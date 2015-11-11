package main

import (
	"encoding/base64"
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
	"os"
	"regexp"
	"strconv"
	"time"
)

// Nyanpass object
type Nyanpass struct {
	Counts    plotter.XYs
	twitter   *Twitter
	imagePath string
	labels    []string
}

// TwitterInfo onject
type TwitterInfo struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// NewNyanpass is constructor of Nyanpass
func NewNyanpass() *Nyanpass {
	info := TwitterInfo{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}

	return &Nyanpass{
		twitter: NewTwitter(info),
	}
}

// GetNyanpassWithDays get nyanpass for days you decided
func (n *Nyanpass) GetNyanpassWithDays(days int) ([]string, error) {
	userID := os.Getenv("USER_ID")
	tweets, err := n.twitter.GetTweets(days, userID)
	if err != nil {
		return nil, err
	}

	reverseTweets(tweets)

	n.Counts = make(plotter.XYs, days)
	now := time.Now()
	cnt := 0
	for i := -days; i < 0; i++ {
		re, err := regexp.Compile("[0-9]+")
		if err != nil {
			return tweets, err
		}
		all := re.FindAllString(tweets[cnt], -1)
		if len(all) != 2 {
			continue
		}
		countF, err := strconv.ParseFloat(all[0], 64)
		if err != nil {
			return tweets, err
		}
		n.Counts[cnt].X = float64(cnt)
		n.Counts[cnt].Y = countF
		n.labels = append(n.labels, now.AddDate(0, 0, i).Format("01/02"))
		cnt++
	}

	return tweets, nil
}

// CreateImage creates graph of nyanpass
func (n *Nyanpass) CreateImage(fileName string) error {
	if n.Counts == nil {
		return errors.New("Count is not defined.")
	}

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.Title.Text = "Nyanpass Graph"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Nyanpass count"
	p.Y.Tick.Marker = &CommaTicks{}
	p.NominalX(n.labels...)

	err = plotutil.AddLinePoints(p, "Nyanpass", n.Counts)
	if err != nil {
		return err
	}
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fileName); err != nil {
		return err
	}
	n.imagePath = fileName
	return nil
}

// PostGraph post Nyanpass graph to Twitter.
func (n *Nyanpass) PostGraph(text string) (*anaconda.Tweet, error) {
	if n.imagePath == "" {
		return nil, errors.New("Could not find images.")
	}

	image, err := os.Open(n.imagePath)
	if err != nil {
		return nil, err
	}
	defer image.Close()
	status, err := image.Stat()
	if err != nil {
		return nil, err
	}
	imageSize := status.Size()
	data := make([]byte, imageSize)
	image.Read(data)

	if text == "" {
		text = n.labels[0] + "~" + n.labels[len(n.labels)-1]
	}

	tweet, err := n.twitter.PostMediaTweet(text, base64.StdEncoding.EncodeToString(data))
	return &tweet, err
}

func reverseTweets(tweets []string) {
	for i, j := 0, len(tweets)-1; i < j; i, j = i+1, j-1 {
		tweets[i], tweets[j] = tweets[j], tweets[i]
	}
}

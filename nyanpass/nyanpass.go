package nyanpass

import (
	"encoding/base64"
	"errors"
	"github.com/ChimeraCoder/anaconda"
	"github.com/Rompei/nyanpass-graph2/twitter"
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
	Counts    plotter.Values
	twitter   *twitter.Twitter
	imagePath string
	labels    []string
}

// NewNyanpass is constructor of Nyanpass
func NewNyanpass(info twitter.TwitterInfo) *Nyanpass {

	return &Nyanpass{
		twitter: twitter.NewTwitter(info),
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

	now := time.Now()
	cnt := 0
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		return tweets, err
	}
	for i := -days; i < 0; i++ {
		all := re.FindAllString(tweets[cnt], -1)
		if len(all) != 3 {
			cnt++
			continue
		}
		countF, err := strconv.ParseFloat(all[0], 64)
		if err != nil {
			return tweets, err
		}
		n.Counts = append(n.Counts, countF)
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

	bar, err := plotter.NewBarChart(n.Counts, vg.Points(20))
	if err != nil {
		return err
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(2)

	p.Add(bar)
	p.Title.Text = "Nyanpass Graph"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Nyanpass count"
	p.NominalX(n.labels...)
	p.Y.Tick.Marker = RelabelTicks{}

	if err := p.Save(6*vg.Inch, 6*vg.Inch, fileName); err != nil {
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

	twText := n.labels[0] + "~" + n.labels[len(n.labels)-1] + "のにゃんぱす\n" + text

	tweet, err := n.twitter.PostMediaTweet(twText, base64.StdEncoding.EncodeToString(data))
	return &tweet, err
}

func reverseTweets(tweets []string) {
	for i, j := 0, len(tweets)-1; i < j; i, j = i+1, j-1 {
		tweets[i], tweets[j] = tweets[j], tweets[i]
	}
}

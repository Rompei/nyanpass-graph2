package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
)

// Twitter object
type Twitter struct {
	api *anaconda.TwitterApi
}

// TwitterInfo onject
type TwitterInfo struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// NewTwitter : constructor of Twitter
func NewTwitter(info TwitterInfo) *Twitter {
	anaconda.SetConsumerKey(info.ConsumerKey)
	anaconda.SetConsumerSecret(info.ConsumerSecret)
	return &Twitter{anaconda.NewTwitterApi(info.AccessToken, info.AccessTokenSecret)}
}

// GetTweets returns the number of tweets you decided
func (t *Twitter) GetTweets(count int, userID string) ([]string, error) {
	v := url.Values{}
	v.Set("count", strconv.Itoa(count))
	v.Set("user_id", userID)
	tweets, err := t.api.GetUserTimeline(v)
	if err != nil {
		return nil, err
	}

	var response []string
	for _, tweet := range tweets {
		response = append(response, tweet.Text)
	}
	return response, nil
}

//PostMediaTweet posts tweet with media
func (t *Twitter) PostMediaTweet(text string, mediaString string) (tweet anaconda.Tweet, err error) {
	media, err := t.api.UploadMedia(mediaString)
	if err != nil {
		panic(err)
	}

	v := url.Values{}
	v.Set("media_ids", media.MediaIDString)
	return t.api.PostTweet(text, v)
}

package main

import (
	"encoding/base64"
	"os"
	"testing"
)

func TestGetTweets(t *testing.T) {
	twitter := NewTwitter(prepareTwitterInfo())
	tweets := twitter.GetTweets(7, os.Getenv("USER_ID"))
	if tweets == nil || len(tweets) != 7 {
		t.Errorf("The number of tweets was wrong.")
	}
}

func TestPostMediaTweet(t *testing.T) {
	twitter := NewTwitter(prepareTwitterInfo())
	image, err := os.Open("lena.png")
	if err != nil {
		t.Errorf("Image was not found")
	}
	defer image.Close()
	stat, err := image.Stat()
	if err != nil {
		t.Errorf("Could not take status")
	}
	size := stat.Size()
	data := make([]byte, size)
	image.Read(data)

	tweet, err := twitter.PostMediaTweet("Test", base64.StdEncoding.EncodeToString(data))
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if tweet.Text != "Test" {
		t.Errorf("Text was wrong")
	}
}

func prepareTwitterInfo() TwitterInfo {
	return TwitterInfo{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}
}

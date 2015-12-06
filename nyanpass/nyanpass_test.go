package nyanpass

import (
	"github.com/Rompei/nyanpass-graph2/twitter"
	"os"
	"testing"
)

func TestCreateImage(t *testing.T) {
	nyanpass := NewNyanpass(twitter.TwitterInfo{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	})
	_, err := nyanpass.GetNyanpassWithDays(7)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = nyanpass.CreateImage("test_graph7.png")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	_, err = nyanpass.GetNyanpassWithDays(30)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	err = nyanpass.CreateImage("test_graph30.png")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

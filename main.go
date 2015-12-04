package main

import (
	"fmt"
	"github.com/Rompei/nyanpass-graph2/nyanpass"
	"github.com/Rompei/nyanpass-graph2/twitter"
	flags "github.com/jessevdk/go-flags"
	"os"
)

func main() {

	opts := parseOptions()

	info := twitter.TwitterInfo{
		ConsumerKey:       opts.ConsumerKey,
		ConsumerSecret:    opts.ConsumerSecret,
		AccessToken:       opts.AccessToken,
		AccessTokenSecret: opts.AccessTokenSecret,
	}

	n := nyanpass.NewNyanpass(info)
	tweets, err := n.GetNyanpassWithDays(7)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		fmt.Println(tweet)
	}

	err = n.CreateImage("nyanpass.png")
	if err != nil {
		panic(err)
	}

	tweet, err := n.PostGraph("")
	if err != nil {
		panic(err)
	}

	fmt.Println(tweet.Text)
}

// Options flags of this program
type Options struct {
	ConsumerKey       string `short:"c" long:"consumer-key" description:"Twitter consumer key" required:"true"`
	ConsumerSecret    string `short:"s" long:"consumer-secret" description:"Twitter consumer secret" required:"true"`
	AccessToken       string `short:"a" long:"access-token" description:"Twitter access token" required:"true"`
	AccessTokenSecret string `short:"t" long:"access-token-secret" description:"Twitter access token secret" required:"true"`
}

func parseOptions() *Options {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "Nyanpass-graph2"
	parser.Usage = "[OPTIONS]"
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	if opts.ConsumerKey == "" || opts.ConsumerSecret == "" || opts.AccessToken == "" || opts.AccessTokenSecret == "" {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	return &opts
}

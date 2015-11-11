package main

import (
	"fmt"
	"github.com/Rompei/nyanpass-graph2/nyanpass"
)

func main() {
	n := nyanpass.NewNyanpass()
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

package main

import (
	"fmt"
)

func main() {
	nyanpass := NewNyanpass()
	tweets, err := nyanpass.GetNyanpassWithDays(7)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		fmt.Println(tweet)
	}

	err = nyanpass.CreateImage("nyanpass.png")
	if err != nil {
		panic(err)
	}

	tweet, err := nyanpass.PostGraph("")
	if err != nil {
		panic(err)
	}

	fmt.Println(tweet.Text)
}

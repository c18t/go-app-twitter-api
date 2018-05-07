package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	homedir "github.com/mitchellh/go-homedir"
)

type GatapiConfig struct {
	Key KeyConfig
}

type KeyConfig struct {
	Consumer          string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func main() {
	// get consumer key settings
	var gatapiConfig GatapiConfig
	var configFilePath string
	var err error
	configFilePath, err = homedir.Expand("~/.config/gatapi/config.toml")
	if err != nil {
		os.Exit(1)
	}

	_, err = toml.DecodeFile(configFilePath, &gatapiConfig)
	if err != nil {
		os.Exit(1)
	}

	config := oauth1.NewConfig(gatapiConfig.Key.Consumer, gatapiConfig.Key.ConsumerSecret)
	token := oauth1.NewToken(gatapiConfig.Key.AccessToken, gatapiConfig.Key.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// User Timeline
	tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count: 10,
	})

	fmt.Printf("User's USER TIMELINE:\n")
	for i := 0; i < len(tweets); i++ {
		fmt.Printf("@%+v: %+v\n", tweets[i].User.ScreenName, tweets[i].Text)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// TweetParams represents ...
//   ID:
//   To:
type TweetParams struct {
	ID int64
	To int64
}

var tweetParams, sayParams TweetParams

// tweetCmd represents the tweet command
var tweetCmd = &cobra.Command{
	Use:   "tweet <text>",
	Short: "Post tweet",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Post tweet`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		var reply bool = false
		var targetTweet = &twitter.Tweet{}
		var updateParams = &twitter.StatusUpdateParams{}
		if tweetParams.ID != 0 {
			var err error
			reply = true
			targetTweet, _, err = client.Statuses.Show(tweetParams.ID, nil)
			if err != nil {
				fmt.Printf("Update Error: %+v\n", err)
				os.Exit(1)
			}

			updateParams.InReplyToStatusID = tweetParams.ID
		}

		tweet, _, _ := client.Statuses.Update(args[0], updateParams)

		if reply == false {
			fmt.Printf("Update Tweet: [%+v] @%+v: %+v\n", tweet.IDStr, tweet.User.ScreenName, tweet.Text)
		} else {
			fmt.Printf("Update Tweet: [%+v] @%+v: %+v (reply to @%+v: %+v)\n", tweet.IDStr, tweet.User.ScreenName, tweet.Text, targetTweet.User.ScreenName, targetTweet.Text)
		}
	},
}

var sayCmd = &cobra.Command{
	Use:   "say <text>",
	Short: tweetCmd.Short,
	Long:  tweetCmd.Long,
	Args:  tweetCmd.Args,
	PreRun: func(cmd *cobra.Command, args []string) {
		tweetParams = sayParams
	},
	Run: tweetCmd.Run,
}

func init() {
	rootCmd.AddCommand(tweetCmd, sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tweetCmd.Flags().Int64Var(&tweetParams.ID, "to", 0, "reply to <tweet id>") // (it takes priority of '--to' option)")
	//tweetCmd.Flags().StrVar(&tweetParams.To, "to", nil, "reply to <screen name>")
	sayCmd.Flags().Int64Var(&sayParams.ID, "id", 0, "reply to <tweet id>") // (it takes priority of '--to' option)")
	//sayCmd.Flags().StrVar(&sayParams.To, "to", nil, "reply to <screen name>")
}

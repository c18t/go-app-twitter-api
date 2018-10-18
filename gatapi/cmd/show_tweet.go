package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowTweetParams represents ...
//   Limit:
type ShowTweetParams struct {
	Limit int
}

var showTweetParams ShowTweetParams

// showTweetCmd represents the show tweet command
var showTweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Show user's user timeline",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show user's user timeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		// User Timeline
		tweets, _, _ := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
			Count: showTweetParams.Limit,
		})

		fmt.Println("User's USER TIMELINE:")
		for i := 0; i < len(tweets); i++ {
			fmt.Printf("[%+v] @%+v: %+v\n", tweets[i].IDStr, tweets[i].User.ScreenName, tweets[i].Text)
		}
	},
}

func init() {
	showCmd.AddCommand(showTweetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showTweetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showTweetCmd.Flags().IntVarP(&showTweetParams.Limit, "limit", "l", 20, "tweet count")
}

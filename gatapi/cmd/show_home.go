package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowHomeParams represents ...
//   Limit:
type ShowHomeParams struct {
	Limit int
}

var showHomeParams ShowHomeParams

// showHomeCmd represents the show tweet command
var showHomeCmd = &cobra.Command{
	Use:   "home",
	Short: "Show user's home timeline",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show user's home timeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		// Home Timeline
		tweets, _, _ := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
			Count: showHomeParams.Limit,
		})

		fmt.Println("User's HOME TIMELINE:")
		for i := 0; i < len(tweets); i++ {
			fmt.Printf("[%+v] @%+v: %+v\n", tweets[i].IDStr, tweets[i].User.ScreenName, tweets[i].Text)
		}
	},
}

func init() {
	showCmd.AddCommand(showHomeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showHomeCmd.Flags().IntVarP(&showHomeParams.Limit, "limit", "l", 20, "tweet count")
}

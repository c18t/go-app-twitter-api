package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowTweetParams represents ...
//   Limit:
type ShowTweetParams struct {
	Limit   int
	Page    int
	SinceID int64
	MaxID   int64
	Retry   int
}

var showTweetParams ShowTweetParams

// showTweetCmd represents the show tweet command
var showTweetCmd = &cobra.Command{
	Use:   "tweet [<screen name>]",
	Short: "Show <screen name>'s user timeline",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show <screen name>'s user timeline.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		var page, retry int
		var sinceID, maxID int64
		var screenName string

		// <screen name>の指定があれば対象を設定
		if len(args) > 0 {
			screenName = args[0]
		}

		page = showTweetParams.Page
		retry = showTweetParams.Retry
		sinceID = showTweetParams.SinceID
		maxID = showTweetParams.MaxID

		if showTweetParams.Page == 0 {
			if showTweetParams.SinceID == 0 {
				// 取得条件の指定がなければ1ページだけ取得
				page = 1
			} else {
				// --since-id まで取得
				page = -1
			}
		}

		for page != 0 && retry > 0 {
			// User Timeline
			params := &twitter.UserTimelineParams{
				Count: showTweetParams.Limit,
			}
			if screenName != "" {
				params.ScreenName = screenName
			}
			if sinceID != 0 {
				params.SinceID = sinceID
			}
			if maxID != 0 {
				params.MaxID = maxID
			}
			tweets, _, err := client.Timelines.UserTimeline(params)
			if err != nil {
				fmt.Printf("Show Tweet Error: %+v\n", err)
				retry--
				continue
			}

			if len(tweets) < 1 {
				// 取得できるツイートがなければ終了
				break
			}

			for i := 0; i < len(tweets); i++ {
				fmt.Printf("[%+v] @%+v: %+v\n", tweets[i].IDStr, tweets[i].User.ScreenName, tweets[i].Text)
			}

			maxID = tweets[len(tweets)-1].ID - 1
			if page > 0 {
				page--
			}
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
	showTweetCmd.Flags().IntVarP(&showTweetParams.Limit, "limit", "l", 20, "tweet count (max 200)")
	showTweetCmd.Flags().IntVarP(&showTweetParams.Page, "page", "p", 0, "page")
	showTweetCmd.Flags().Int64VarP(&showTweetParams.SinceID, "since-id", "s", 0, "since <tweet id>") // (it takes priority of '--to' option)")
	showTweetCmd.Flags().Int64VarP(&showTweetParams.MaxID, "max-id", "m", 0, "max <tweet id>")       // (it takes priority of '--to' option)")
	showTweetCmd.Flags().IntVarP(&showTweetParams.Retry, "retry", "r", 3, "retry count")
}

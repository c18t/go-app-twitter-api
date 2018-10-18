package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowHomeParams represents ...
//   Limit:
type ShowHomeParams struct {
	Limit   int
	Page    int
	SinceID int64
	MaxID   int64
	Retry   int
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

		var page, retry int
		var sinceID, maxID int64

		page = showHomeParams.Page
		retry = showHomeParams.Retry
		sinceID = showHomeParams.SinceID
		maxID = showHomeParams.MaxID

		if showHomeParams.Page == 0 {
			if showHomeParams.SinceID == 0 {
				// 取得条件の指定がなければ1ページだけ取得
				page = 1
			} else {
				// --since-id まで取得
				page = -1
			}
		}

		for page != 0 && retry > 0 {
			// User Timeline
			params := &twitter.HomeTimelineParams{
				Count: showHomeParams.Limit,
			}
			if sinceID != 0 {
				params.SinceID = sinceID
			}
			if maxID != 0 {
				params.MaxID = maxID
			}
			tweets, _, err := client.Timelines.HomeTimeline(params)
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
	showCmd.AddCommand(showHomeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showHomeCmd.Flags().IntVarP(&showHomeParams.Limit, "limit", "l", 20, "tweet count (max 200)")
	showHomeCmd.Flags().IntVarP(&showHomeParams.Page, "page", "p", 0, "page")
	showHomeCmd.Flags().Int64VarP(&showHomeParams.SinceID, "since-id", "s", 0, "since <tweet id>") // (it takes priority of '--to' option)")
	showHomeCmd.Flags().Int64VarP(&showHomeParams.MaxID, "max-id", "m", 0, "max <tweet id>")       // (it takes priority of '--to' option)")
	showHomeCmd.Flags().IntVarP(&showHomeParams.Retry, "retry", "r", 3, "retry count")
}

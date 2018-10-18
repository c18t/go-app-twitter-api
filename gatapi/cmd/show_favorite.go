package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowFavoriteParams represents ...
//   Limit:
type ShowFavoriteParams struct {
	Limit   int
	Page    int
	SinceID int64
	MaxID   int64
	Retry   int
}

var showFavoriteParams ShowFavoriteParams

// showFavoriteCmd represents the show favorite command
var showFavoriteCmd = &cobra.Command{
	Use:   "favorite [<screen name>]",
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

		page = showFavoriteParams.Page
		retry = showFavoriteParams.Retry
		sinceID = showFavoriteParams.SinceID
		maxID = showFavoriteParams.MaxID

		if showFavoriteParams.Page == 0 {
			if showFavoriteParams.SinceID == 0 {
				// 取得条件の指定がなければ1ページだけ取得
				page = 1
			} else {
				// --since-id まで取得
				page = -1
			}
		}

		for page != 0 && retry > 0 {
			// User Timeline
			params := &twitter.FavoriteListParams{
				Count: showFavoriteParams.Limit,
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
			tweets, _, err := client.Favorites.List(params)
			if err != nil {
				fmt.Printf("Show Favorite Error: %+v\n", err)
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
	showCmd.AddCommand(showFavoriteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showTweetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showFavoriteCmd.Flags().IntVarP(&showFavoriteParams.Limit, "limit", "l", 20, "tweet count")
	showFavoriteCmd.Flags().IntVarP(&showFavoriteParams.Page, "page", "p", 0, "page")
	showFavoriteCmd.Flags().Int64VarP(&showFavoriteParams.SinceID, "since-id", "s", 0, "since <tweet id>") // (it takes priority of '--to' option)")
	showFavoriteCmd.Flags().Int64VarP(&showFavoriteParams.MaxID, "max-id", "m", 0, "max <tweet id>")       // (it takes priority of '--to' option)")
	showFavoriteCmd.Flags().IntVarP(&showFavoriteParams.Retry, "retry", "r", 3, "retry count")
}

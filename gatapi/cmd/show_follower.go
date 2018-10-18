package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowFollowerParams represents ...
//   Limit:
type ShowFollowerParams struct {
	Limit int
	Page  int
	All   bool
	AsID  bool
	Retry int
}

var showFollowerParams ShowFollowerParams

// showHomeCmd represents the show tweet command
var showFollowerCmd = &cobra.Command{
	Use:   "follower [<screen name>]",
	Short: "Show <screen name>'s follower",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show <screen name>'s follower.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		var page, retry, limit int
		var screenName string
		var cursor int64 = -1

		page = showFollowerParams.Page
		retry = showFollowerParams.Retry
		limit = showFollowerParams.Limit

		// <screen name>の指定があれば対象を設定
		if len(args) > 0 {
			screenName = args[0]
		}

		if showFollowerParams.AsID == true {
			for retry > 0 {
				params := &twitter.FollowerIDParams{
					Count:  5000, // max
					Cursor: cursor,
				}
				if screenName != "" {
					params.ScreenName = screenName
				}
				follower, _, err := client.Followers.IDs(params)
				if err != nil {
					fmt.Printf("Show Follower Error: %+v\n", err)
					retry--
					continue
				}

				if len(follower.IDs) < 1 {
					// 取得できるユーザーがなければ終了
					break
				}

				for i := 0; i < len(follower.IDs); i++ {
					fmt.Println(follower.IDs[i])
				}

				cursor = follower.NextCursor

				if cursor == 0 {
					break
				}
			}
		} else {
			if page == 0 {
				// 未指定時は1ページ取得
				page = 1
			}

			if showFollowerParams.All == true {
				// --all指定時は最大数で全て取得
				limit = 200
				page = -1
			}

			for page != 0 && retry > 0 {
				params := &twitter.FollowerListParams{
					Count:  limit,
					Cursor: cursor,
				}
				if screenName != "" {
					params.ScreenName = screenName
				}
				follower, _, err := client.Followers.List(params)
				if err != nil {
					fmt.Printf("Show Follower Error: %+v\n", err)
					retry--
					continue
				}

				if len(follower.Users) < 1 {
					// 取得できるユーザーがなければ終了
					break
				}

				for i := 0; i < len(follower.Users); i++ {
					fmt.Printf("@%+v: %+v (%+v)\n", follower.Users[i].ScreenName, follower.Users[i].Name, follower.Users[i].IDStr)
				}

				cursor = follower.NextCursor
				if page > 0 {
					page--
				}

				if cursor == 0 {
					break
				}
			}
		}
	},
}

func init() {
	showCmd.AddCommand(showFollowerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showFollowerCmd.Flags().IntVarP(&showFollowerParams.Limit, "limit", "l", 20, "user count (max 200)")
	showFollowerCmd.Flags().IntVarP(&showFollowerParams.Page, "page", "p", 0, "page")
	showFollowerCmd.Flags().BoolVarP(&showFollowerParams.All, "all", "a", false, "show all users (it takes priority of '--page' option)")
	showFollowerCmd.Flags().BoolVar(&showFollowerParams.AsID, "as-id", false, "get user's ids")
	showFollowerCmd.Flags().IntVarP(&showFollowerParams.Retry, "retry", "r", 3, "retry count")
}

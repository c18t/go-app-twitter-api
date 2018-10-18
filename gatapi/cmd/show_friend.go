package cmd

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowFriendParams represents ...
//   Limit:
type ShowFriendParams struct {
	Limit int
	Page  int
	All   bool
	AsID  bool
	Retry int
}

var showFriendParams ShowFriendParams

// showHomeCmd represents the show tweet command
var showFriendCmd = &cobra.Command{
	Use:   "friend [<screen name>]",
	Short: "Show <screen name>'s friends",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show <screen name>'s friends.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		var page, retry, limit int
		var screenName string
		var cursor int64 = -1

		page = showFriendParams.Page
		retry = showFriendParams.Retry
		limit = showFriendParams.Limit

		// <screen name>の指定があれば対象を設定
		if len(args) > 0 {
			screenName = args[0]
		}

		if showFriendParams.AsID == true {
			for retry > 0 {
				params := &twitter.FriendIDParams{
					Count:  5000, // max
					Cursor: cursor,
				}
				if screenName != "" {
					params.ScreenName = screenName
				}
				friends, _, err := client.Friends.IDs(params)
				if err != nil {
					fmt.Printf("Show Friend Error: %+v\n", err)
					retry--
					continue
				}

				if len(friends.IDs) < 1 {
					// 取得できるユーザーがなければ終了
					break
				}

				for i := 0; i < len(friends.IDs); i++ {
					fmt.Println(friends.IDs[i])
				}

				cursor = friends.NextCursor

				if cursor == 0 {
					break
				}
			}
		} else {
			if page == 0 {
				// 未指定時は1ページ取得
				page = 1
			}

			if showFriendParams.All == true {
				// --all指定時は最大数で全て取得
				limit = 200
				page = -1
			}

			for page != 0 && retry > 0 {
				params := &twitter.FriendListParams{
					Count:  limit,
					Cursor: cursor,
				}
				if screenName != "" {
					params.ScreenName = screenName
				}
				friends, _, err := client.Friends.List(params)
				if err != nil {
					fmt.Printf("Show Friend Error: %+v\n", err)
					retry--
					continue
				}

				if len(friends.Users) < 1 {
					// 取得できるユーザーがなければ終了
					break
				}

				for i := 0; i < len(friends.Users); i++ {
					fmt.Printf("@%+v: %+v (%+v)\n", friends.Users[i].ScreenName, friends.Users[i].Name, friends.Users[i].IDStr)
				}

				cursor = friends.NextCursor
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
	showCmd.AddCommand(showFriendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showFriendCmd.Flags().IntVarP(&showFriendParams.Limit, "limit", "l", 20, "user count (max 200)")
	showFriendCmd.Flags().IntVarP(&showFriendParams.Page, "page", "p", 0, "page")
	showFriendCmd.Flags().BoolVarP(&showFriendParams.All, "all", "a", false, "show all users (it takes priority of '--page' option)")
	showFriendCmd.Flags().BoolVar(&showFriendParams.AsID, "as-id", false, "get user's ids")
	showFriendCmd.Flags().IntVarP(&showFriendParams.Retry, "retry", "r", 3, "retry count")
}

package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/spf13/cobra"
)

// ShowListParams represents ...
//   Limit:
type ShowListParams struct {
	Limit int
	Page  int
	All   bool
	AsID  bool
	Retry int
}

var showListParams ShowListParams

// showHomeCmd represents the show tweet command
var showListCmd = &cobra.Command{
	Use:   "list [<screen name>]",
	Short: "Show <screen name>'s user lists",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show <screen name>'s user lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getAnacondaClient()

		var page, retry int
		var screenName string

		page = showListParams.Page
		retry = showListParams.Retry

		// <screen name>の指定があれば対象を設定
		if len(args) > 0 {
			screenName = args[0]
		}

		user, err := client.GetUsersShow(screenName, nil)
		if err != nil {
			fmt.Printf("Show List Error: %+v\n", err)
			os.Exit(1)
		}

		if page == 0 {
			// 未指定時は1ページ取得
			page = 1
		}

		v := url.Values{}
		v.Set("count", "200")
		if screenName != "" {
			v.Set("screenName", screenName)
		}

		var lists []anaconda.List
		for retry > 0 {
			lists, err = client.GetListsOwnedBy(user.Id, v)
			if err != nil {
				fmt.Printf("Show List Error: %+v\n", err)
				retry--
				continue
			}
			break
		}

		for _, list := range lists {
			fmt.Printf("[%+v] %+v: %+v\n", list.Id, list.Slug, list.Description)
		}

		// cursor = lists.NextCursor
		// if page > 0 {
		// 	page--
		// }

		// if cursor == 0 {
		// 	break
		// }
	},
}

func init() {
	showCmd.AddCommand(showListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showListCmd.Flags().IntVarP(&showListParams.Limit, "limit", "l", 20, "user count (max 200)")
	// showListCmd.Flags().IntVarP(&showListParams.Page, "page", "p", 0, "page")
	// showListCmd.Flags().BoolVarP(&showListParams.All, "all", "a", false, "show all users (it takes priority of '--page' option)")
	showListCmd.Flags().IntVarP(&showListParams.Retry, "retry", "r", 3, "retry count")
}

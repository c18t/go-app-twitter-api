package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

type ShowListMemberParams struct {
	ListID int64
	Limit  int
	Page   int
	All    bool
	AsID   bool
	Retry  int
}

var showListMemberParams ShowListMemberParams

// showListMemberCmd represents the showListMember command
var showListMemberCmd = &cobra.Command{
	Use:   "list-member [[<screen name>'/']<list-slug>]",
	Short: "Show members of the specified list",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show members of the specified list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getTwitterClient()

		var screenName, slug string
		var page, retry, limit int
		var listID int64 = -1
		var cursor int64 = -1

		listID = showListMemberParams.ListID
		page = showListMemberParams.Page
		retry = showListMemberParams.Retry
		limit = showListMemberParams.Limit

		if listID == 0 {
			if len(args) < 1 {
				return fmt.Errorf("accepts %d arg(s), received %d", 1, len(args))
			} else if strings.Index(args[0], "/") != -1 {
				temp := strings.SplitN(args[0], "/", 2)
				screenName = temp[0]
				slug = temp[1]
			} else {
				params := &twitter.AccountVerifyParams{
					IncludeEntities: twitter.Bool(false),
					SkipStatus:      twitter.Bool(true),
					IncludeEmail:    twitter.Bool(false),
				}
				verify, _, err := client.Accounts.VerifyCredentials(params)
				if err != nil {
					fmt.Printf("Show List's Members Error: %+v\n", err)
					os.Exit(1)
				}
				screenName = verify.ScreenName
				slug = args[0]
			}
		}

		if page == 0 {
			// 未指定時は1ページ取得
			page = 1
		}

		if showListMemberParams.All == true {
			// --all指定時は最大数で全て取得
			limit = 5000
			page = -1
		}

		for page != 0 && retry > 0 {
			params := &twitter.ListsMembersParams{
				Count:           limit,
				Cursor:          cursor,
				IncludeEntities: twitter.Bool(false),
				SkipStatus:      twitter.Bool(true),
			}
			if listID != 0 {
				params.ListID = listID
			} else {
				params.Slug = slug
				if screenName != "" {
					params.OwnerScreenName = screenName
				}
			}
			members, _, err := client.Lists.Members(params)
			if err != nil {
				fmt.Printf("Show List's Members Error: %+v\n", err)
				retry--
				continue
			}

			if len(members.Users) < 1 {
				// 取得できるユーザーがなければ終了
				break
			}

			for i := 0; i < len(members.Users); i++ {
				if showListMemberParams.AsID == true {
					fmt.Println(members.Users[i].IDStr)
				} else {
					fmt.Printf("@%+v: %+v (%+v)\n", members.Users[i].ScreenName, members.Users[i].Name, members.Users[i].IDStr)
				}
			}

			cursor = members.NextCursor
			if page > 0 {
				page--
			}

			if cursor == 0 {
				break
			}
		}

		return nil
	},
}

func init() {
	showCmd.AddCommand(showListMemberCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showListMemberCmd.Flags().Int64Var(&showListMemberParams.ListID, "list-id", 0, "target list id")
	showListMemberCmd.Flags().IntVarP(&showListMemberParams.Limit, "limit", "l", 20, "member count (max 5000)")
	showListMemberCmd.Flags().IntVarP(&showListMemberParams.Page, "page", "p", 0, "page")
	showListMemberCmd.Flags().BoolVarP(&showListMemberParams.All, "all", "a", false, "show all members (it takes priority of '--page' option)")
	showListMemberCmd.Flags().BoolVar(&showListMemberParams.AsID, "as-id", false, "get member's ids")
	showListMemberCmd.Flags().IntVarP(&showListMemberParams.Retry, "retry", "r", 3, "retry count")
}

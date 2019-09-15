package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ListAddParams represents ...
//   Limit:
type ListAddParams struct {
	ListID int64
}

var listAddParams ListAddParams

// showHomeCmd represents the show tweet command
var listAddCmd = &cobra.Command{
	Use:   "add <screen name> [, <screen name>]*",
	Short: "Add <screen name> to user lists",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Add <screen name> to user lists.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		params := &twitter.ListsMembersCreateAllParams{
			ListID:     listAddParams.ListID,
			ScreenName: strings.Join(args, ","),
		}
		_, err := client.Lists.MembersCreateAll(params)
		if err != nil {
			fmt.Printf("Add List Member Error: %+v\n", err)
			os.Exit(1)
		}

		//fmt.Printf("Added List Member to [%+v] %+v: %+v\n", list.Id, list.Slug, list.Description)
		fmt.Printf("Added List Member to [%+v]\n", params.ListID)
	},
}

func init() {
	listCmd.AddCommand(listAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listAddCmd.Flags().Int64VarP(&listAddParams.ListID, "list-id", "l", 0, "target list id")
	listAddCmd.MarkFlagRequired("list-id")
}

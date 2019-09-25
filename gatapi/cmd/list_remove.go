package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

type ListRemoveParams struct {
	ListID int64
}

var listRemoveParams ListRemoveParams

// listRemoveCmd represents the listRemove command
var listRemoveCmd = &cobra.Command{
	Use:   "remove <screen name> [, <screen name>]*",
	Short: "Remove <screen name> from user lists",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Remove <screen name> from user lists`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		params := &twitter.ListsMembersDestroyAllParams{
			ListID:     listRemoveParams.ListID,
			ScreenName: strings.Join(args, ","),
		}
		_, err := client.Lists.MembersDestroyAll(params)
		if err != nil {
			fmt.Printf("Remove List Member Error: %+v\n", err)
			os.Exit(1)
		}

		//fmt.Printf("Removed List Member from [%+v] %+v: %+v\n", list.Id, list.Slug, list.Description)
		fmt.Printf("Removed List Member from [%+v]\n", params.ListID)
	},
}

func init() {
	listCmd.AddCommand(listRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listRemoveCmd.Flags().Int64VarP(&listRemoveParams.ListID, "list-id", "l", 0, "target list id")
	listRemoveCmd.MarkFlagRequired("list-id")
}

package cmd

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// ShowProfileParams represents ...
//   Limit:
type ShowProfileParams struct {
	ByID bool
}

var showProfileParams ShowProfileParams

// showHomeCmd represents the show tweet command
var showProfileCmd = &cobra.Command{
	Use:   "profile (<screen name>, [<screen name>*] | <user id list>)",
	Short: "Show <screen name>'s profile",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Show <screen name>'s profile.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		for i := 0; i < len(args); i++ {
			params := &twitter.UserShowParams{}

			if showProfileParams.ByID == true {
				id, err := strconv.ParseInt(args[i], 10, 64)
				if err != nil {
					fmt.Printf("Show profile failed: [%+v] %+v\n", args[i], err)
					continue
				}
				params.UserID = id
			} else {
				params.ScreenName = args[i]
			}

			user, _, err := client.Users.Show(params)
			if err != nil {
				fmt.Printf("Show Profile Error: %+v\n", err)
				continue
			}

			fmt.Printf("@%+v: %+v (%+v)\n", user.ScreenName, user.Name, user.IDStr)
		}
	},
}

func init() {
	showCmd.AddCommand(showProfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showHomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	showProfileCmd.Flags().BoolVar(&showProfileParams.ByID, "by-id", false, "get profile by user id")
}

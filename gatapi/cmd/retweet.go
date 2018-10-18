package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var retweetCmd = &cobra.Command{
	Use:   "retweet <tweet id>",
	Short: "Retweet some tweets",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Retweet some tweets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		for i := 0; i < len(args); i++ {
			id, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				fmt.Printf("Retweet failed: [%+v] %+v\n", args[i], err)
				continue
			}
			tweet, _, err := client.Statuses.Retweet(id, nil)
			if err != nil {
				fmt.Printf("Retweet failed: [%+v] %+v\n", args[i], err)
				continue
			}

			fmt.Printf("Retweet Tweet: [%+v] @%+v: %+v\n", tweet.IDStr, tweet.User.ScreenName, tweet.Text)
		}
	},
}

var rtCmd = &cobra.Command{
	Use:   "rt <text>",
	Short: retweetCmd.Short,
	Long:  retweetCmd.Long,
	Args:  retweetCmd.Args,
	Run:   retweetCmd.Run,
}

func init() {
	rootCmd.AddCommand(retweetCmd, rtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

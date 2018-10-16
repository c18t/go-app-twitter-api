package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var tweetCmd = &cobra.Command{
	Use:   "tweet <text>",
	Short: "Post tweet",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Post tweet`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		tweet, _, _ := client.Statuses.Update(args[0], nil)

		fmt.Printf("Update Tweet: %s\n", tweet.Text)
	},
}

var sayCmd = &cobra.Command{
	Use:   "say <text>",
	Short: tweetCmd.Short,
	Long:  tweetCmd.Long,
	Args:  tweetCmd.Args,
	Run:   tweetCmd.Run,
}

func init() {
	rootCmd.AddCommand(tweetCmd, sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

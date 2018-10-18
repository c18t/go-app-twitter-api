package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// QuoteTweetParams represents ...
//   Limit:
type QuoteTweetParams struct {
	ID int64
}

var quoteTweetParamas, qtParamas QuoteTweetParams

// quoteCmd represents the quote twewet command
var quoteCmd = &cobra.Command{
	Use:   "quote <text>",
	Short: "Post quote tweet",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Post quote tweet`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		targetTweet, _, err := client.Statuses.Show(quoteTweetParamas.ID, nil)
		if err != nil {
			fmt.Printf("Update Error: %+v\n", err)
			os.Exit(1)
		}

		text := fmt.Sprintf("%+v https://twitter.com/%+v/status/%+v", args[0], targetTweet.User.ScreenName, quoteTweetParamas.ID)
		tweet, _, err := client.Statuses.Update(text, nil)
		if err != nil {
			fmt.Printf("Update Error: %+v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Update Tweet: [%+v] @%+v: %+v\n", tweet.IDStr, tweet.User.ScreenName, tweet.Text)
	},
}

var qtCmd = &cobra.Command{
	Use:   "qt <text>",
	Short: quoteCmd.Short,
	Long:  quoteCmd.Long,
	Args:  quoteCmd.Args,
	PreRun: func(cmd *cobra.Command, args []string) {
		quoteTweetParamas = qtParamas
	},
	Run: quoteCmd.Run,
}

func init() {
	rootCmd.AddCommand(quoteCmd, qtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	quoteCmd.Flags().Int64Var(&quoteTweetParamas.ID, "id", 0, "tweet id")
	quoteCmd.MarkFlagRequired("id")
	qtCmd.Flags().Int64Var(&qtParamas.ID, "id", 0, "tweet id")
	qtCmd.MarkFlagRequired("id")
}

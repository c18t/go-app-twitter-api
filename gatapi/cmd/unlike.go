package cmd

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var unlikeCmd = &cobra.Command{
	Use:   "unlike <tweet id> [, <tweet id>]*",
	Short: "Unlike some tweets",
	Long: `Gatapi - Go Application for Twitter API. (for study)
Unlike some tweets`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getTwitterClient()

		for i := 0; i < len(args); i++ {
			id, err := strconv.ParseInt(args[i], 10, 64)
			if err != nil {
				fmt.Printf("Unlike failed: [%+v] %+v\n", args[i], err)
				continue
			}
			tweet, _, err := client.Favorites.Destroy(&twitter.FavoriteDestroyParams{ID: id})
			if err != nil {
				fmt.Printf("Unlike failed: [%+v] %+v\n", args[i], err)
				continue
			}

			fmt.Printf("Unlike Tweet: [%+v] @%+v: %+v\n", tweet.IDStr, tweet.User.ScreenName, tweet.Text)
		}
	},
}

var unfavCmd = &cobra.Command{
	Use:   "unf <text>",
	Short: unlikeCmd.Short,
	Long:  unlikeCmd.Long,
	Args:  unlikeCmd.Args,
	Run:   unlikeCmd.Run,
}

func init() {
	rootCmd.AddCommand(unlikeCmd, unfavCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

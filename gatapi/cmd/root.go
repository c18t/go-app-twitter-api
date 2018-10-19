package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/ChimeraCoder/anaconda"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gatapi",
	Short: "Gatapi - Go Application for Twitter API. (for study)",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/gatapi/config.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defaultConfigDir := path.Join(home, ".config")
		configDir := os.Getenv("XDG_CONFIG_HOME")
		if configDir == "" || path.IsAbs(configDir) == false {
			configDir = defaultConfigDir
		}
		configPath := path.Join(configDir, "gatapi")

		// Search config in home directory with name ".gatapi" (without extension).
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getTwitterClient() *twitter.Client {
	config := oauth1.NewConfig(viper.GetString("Key.Consumer"), viper.GetString("Key.ConsumerSecret"))
	token := oauth1.NewToken(viper.GetString("Key.AccessToken"), viper.GetString("Key.AccessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)
}

func getAnacondaClient() *anaconda.TwitterApi {
	// Twitter client
	return anaconda.NewTwitterApiWithCredentials(viper.GetString("Key.AccessToken"), viper.GetString("Key.AccessTokenSecret"), viper.GetString("Key.Consumer"), viper.GetString("Key.ConsumerSecret"))
}

package cmd

import (
	"fmt"
	// "log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config string

type Opts struct {
	BaseUrl string
}

var opts Opts

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "~/.elastic", "config file")
	rootCmd.PersistentFlags().StringVarP(&opts.BaseUrl, "baseurl", "b", "http://localhost:9200", "Elasticsearch base url")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("baseurl", rootCmd.PersistentFlags().Lookup("baseurl"))

	viper.SetDefault("config", "~/.elastic")
	viper.SetDefault("baseurl", "http://localhost:9200")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	viper.SetConfigType("yaml")
	if viper.GetString("config") != "" {
		// fmt.Printf("Setting config auto: %s\n", viper.GetString("config"))

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// log.Printf("Found home: %s\n", home)

		config := viper.GetString("config")
		// log.Printf("Current config: %s\n", config)
		if strings.HasPrefix(config, "$HOME") || strings.HasPrefix(config, "~") {
			// log.Printf("Replacing config: %s\n", config)
			config = strings.Replace(config, "$HOME", home, 1)
			config = strings.Replace(config, "~", home, 1)
		}
		viper.Set("config", config)
		viper.SetConfigFile(viper.GetString("config"))
	}

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println("Can't read config:", err)
	}

	config = viper.GetString("config")
	opts.BaseUrl = viper.GetString("baseurl")

	// log.Printf("Config: %v\n", opts)
}

var rootCmd = &cobra.Command{
	Use:   "elastic",
	Short: "elastic is a handy commandline Elasticsearch client meant for admins",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

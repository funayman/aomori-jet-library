package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/funayman/aomori-library/client"
	"github.com/funayman/aomori-library/db"
)

const (
	defaultConfigFile = ".honshitsu"
)

var (
	cmdName    = os.Args[0]
	dbReadOnly bool
	database   string
	cfgFile    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   cmdName,
	Short: "The Aomori JET Public Library",
	Long: `This allows you to run the application server along with other
helpful commands.

The Aomori JET Public Library is an open sourced web application
written in Go. And is made possible by the help of Dan Hantos
and Dave Derderian (https://drt.sh).`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("connecting to db...")
		db.Connect(viper.GetString("db.conn"), viper.GetBool("db.readonly"))

		fmt.Println("initing isbn/book clients")
		client.Init(&client.Cfg{
			IsbnDbKey:    viper.GetString("api.isbndb"),
			GoodReadsKey: viper.GetString("api.goodreads"),
		})

	},

	Run: func(cmd *cobra.Command, args []string) {},
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/"+defaultConfigFile+".yaml)")

	rootCmd.PersistentFlags().StringP("database", "d", db.DefaultConnection, "connects to BoltDB database at specified location (default ./"+db.DefaultConnection+")")
	rootCmd.PersistentFlags().BoolP("readonly", "r", false, "connects to database in read-only mode. changes cannot be made")

	viper.BindPFlag("db.conn", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("db.readonly", rootCmd.PersistentFlags().Lookup("readonly"))
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

		// Search config in home directory and current directory with name ".honshitsu" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".honshitsu")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

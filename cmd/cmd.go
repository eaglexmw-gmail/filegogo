package cmd

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "filegogo",
		Short: "a p2p file transfer tool",
		Long:  `A p2p file transfer tool that can be used in the webrtc p2p`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/filegogo.toml)")
	rootCmd.PersistentFlags().StringP("server", "s", "", "Signal Server Address (default is https://send.22333.fun)")
	rootCmd.PersistentFlags().BoolP("show-qrcode", "q", false, "Show QRcode")
	rootCmd.PersistentFlags().BoolP("show-progress", "p", true, "Show Progress Bar")

	rootCmd.PersistentFlags().StringP("level", "", "info", "log level")
	//viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlags(rootCmd.PersistentFlags())

	viper.SetDefault("server", "http://localhost:8033")
	viper.SetDefault("level", "info")

	// server
	viper.SetDefault("listen", "0.0.0.0:8033")
	viper.SetDefault("browser", "./config.json")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.SetConfigName("filegogo")
		viper.SetConfigType("toml")

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath(home + "/.config/")
		viper.AddConfigPath("/etc/filegogo/")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//log.SetReportCaller(true)
	if level, err := log.ParseLevel(viper.GetString("level")); err != nil {
		fmt.Println(err)
	} else {
		log.SetLevel(level)
	}
	log.SetFormatter(&log.TextFormatter{
		//FullTimestamp: true,
	})
}

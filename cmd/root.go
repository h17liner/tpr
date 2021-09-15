package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	config string
	debug  bool

	rootCmd = &cobra.Command{
		Use:   "tpr",
		Short: "TPR: Terraform private registry.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, setLogger)

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Set logger in debug")
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file")
}

// Setting up logger
func setLogger() {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")

	if config != "" {
		viper.SetConfigFile(config)
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Error to reading config file (%s): %s\n", config, errors.WithStack(err))
		os.Exit(1)
	}

	viper.SetEnvPrefix("TPR")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

}

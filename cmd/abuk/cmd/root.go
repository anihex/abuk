/*
Copyright © 2020 Chris <chris@suletuxe.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/anihex/abuk/cli"
	"github.com/anihex/abuk/internal"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abuk",
	Short: "A player for audiobooks",
	Long: `A player for audiobooks that runs on the CLI.
	
	It uses ffmpeg, ffprobe, ffplay and rcrack to play the audiobooks.
	Each program can be configured using environment variable, a config
	file or by using parameters on the CLI.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		config := &internal.Config{}
		app := internal.NewAbuk(*config)
		cli.CLI(app)
	},

	Version: "v1.0.0",
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.abuk.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("file", "f", "", "Path to a single file to play (overrides 'directory')")
	rootCmd.Flags().StringP("directory", "d", "", "Path to the audiobook library")
	rootCmd.Flags().String("ffmpeg", "", "Path to the ffmpeg binary")
	rootCmd.Flags().String("ffprobe", "", "Path to the ffprobe binary")
	rootCmd.Flags().String("ffplay", "", "Path to the ffplay binary")
	rootCmd.Flags().String("rcrack", "", "Path to the rcrack folder")
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

		// Search config in home directory with name ".abuk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".abuk")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

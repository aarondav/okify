/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"os"
	"github.com/briandowns/spinner"
	"time"
	"math/rand"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "okify",
	Short: "Feelings are more important than production.",
	Long: `Avoid getting non-positive feedback from anything.

Example:
	ls nonexistent-file || okify
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
			s.Suffix = " Calculating non-offensive response"
			s.FinalMSG = randomCompliment()
			s.Start()
			time.Sleep(4 * time.Second)
			s.Stop()
			os.Exit(0)
		},
}

func randomCompliment() string {
	rand.Seed(time.Now().Unix())
	compliments := []string{
		"Everything is fine!\n",
		"Ignore the haters!\n",
		"Looks good to me!\n",
		"You are doing such a good job!\n",
		"No one is as good as you are!\n",
		"How is that you are still single!\n",
		"You are so handsome!\n",
	}
	complimentIndex := rand.Intn(len(compliments))
	return compliments[complimentIndex]
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("You are doing great")
		os.Exit(0)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.okify.yaml)")
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
			fmt.Println("You are a great person!")
			os.Exit(0)
		}

		// Search config in home directory with name ".okify" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".okify")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
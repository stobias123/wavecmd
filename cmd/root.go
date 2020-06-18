// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spaceapegames/go-wavefront"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var wavefrontToken string
var wavefrontAddress string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wavecmd",
	Short: "A quick cli for posting data to wavefront",
	Long: `This is a cli utility packaged as a single binary to allow for quick wavefront utilities.. For example:

Post events. `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	//rootCmd.SetEnvPrefix("wavefront")

	rootCmd.PersistentFlags().StringVar(&wavefrontToken, "token", "", "Wavefront Token")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))

	rootCmd.PersistentFlags().StringVar(&wavefrontAddress, "address", "", "Wavefront Address")
	viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wavecmd.yaml)")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".wavecmd" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wavecmd")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Debug("Using config file:", viper.ConfigFileUsed())
	}

	wavefrontToken = viper.GetString("token")
	wavefrontAddress = viper.GetString("address")
}

// GetClient returns a wavefront cli client from the parsed config.
func GetClient(cmd *cobra.Command, args []string) *wavefront.Client {
	token, err := rootCmd.PersistentFlags().GetString("token")
	if err != nil {
		log.Fatal(err)
	}
	address, err := rootCmd.PersistentFlags().GetString("address")
	if err != nil {
		log.Fatal(err)
	}

	config := &wavefront.Config{
		Address: address,
		Token:   token,
	}

	client, err := wavefront.NewClient(config)
	return client
}

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
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spaceapegames/go-wavefront"
	"github.com/spf13/cobra"
)

// eventCmd represents the event command
var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var createEventCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an event",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient(cmd, args)
		events := client.Events()

		eventName, _ := cmd.Flags().GetString("name")
		eventType, _ := cmd.Flags().GetString("type")
		eventDetails, _ := cmd.Flags().GetString("message")
		eventTags, _ := cmd.Flags().GetStringArray("tags")

		e := &wavefront.Event{
			Name:    eventName,
			Type:    eventType,
			Details: eventDetails,
			Tags:    eventTags,
		}

		err := events.Create(e)
		if err != nil {
			log.Fatal(err)
		}

		res, err := json.MarshalIndent(e, "", "  ")
		fmt.Println(string(res))
	},
}

var searchEventCmd = &cobra.Command{
	Use:   "search",
	Short: "Search an event by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient(cmd, args)

		id, err := cmd.Flags().GetString("id")

		if err != nil {
			log.Fatal(err)
		}

		events := client.Events()

		e, err := events.FindByID(id)
		if e == nil {
			log.Fatal("Nothing found")
		}
		res, err := json.MarshalIndent(e, "", "  ")
		fmt.Println(string(res))
	},
}

var closeEventCmd = &cobra.Command{
	Use:   "close",
	Short: "Close an event by ID",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetClient(cmd, args)
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}

		events := client.Events()

		e, err := events.FindByID(id)
		if e == nil {
			log.Fatal("Event not found")
		}

		err = events.Close(e)
		if err != nil {
			log.Fatal(err)
		}

		res, err := json.MarshalIndent(e, "", "  ")
		fmt.Println(string(res))
	},
}

func init() {
	searchEventCmd.Flags().String("id", "", "Id to search for.")
	searchEventCmd.MarkFlagRequired("id")

	closeEventCmd.Flags().String("id", "", "Id to search for.")
	closeEventCmd.MarkFlagRequired("id")

	createEventCmd.Flags().StringP("name", "n", "", "Name of the event")
	createEventCmd.Flags().StringP("message", "m", "", "Message to include with the event")
	createEventCmd.Flags().StringP("type", "t", "", "Type of the event. Defaults to cli.")
	createEventCmd.Flags().StringArray("tags", []string{"cli"}, "Tags to include.")
	createEventCmd.MarkFlagRequired("name")
	createEventCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(eventCmd)
	eventCmd.AddCommand(createEventCmd)
	eventCmd.AddCommand(searchEventCmd)
	eventCmd.AddCommand(closeEventCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// eventCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/radovskyb/watcher"
	"github.com/spf13/cobra"
)

var importerSource string
var importerDestination string
var lookupDelay int

// importCmd represents the import command
var importerCmd = &cobra.Command{
	Use:   "importer",
	Short: "manager the importers and trigger import",
	Long: `miam use importer specified in the configuration.
	This command can list them but also trigger import to one of all of them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("import called")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a new importer",
	Long:  "Start a new importer from the desired type. The importer will remain active in the background until stoped",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start` called")

		w := watcher.New()

		// SetMaxEvents to 1 to allow at most 1 event's to be received
		// on the Event channel per watching cycle.
		//
		// If SetMaxEvents is not set, the default is to send all events.
		w.SetMaxEvents(1)
		// Only notify rename and move events.
		w.FilterOps(watcher.Write)

		go func() {
			for {
				select {
				case event := <-w.Event:
					fmt.Println(event) // Print the event's info.
				case err := <-w.Error:
					log.Fatalln(err)
				case <-w.Closed:
					return
				}
			}
		}()

		// Watch this folder for changes.
		if err := w.AddRecursive(importerSource); err != nil {
			log.Fatalln(err)
		}

	},
}

func initStartCmd() {
	startCmd.Flags().StringVarP(&importerSource, "source", "s", ".", "The source folder that container images to import")
	startCmd.Flags().StringVarP(&importerSource, "destination", "d", "", "The destination folder to copy images")
	startCmd.Flags().IntVarP(&lookupDelay, "lookup delay", "t", 10, "When importer is not treating files, it will wait this delay before checking the source folder for new files")
	importerCmd.AddCommand(startCmd)
}

func init() {
	initStartCmd()
	rootCmd.AddCommand(importerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

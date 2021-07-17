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

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yarmand/miam/importer"
)

var importerSource string
var importerDestination string

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
		i := importer.New(
			afero.NewOsFs(),
			importerSource,
			importer.PLogFilename(fmt.Sprintf("Processing file: %s/", importerSource), "")).
			Then(importer.PMoveToDateFolder(importerDestination)).
			Then(importer.PLogFilename("", "Processed"))
		err := i.Import()
		if err != nil {
			fmt.Printf("Error in Importer:%v\n", err)
		}
	},
}

func initStartCmd() {
	startCmd.Flags().StringVarP(&importerSource, "source", "s", ".", "The source folder that container images to import")
	startCmd.Flags().StringVarP(&importerDestination, "destination", "d", "", "The destination folder to copy images")
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

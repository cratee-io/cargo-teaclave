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

	"github.com/sammyne/cargo-teaclave/tools/testings"
	"github.com/spf13/cobra"
)

var (
	testCmdCratePath string
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, err := testings.NewWorkspace(testCmdCratePath)
		if err != nil {
			panic(err)
		}
		//defer os.RemoveAll(workingDir)

		if err := testings.Run(workingDir); err != nil {
			panic(err)
		}

		fmt.Println(workingDir)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.
	testCmd.Flags().StringVar(&testCmdCratePath, "crate", "", "path of crate to test")

	if err := testCmd.MarkFlagRequired("crate"); err != nil {
		panic(err)
	}
}

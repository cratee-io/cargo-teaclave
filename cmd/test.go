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
	"os"

	"github.com/sammyne/cargo-teaclave/tools/testings"
	"github.com/spf13/cobra"
)

var (
	testCmdCratePath string
	testCmdDriverTag string
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test a given teaclave-sgx-sdk-ported crate",
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, err := testings.NewWorkspace(testCmdCratePath, testCmdDriverTag)
		if err != nil {
			panic(err)
		}
		defer os.RemoveAll(workingDir)

		if err := testings.Run(workingDir); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.
	workingDir, _ := os.Getwd()
	testCmd.Flags().StringVar(&testCmdCratePath, "crate", workingDir,
		"path of the crate to test")

	testCmd.Flags().StringVar(&testCmdDriverTag, "driver", "v1.1.2",
		"tag of the driver to use, e.g. v1.1.2")
}

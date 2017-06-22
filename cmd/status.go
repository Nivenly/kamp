// Copyright © 2017 The Kamp Authors
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
	"github.com/spf13/cobra"
	"os"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Return status of kamp container(s)",
	Long:  KampBannerMessage("Return pod information and metrics about your kamp instance."),
	Run: func(cmd *cobra.Command, args []string) {
		err := RunStatus(statusOpt)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}

type StatusOptions struct {
	Options
}

var statusOpt = &StatusOptions{}

func RunStatus(options *StatusOptions) error {

	// Todo (@Grillz) start coding here

	return nil
}

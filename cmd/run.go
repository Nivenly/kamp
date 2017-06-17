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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a kamp instance in a Kubernetes cluster",
	Long:  KampBannerMessage("Run and attach to an arbitrary container in a Kubernetes cluster as a pod."),
	Run: func(cmd *cobra.Command, args []string) {
		err := RunRun(runOpt)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&runOpt.Command, "cmd", "c", "/bin/bash", "The command to execute in the container.")
	runCmd.Flags().StringVarP(&runOpt.KubernetesNamespace, "namespace", "n", "default", "The Kubernetes namespace to run the container in.")
	runCmd.Flags().StringVarP(&runOpt.Volume, "volume", "V", "", "The volume string to mount CIFS volumes with. <local>:<remote>")
	runCmd.SetUsageTemplate(UsageTemplate)
	args := os.Args
	if len(args) < 3 {
		runCmd.Help()
		os.Exit(0)
	}
}

type RunOptions struct {
	Options
	ImageQuery          string
	Command             string
	KubernetesNamespace string
	Volume              string
}

var runOpt = &RunOptions{}

func RunRun(options *RunOptions) error {

	// Todo (@Grillz) start coding here

	return nil
}

const UsageTemplate = `Usage:{{if .Runnable}}
  {{if .HasAvailableFlags}}{{appendIfNotPresent .UseLine "<image>:(tag) [flags]"}}{{else}}{{.UseLine}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
  {{ .CommandPath}} [command]{{end}}{{if gt .Aliases 0}}
Aliases:
  {{.NameAndAliases}}
{{end}}{{if .HasExample}}
Examples:
{{ .Example }}{{end}}{{ if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableLocalFlags}}
Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{end}}{{ if .HasAvailableInheritedFlags}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsHelpCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{ if .HasAvailableSubCommands }}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

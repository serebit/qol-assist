// Copyright © 2020-2020 Solus Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
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
)

const QolAssistVersion = "0.5.0"

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "show version",
	Long: "Print the qol-assist version and exit",
	Run: printVersion,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func printVersion(_ *cobra.Command, _ []string) {
	fmt.Printf("qol-assist version %v\n\nCopyright © 2017-2020 Solus Project\n", QolAssistVersion)
	fmt.Println("Licensed under the Apache License, Version 2.0")
}

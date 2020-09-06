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

import "github.com/spf13/cobra"

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Perform migration functions",
	Long: "Applies migrations that are available on the system",
	Aliases: []string{"m"},
	Run: migrate,
}

func init() {
	RootCmd.AddCommand(listUsersCmd)
}

func migrate(_ *cobra.Command, _ []string) {

}

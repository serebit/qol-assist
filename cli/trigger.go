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

package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/getsolus/qol-assist/migration"
	"os"
)

var trigger = &cmd.CMD{
	Name:  "trigger",
	Short: "Schedule migration on next boot",
	Alias: "t",
	Args:  &struct{}{},
	Run: func(_ *cmd.RootCMD, _ *cmd.CMD) {
		if os.Geteuid() != 0 || os.Getegid() != 0 {
			fmt.Println("This command must be run with root privileges.")
			return
		}

		if err := migration.CreateTriggerFile(); err != nil {
			fmt.Printf("Failed to create trigger file %s: %s\n", migration.TriggerFile, err)
			return
		}

		fmt.Println("Migration will run on next boot.")
	},
}

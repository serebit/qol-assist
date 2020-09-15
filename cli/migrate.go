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
	"github.com/getsolus/qol-assist/core"
	"os"
)

var migrate = &cmd.CMD{
	Name:  "migrate",
	Short: "Applies migrations that are available on the system",
	Alias: "m",
	Args:  &struct{}{},
	Run: func(_ *cmd.RootCMD, _ *cmd.CMD) {
		if os.Geteuid() != 0 || os.Getegid() != 0 {
			fmt.Println("This command must be run with root privileges.")
			return
		}

		if !core.TriggerFileExists() {
			fmt.Println("Refusing to run migration without trigger file.")
			return
		}

		var _, err = core.NewContext()
		if err != nil {
			fmt.Printf("Unable to gather system info: %s\n", err)
		}
	},
}

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

import "github.com/DataDrake/cli-ng/cmd"

var RootCMD = &cmd.RootCMD{
	Name:  "qol-assist",
	Short: "QoL assistance to help Solus roll!",
}

func init() {
	RootCMD.RegisterCMD(&cmd.Help)
	RootCMD.RegisterCMD(triggerCMD)
	RootCMD.RegisterCMD(versionCMD)
	RootCMD.RegisterCMD(migrateCMD)
	RootCMD.RegisterCMD(listUsersCMD)
}

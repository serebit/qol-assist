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

package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Migration struct {
	Name string
	Path string

	Description string        `toml:"description"`
	UpdateUsers []UpdateUsers `toml:"users-update"`
	UpdateGroup []UpdateGroup `toml:"group-update"`
}

type UpdateUsers struct {
	UserFilters []string `toml:"only"`
	GroupName   string   `toml:"group"`
}

type UpdateGroup struct {
	GroupName  string `toml:"name"`
	NewGroupID int32  `toml:"id"`
}

// Load reads a Migration configuration from a file and parses it
func (m *Migration) Load(path string) error {
	// Read the configuration into the program
	var cfg, err = readFile(path)
	if err != nil {
		return fmt.Errorf("unable to read config file located at %s", path)
	}

	// Save the configuration into the content structure
	if err := toml.Unmarshal(cfg, m); err != nil {
		return fmt.Errorf("unable to read config file located at %s due to %s", path, err.Error())
	}
	return nil
}

func (m *Migration) Validate() error {
	// Verify that there is at least one binary to execute, otherwise there
	// is no need to continue
	if len(m.UpdateUsers) == 0 && len(m.UpdateGroup) == 0 {
		return fmt.Errorf("migrations must contain at least one modification")
	}
	return nil
}

func (m *Migration) Run(context *Context) {
	for _, task := range m.UpdateUsers {
		var filtered = context.FilterUsers(task.UserFilters...)

		for _, user := range filtered {
			var _, err = context.AddToGroup(user, task.GroupName)
			if err != nil {
				fmt.Printf("Failed to run migration %s due to error: %s", m.Name, err)
			}
		}
	}
}

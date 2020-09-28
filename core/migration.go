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
	"io/ioutil"
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

func LoadMigrations() []Migration {
	var allMigrations = make([]Migration, 0)

	if sysFiles, err := ioutil.ReadDir(SysDir); err != nil {
		fmt.Printf("System directory for migrations at %s is unreadable, skipping...\n", SysDir)
	} else {
		for _, it := range sysFiles {
			allMigrations = appendMigrationFrom(allMigrations, SysDir, it.Name())
		}
	}

	if usrFiles, err := ioutil.ReadDir(UsrDir); err != nil {
		fmt.Printf("User directory for migrations at %s is unreadable, skipping...\n", UsrDir)
	} else {
		for _, it := range usrFiles {
			allMigrations = appendMigrationFrom(allMigrations, UsrDir, it.Name())
		}
	}

	return allMigrations
}

func appendMigrationFrom(migrations []Migration, dir string, name string) []Migration {
	if migration, err := parseMigration(dir, name); err != nil {
		fmt.Println(err)
	} else {
		migrations = append(migrations, migration)
	}
	return migrations
}

func parseMigration(dir string, name string) (Migration, error) {
	var migration Migration

	// Read the configuration into the program
	var path = fmt.Sprintf("%s/%s", dir, name)
	var cfg, err = readFile(path)
	if err != nil {
		return migration, fmt.Errorf("unable to read migration file located at %s due to %s", path, err)
	}

	// Save the configuration into the content structure
	if err := toml.Unmarshal(cfg, &migration); err != nil {
		return migration, fmt.Errorf("unable to parse migration file located at %s due to %s", path, err.Error())
	}

	migration.Name = name
	migration.Path = path

	return migration, nil
}

func (m Migration) Validate() error {
	if len(m.UpdateUsers) == 0 && len(m.UpdateGroup) == 0 {
		return fmt.Errorf("migrations must contain at least one modification")
	}
	return nil
}

func (m Migration) Run(context *Context) {
	fmt.Printf("Running migration %s\n", m.Name)
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

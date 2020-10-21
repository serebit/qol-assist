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
	"github.com/DataDrake/waterlog"
	"io/ioutil"
	gouser "os/user"
	"strconv"
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
	NewGroupID int    `toml:"id"`
}

func LoadMigrations() []Migration {
	var allMigrations = make([]Migration, 0)

	waterlog.Debugln("Loading migrations...")

	if sysFiles, err := ioutil.ReadDir(SysDir); err != nil {
		waterlog.Debugf("System directory for migrations at %s is unreadable, skipping...\n", SysDir)
	} else {
		for _, it := range sysFiles {
			allMigrations = appendMigrationFrom(allMigrations, SysDir, it.Name())
		}
	}

	if usrFiles, err := ioutil.ReadDir(UsrDir); err != nil {
		waterlog.Debugf("User directory for migrations at %s is unreadable, skipping...\n", UsrDir)
	} else {
		for _, it := range usrFiles {
			allMigrations = appendMigrationFrom(allMigrations, UsrDir, it.Name())
		}
	}

	waterlog.Debugln("Finished loading migrations.")

	return allMigrations
}

func appendMigrationFrom(migrations []Migration, dir string, name string) []Migration {
	if migration, err := parseMigration(dir, name); err != nil {
		waterlog.Warnf("Failed to parse migration %s in dir %s: %s", name, dir, err)
	} else {
		migrations = append(migrations, migration)
		waterlog.Debugf("    Loaded migration %s from directory %s\n", name, dir)
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
		return migration, fmt.Errorf("unable to parse migration file located at %s due to %s", path, err)
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
	waterlog.Debugf("Running migration %s...\n", m.Name)
	for _, task := range m.UpdateUsers {
		m.updateUsers(context, task)
	}
	for _, task := range m.UpdateGroup {
		m.updateGroup(context, task)
	}
}

func (m Migration) updateUsers(context *Context, task UpdateUsers) {
	var filtered = context.FilterUsers(task.UserFilters...)

	for _, user := range filtered {
		if ran, err := context.AddToGroup(user, task.GroupName); err != nil {
			waterlog.Warnf("    failed to add group %s to user %s due to error: %s\n", task.GroupName, user.Name, err)
		} else if ran {
			waterlog.Debugf("    successfully added group %s to user %s\n", task.GroupName, user.Name)
		} else {
			waterlog.Debugf("    user %s already has group %s, skipping\n", user.Name, task.GroupName)
		}
	}
}

func (m Migration) updateGroup(context *Context, task UpdateGroup) {
	var byName *gouser.Group = nil
	var byID *gouser.Group = nil

	var gid = strconv.Itoa(task.NewGroupID)
	for _, group := range context.groups {
		switch {
		case group.Name == task.GroupName:
			byName = &group
		case group.Gid == strconv.Itoa(task.NewGroupID):
			byID = &group
		}
	}

	if byName == nil && byID == nil {
		// group doesn't exist, create it
		if err := context.CreateGroup(task.GroupName, gid); err != nil {
			waterlog.Warnf("    failed to create group with name %s and GID %s due to error %s\n", task.GroupName, gid, err)
		} else {
			waterlog.Debugf("    successfully created group %s with GID %s\n", task.GroupName, gid)
		}
	} else if byName != nil && byID == nil {
		// group has wrong ID, fix it
		if err := context.UpdateGroupID(task.GroupName, gid); err != nil {
			waterlog.Warnf("    failed to update group with name %s to new GID %s due to error %s\n", task.GroupName, gid, err)
		} else {
			waterlog.Debugf("    successfully updated group with name %s to new GID %s\n", task.GroupName, gid)
		}
	} else if byName != byID {
		// there's a group with our desired ID, and it isn't supposed to have it. Fail.
		waterlog.Warnf("    another group already exists with desired GID %s, skipping update for group %s\n", gid, task.GroupName)
	}
}

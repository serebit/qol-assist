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

// #include <pwd.h>
// #include <unistd.h>
import "C"
import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	users "os/user"
	"strconv"
	"strings"
)

const minimumUID = 1000
const wheelGroup = "sudo"

var listUsers = &cmd.CMD{
	Name:  "list-users",
	Short: "List users on the system and their associated groups",
	Alias: "l",
	Args:  &ListArgs{},
	Run: func(_ *cmd.RootCMD, command *cmd.CMD) {
		var filter = command.Args.(*ListArgs).Filter

		for _, it := range obtainAllUsers() {
			switch {
			case filter == "all":
				fallthrough
			case filter == "active" && it.isActive:
				fallthrough
			case filter == "system" && !it.isActive:
				fallthrough
			case filter == "admin" && (it.isRoot || it.isAdmin):
				fmt.Printf("User: %s (%s)\n", it.name, strings.Join(it.groups, ":"))
			}
		}
	},
}

type ListArgs struct {
	Filter string `desc:"[system|all|admin|active]"`
}

type user struct {
	name     string
	groups   []string
	isActive bool
	isRoot   bool
	isAdmin  bool
}

func obtainAllUsers() []user {
	var shells = activeShells()
	var allUsers = make([]user, 0)

	C.setpwent()
	for {
		var pwd = C.getpwent()
		if pwd == nil {
			break
		}

		var uid = int(pwd.pw_uid)
		var it, err = users.LookupId(strconv.Itoa(uid))
		if err != nil {
			break
		}

		var groupIDs []string
		if groupIDs, err = it.GroupIds(); err != nil {
			break
		}

		var groupNames = groupNamesFromGUIDs(groupIDs)
		allUsers = append(allUsers, user{
			name:     C.GoString(pwd.pw_name),
			groups:   groupNames,
			isActive: uid >= minimumUID && contains(shells, C.GoString(pwd.pw_shell)),
			isRoot:   uid == 0 && int(pwd.pw_gid) == 0,
			isAdmin:  contains(groupNames, wheelGroup),
		})
	}
	C.endpwent()

	return allUsers
}

func groupNamesFromGUIDs(guids []string) []string {
	var groupNames = make([]string, 0)
	for i := range guids {
		var group, err = users.LookupGroupId(guids[i])
		if err == nil {
			groupNames = append(groupNames, group.Name)
		}
	}
	return groupNames
}

func activeShells() []string {
	var shells = make([]string, 0)

	C.setusershell()
	for {
		var cShell = C.getusershell()
		if cShell == nil {
			break
		}
		shells = append(shells, C.GoString(cShell))
	}
	C.endusershell()

	return shells
}

func contains(list []string, item string) bool {
	for _, it := range list {
		if it == item {
			return true
		}
	}
	return false
}

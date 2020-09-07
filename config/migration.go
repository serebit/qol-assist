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

package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const TrackDir = "/var/lib/qol-assist"
const TriggerFile = TrackDir + "/trigger"
const StatusFile = TrackDir + "/status"

type Migration struct {
	Name string
	Path string

	Description     string            `toml:"description"`
	AddUsersToGroup []AddUsersToGroup `toml:"add-users-to-group"`
	UpdateGroupID   []UpdateGroupID   `toml:"update-group-id"`
}

type AddUsersToGroup struct {
	UserFilters []string `toml:"user-filters"`
	GroupName   string   `toml:"group-name"`
}

type UpdateGroupID struct {
	GroupName  string `toml:"group-name"`
	NewGroupID int32  `toml:"new-group-id"`
}

// Load reads a Migration configuration from a file and parses it
func (m *Migration) Load(path string) error {
	// Check if this is a valid file path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	// Read the configuration into the program
	cfg, err := ioutil.ReadFile(filepath.Clean(path))
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
	if len(m.AddUsersToGroup) == 0 && len(m.UpdateGroupID) == 0 {
		return fmt.Errorf("migrations must contain at least one modification")
	}
	return nil
}

// Grab the current migration level from the status file.
// If the file doesn't exist, or we have parse issues, we default to 0.
func GetMigrationLevel() int {
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(StatusFile); err != nil {
		// file doesn't exist, default
		return 0
	}

	if level, err := strconv.Atoi(string(bytes)); err != nil || level < 0 {
		// parse error, default
		return 0
	} else {
		return level
	}
}

func ConstructTrackDir() error {
	if _, err := os.Stat(TrackDir); os.IsNotExist(err) {
		if err := os.Mkdir(TrackDir, 0o755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func CreateTriggerFile() error {
	if err := ConstructTrackDir(); err != nil {
		return err
	}

	if _, err := os.Stat(TriggerFile); os.IsNotExist(err) {
		if _, err := os.Create(TriggerFile); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

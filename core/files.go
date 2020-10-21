// Copyright © 2020 Solus Project
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
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	TrackDir    string
	SysDir      string
	UsrDir      string
	TriggerFile = TrackDir + "/trigger"
)

func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func createFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	// Check if this is a valid file path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []byte{}, err
	}

	// Read the file as bytes, returning both bytes and a possible error
	return ioutil.ReadFile(filepath.Clean(path))
}

func ConstructTrackDir() error {
	return createDir(TrackDir)
}

func CreateTriggerFile() error {
	if err := ConstructTrackDir(); err != nil {
		return err
	}
	return createFile(TriggerFile)
}

func RemoveTriggerFile() error {
	if _, err := os.Stat(TriggerFile); err == nil {
		if err := os.Remove(TriggerFile); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func TriggerFileExists() bool {
	var _, err = os.Stat(TriggerFile)
	return err == nil
}

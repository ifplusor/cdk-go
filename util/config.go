// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParseConfig(file string, v interface{}) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	if strings.HasSuffix(file, "json") {
		err = json.Unmarshal(data, v)
	} else {
		err = yaml.Unmarshal(data, v)
	}
	return err
}

// Copyright 2023 Linkall Inc.
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

package internal

import (
	"fmt"
	"strings"

	cdkgo "github.com/vanus-labs/cdk-go"
)

var _ cdkgo.SourceConfigAccessor = &exampleConfig{}

type exampleConfig struct {
	cdkgo.SourceConfig `json:",inline" yaml:",inline"`
	Source             string `json:"source" yaml:"source" validate:"required"`
	Secret             Secret `json:"secret" yaml:"secret"`
}

func NewExampleConfig() cdkgo.SourceConfigAccessor {
	return &exampleConfig{}
}

func (c *exampleConfig) GetSecret() cdkgo.SecretAccessor {
	return &c.Secret
}

func (c *exampleConfig) Validate() error {
	if !strings.HasPrefix(c.Source, "example") {
		return fmt.Errorf("source is invlaid")
	}
	return c.SourceConfig.Validate()
}

type Secret struct {
	Host     string `json:"host" yaml:"host" validate:"required"`
	Username string `json:"username" yaml:"username" validate:"required"`
	Password string `json:"password" yaml:"password" validate:"required"`
}

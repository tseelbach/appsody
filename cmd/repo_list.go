// Copyright © 2019 IBM Corporation and others.
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

package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type repoListCommandConfig struct {
	*RootCommandConfig
	output string
}

func newRepoListCmd(config *RootCommandConfig) *cobra.Command {
	repoListConfig := &repoListCommandConfig{RootCommandConfig: config}
	// repo list represent repo list cmd
	var repoListCmd = &cobra.Command{
		Use:   "list",
		Short: "List configured Appsody repositories",
		Long:  `List configured Appsody repositories. An asterisk denotes the default repository.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var repos RepositoryFile

			list, repoErr := repos.getRepos(config)
			if repoErr != nil {
				return repoErr
			}

			if repoListConfig.output == "" {
				repoList, err := repos.listRepos(config)
				if err != nil {
					return err
				}
				Info.log("\n", repoList)
			} else if repoListConfig.output == "yaml" {
				bytes, err := yaml.Marshal(&list)
				if err != nil {
					return err
				}
				result := string(bytes)
				Info.log("\n", result)
			} else if repoListConfig.output == "json" {
				bytes, err := json.Marshal(&list)
				if err != nil {
					return err
				}
				result := string(bytes)
				Info.log("\n", result)
			}

			return nil
		},
	}

	repoListCmd.PersistentFlags().StringVarP(&repoListConfig.output, "output", "o", "", "Output repo list in yaml or json format")
	return repoListCmd
}

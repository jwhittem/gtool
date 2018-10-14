// Copyright © 2018 Jeremy Whittemore <kbfastcat64@gmail.com>
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
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	mode = 0755
)

// cloneCmd represents the cloneall command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone all repos for a user or organization",

	Run: func(cmd *cobra.Command, args []string) {
		client := NewGithubClient(token)
		if client == nil {
			log.Println("Can't connect to github API.")
			os.Exit(1)
		}

		// iterate through args, once for each organization or user specified
		for _, account := range args {
			allRepos := GetAllRepos(client, account)
			if allRepos == nil {
				log.Println("No repositories found...")
				return
			}

			log.Println("Cloning all repos for:", account)

			if err := os.Mkdir(path+account, mode); err != nil {
				log.Println(err)
			}

			// iterate through all repos for the arg
			for _, repo := range allRepos {
				fullPath := path + account + "/" + repo.GetName()

				_, err := os.Stat(fullPath)
				if os.IsNotExist(err) {
					cmd := exec.Command("git", "clone", repo.GetGitURL(), fullPath)
					if output, err := cmd.CombinedOutput(); err != nil {
						log.Println("error: ", err)
					} else {
						log.Print(string(output))
					}
				}
			}
		}
	},
}

func init() {
	repoCmd.AddCommand(cloneCmd)

}

// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"
	"strings"

	"github.com/iancmcc/jig/fs"
	"github.com/iancmcc/jig/match"
	"github.com/spf13/cobra"
)

var (
	exact bool
	limit int
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List repositories",
	Long:  `List repositories below the current directory, optionally sorted by similarity to a search string`,
	Run: func(cmd *cobra.Command, args []string) {
		here, _ := os.Getwd()
		here = strings.TrimSuffix(here, "/") + "/"
		repos := fs.DefaultFinder().FindBelowWithChildrenNamed(here, ".git", 1)
		if len(args) == 0 {
			var i int
			for repo := range repos {
				if limit > 0 && i >= limit {
					break
				}
				fmt.Println(strings.TrimPrefix(repo, here))
				i++
			}
			return
		}
		matcher := match.DefaultMatcher(args[0])
		for repo := range repos {
			matcher.Add(strings.TrimPrefix(repo, here))
		}
		for i, repo := range matcher.Match() {
			if limit > 0 && i >= limit {
				break
			}
			fmt.Println(repo)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
	lsCmd.PersistentFlags().BoolVarP(&exact, "exact", "x", false, "Return exact matches only (default is fuzzy matching)")
	lsCmd.PersistentFlags().IntVarP(&limit, "limit", "n", 0, "Limit the number of results returned (default is no limit)")
}

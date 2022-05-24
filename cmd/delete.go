/*
Copyright © 2022 yuyicai

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yuyicai/gh-cleaner/pkg/github"
)

var (
	repositories []string
	packages     []string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete github repositories and packages",
	Long:  `delete github repositories and packages`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteRun()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	pf := deleteCmd.Flags()
	pf.StringSliceVar(&repositories, "repos", repositories, "Github repos, example: octocat/Hello-World")
	pf.StringSliceVar(&packages, "packages", packages, "Github container packages, example: octocat/Hello-World")
}

func deleteRun() error {
	reposMap := make(map[string]string)
	packagesMap := make(map[string]string)
	fmt.Println("repositories:", repositories)
	for _, repository := range repositories {
		repository = strings.Replace(repository, " ", "", -1)
		r := strings.Split(repository, "/")
		if len(r) != 2 || r[0] == "" || r[1] == "" {
			return fmt.Errorf("repository format error: %s", repository)
		}
		reposMap[r[0]] = r[1]
	}

	for _, pkg := range packages {
		pkg = strings.Replace(pkg, " ", "", -1)
		p := strings.Split(pkg, "/")
		if len(p) != 2 || p[0] == "" || p[1] == "" {
			return fmt.Errorf("package format error: %s", pkg)
		}
		packagesMap[p[0]] = p[1]
	}

	client := github.NewAuthClient(context.Background(), token)
	if err := client.DeleteContainerPackages(packagesMap); err != nil {
		return err
	}

	return client.DeleteRepositories(reposMap)
}

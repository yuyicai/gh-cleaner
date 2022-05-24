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
	"errors"

	"github.com/spf13/cobra"
	"github.com/yuyicai/gh-cleaner/pkg/github"
)

var (
	repositories []string
	packages     []string
	owner        string
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
	pf.StringSliceVarP(&repositories, "repos", "r", repositories, "Github repos list")
	pf.StringSliceVarP(&packages, "packages", "p", packages, "Github container packages list")
	pf.StringVarP(&owner, "username", "u", owner, "Github user or org name")
	deleteCmd.MarkFlagRequired("username")
}

func deleteRun() error {
	if token == "" {
		return errors.New("please set GitHub PAT token by --token or set env variable GITHUB_TOKEN")
	}
	client := github.NewAuthClient(context.Background(), token)
	if err := client.DeleteContainerPackages(owner, packages); err != nil {
		return err
	}

	return client.DeleteRepositories(owner, repositories)
	return nil
}

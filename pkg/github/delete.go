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
package github

import (
	"fmt"
	"log"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func (c *AuthClient) DeleteRepositories(repos map[string]string) error {
	for owner, repo := range repos {
		if owner == "" || repo == "" {
			return fmt.Errorf("repository format error: %s/%s", owner, repo)
		}
		log.Printf("deleting repository %s/%s", owner, repo)
		if err := wait.PollImmediate(20*time.Second, 1*time.Minute, func() (bool, error) {
			if _, err := c.Repositories.Delete(c.ctx, owner, repo); err != nil {
				if resourceNotFound(err) {
					return false, err
				}
				log.Printf("error deleting repository %s/%s: %v", owner, repo, err)
				log.Println("retrying...")
				return false, nil
			}
			return true, nil
		}); err != nil {
			if resourceNotFound(err) {
				log.Printf("repository %s/%s not found, skip", owner, repo)
				return nil
			}
		}

		log.Printf("deleted repository %s/%s ✅", owner, repo)
	}
	return nil
}

func (c *AuthClient) DeleteContainerPackages(packages map[string]string) error {
	for owner, pkg := range packages {
		if owner == "" || pkg == "" {
			return fmt.Errorf("package format error: %s/%s", owner, pkg)
		}
		log.Printf("deleting container package %s/%s", owner, pkg)
		if err := wait.PollImmediate(20*time.Second, 1*time.Minute, func() (bool, error) {
			if _, err := c.Users.DeletePackage(c.ctx, owner, "container", pkg); err != nil {
				if resourceNotFound(err) {
					return false, err
				}
				log.Printf("error deleting container package %s/%s: %v", owner, pkg, err)
				log.Println("retrying...")
				return false, nil
			}
			return true, nil
		}); err != nil {
			if resourceNotFound(err) {
				log.Printf("package %s/%s not found, skip", owner, pkg)
				return nil
			}
		}
		log.Printf("deleted container package %s/%s ✅", owner, pkg)
	}
	return nil
}

func resourceNotFound(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "not found")
}

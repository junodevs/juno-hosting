//+build mage

/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"runtime"

	"github.com/magefile/mage/sh"
)

// Build generates a binary of the project
func Build() error {
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	return sh.Run("go", "build", "--ldflags", "-s -w", "-o", "bin/hosting-server", "./")
}

// Format lints and fixes all files in the directory
func Format() error {
	return sh.Run("go", "fmt", "./...")
}

// Run builds a binary and executes it
func Run() error {
	if err := Build(); err != nil {
		return err
	}

	if runtime.GOOS == "windows" {
		return sh.RunV("./bin/hosting-server.exe")
	}

	return sh.RunV("./bin/hosting-server")
}

// Test executes all tests in the package
func Test() error {
	return sh.Run("go", "test", "./...")
}

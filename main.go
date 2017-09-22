// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"github.com/circonus-labs/circonus-agent/cmd"
	"github.com/circonus-labs/circonus-agent/internal/release"
)

func main() {
	cmd.Execute()
}

// for goreleaser, see .goreleaser.yml
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	tag     = ""
)

func init() {
	release.VERSION = version
	release.COMMIT = commit
	release.DATE = date
	release.TAG = tag
}
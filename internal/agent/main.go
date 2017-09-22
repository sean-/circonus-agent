// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package agent

import (
	"github.com/circonus-labs/circonus-agent/internal/plugins"
	"github.com/circonus-labs/circonus-agent/internal/reverse"
	"github.com/circonus-labs/circonus-agent/internal/server"
	"github.com/circonus-labs/circonus-agent/internal/statsd"
	"github.com/pkg/errors"
)

// Agent holds the main circonus-agent process
type Agent struct {
	errChan chan error
	plugins *plugins.Plugins
}

// New returns a new agent instance
func New() (*Agent, error) {
	a := Agent{
		errChan: make(chan error),
		plugins: plugins.New(),
	}

	if err := a.plugins.Scan(); err != nil {
		return nil, err
	}

	return &a, nil
}

// Start the agent
func (a *Agent) Start() {
	go func() {
		err := statsd.Start()
		if err != nil {
			a.errChan <- errors.Wrap(err, "Starting StatsD listener")
		}
	}()

	go func() {
		err := reverse.Start()
		if err != nil {
			a.errChan <- errors.Wrap(err, "Unable to start reverse connection")
		}
	}()

	go func() {
		err := server.Start()
		if err != nil {
			a.errChan <- errors.Wrap(err, "Starting server")
		}
	}()
}

// Stop the agent
func (a *Agent) Stop() {
	// noop
}

// Wait for agent components to exit
func (a *Agent) Wait() error {
	select {
	case err := <-a.errChan:
		return err
	}
}

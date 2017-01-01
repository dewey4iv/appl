package appl

import (
	"os"

	"git.dewey4.com/dewey4com/gop/version"
)

// New returns a new Application
func New(opts ...Option) (*Application, error) {
	var a Application

	for _, opt := range opts {
		if err := opt.Apply(&a); err != nil {
			return nil, err
		}
	}

	a.PID = os.Getegid()

	return &a, nil
}

// Application holds the various interface implementations
type Application struct {
	PID   int
	Name  string
	Build version.Build

	apps         []App
	startables   []Startable
	stoppables   []Stoppable
	killables    []Killable
	restartables []Restartable
}

// Start starts all apps and startables
func (a *Application) Start() error {
	for _, a := range a.startables {
		if err := a.Start(); err != nil {
			return err
		}
	}

	return nil
}

// Stop stops all apps and stopables
func (a *Application) Stop() error {
	for _, a := range a.stoppables {
		if err := a.Stop(); err != nil {
			return err
		}
	}

	return nil
}

// Kill kills all apps and killables
func (a *Application) Kill() error {
	for _, a := range a.killables {
		if err := a.Kill(); err != nil {
			return err
		}
	}

	return nil
}

// Restart restarts all apps and restartables
func (a *Application) Restart() error {
	for _, a := range a.restartables {
		if err := a.Restart(); err != nil {
			return err
		}
	}

	return nil
}

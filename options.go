package appl

import "git.dewey4.com/dewey4com/gop/version"

// Option defines something that can be applied to an Application
type Option interface {
	Apply(*Application) error
}

// WithApps takes interfaces and applies the to the Application
func WithApps(apps ...interface{}) Option {
	return &withApps{apps}
}

type withApps struct {
	apps []interface{}
}

func (opt *withApps) Apply(a *Application) error {

	for _, app := range opt.apps {
		if err := addInterface(a, app); err != nil {
			return err
		}
	}

	return nil
}

// WithBuild applies the built to the application
func WithBuild(build version.Build) Option {
	return &withBuild{build}
}

type withBuild struct {
	build version.Build
}

func (opt *withBuild) Apply(a *Application) error {
	a.Build = opt.build

	return nil
}

// WithName appliese the name to the application
func WithName(name string) Option {
	return &withName{name}
}

type withName struct {
	name string
}

func (opt *withName) Apply(a *Application) error {
	a.Name = opt.name

	return nil
}

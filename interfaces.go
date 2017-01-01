package appl

import "errors"

// App defines something that implements the other interfaces
type App interface {
	Startable
	Stoppable
	Killable
	Restartable
}

// Startable defines anything that can be started
type Startable interface {
	Start() error
}

// Stoppable defines something that can be stopped
type Stoppable interface {
	Stop() error
}

// Killable defines something that can be killed
type Killable interface {
	Kill() error
}

// Restartable defines something that can be restarted
type Restartable interface {
	Restart() error
}

func addInterface(a *Application, i interface{}) error {
	var casted bool

	if startable, castable := i.(Startable); castable {
		a.startables = append(a.startables, startable)
		casted = true
	}

	if stoppable, castable := i.(Stoppable); castable {
		a.stoppables = append(a.stoppables, stoppable)
		casted = true
	}

	if killable, castable := i.(Killable); castable {
		a.killables = append(a.killables, killable)
		casted = true
	}

	if restartable, castable := i.(Restartable); castable {
		a.restartables = append(a.restartables, restartable)
		casted = true
	}

	if casted {
		return nil
	}

	return errors.New("Value couldn't be casted")
}

package appl

// Start starts a startabke
func Start(s Startable) error {
	return s.Start()
}

// Stop starts a startabke
func Stop(s Stoppable) error {
	return s.Stop()
}

// Kill starts a startabke
func Kill(s Killable) error {
	return s.Kill()
}

// Restart starts a startabke
func Restart(s Restartable) error {
	return s.Restart()
}

// Wait just stops the rest of the app from running
func Wait() {
	select {}
}

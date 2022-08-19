package signals

import (
	"os"
	"os/signal"
)

var oneSignalHandler = make(chan struct{})

func SetupHandler() (stopCh <-chan struct{}) {
	close(oneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()
	return stop
}

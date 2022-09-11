package stop

import (
	"fmt"
	"os"
	"os/signal"
)

func StopBtInterrup() {
	fmt.Print("\nPress Ctrl + C to complete program")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		os.Exit(0)
	}()
	for {

	}
}

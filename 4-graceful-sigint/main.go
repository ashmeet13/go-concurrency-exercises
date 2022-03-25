//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Create a process
	proc := MockProcess{}

	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		flag := false
		for {
			<-interrupt
			fmt.Println()
			if flag {
				fmt.Println("Exit")
				os.Exit(1)
			}
			fmt.Println("Interrupted")
			fmt.Println("Press Ctrl-C one more time to exit")
			flag = true
			go proc.Stop()
		}
	}()

	// Run the process (blocking)
	proc.Run()
}

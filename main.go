package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// Log configuration.
	file, err := logConfig()
	if err != nil {
		panic(err)
	}
	defer func() {
		// Don't forget to close the file!
		// Changed from simple 'defer file.Close()' to this function to
		// avoid 'Unhandled error' warning
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// Context stuff.
	cancelCtx, cancel := context.WithCancel(context.Background())

	// Handling signals.
	signalsChan := make(chan os.Signal, 1)
	//signal.Notify(signalsChan, os.Interrupt, os.Kill)
	signal.Notify(signalsChan)

	// 'done' is used to stop the for loops user for each task.
	var done = false
	go func() {
		select {
		case s := <-signalsChan:
			done = true
			cancel()

			// Log that the SIGINT was received
			switch s {
			case os.Interrupt:
				log.Println("SIGINT received")

			case os.Kill:
				log.Println("SIGKILL received")

			default:
				log.Printf("Received unhandled signal: %v\n", s)
			}
		}
	}()

	for _, cron := range crons {
		wg.Add(1)

		go func(cron Cron) {
			defer wg.Done()

			for !done {
				cmd := exec.Command(cron.Command.Path, cron.Command.Args...)
				err := cmd.Run()
				if err != nil {
					log.Fatalln(err.Error())
				}

				// Log success message
				log.Printf("Command '%s %s' ran successfully", cron.Command.Path, strings.Join(cron.Command.Args, " "))

				// The program "sleeps" for cron.TimeLapse Minutes.
				timeoutCtx, _ := context.WithTimeout(cancelCtx, cron.TimeLapse*time.Minute)
				select {
				case <-timeoutCtx.Done():
					if cancelCtx.Err() != nil {
						log.Printf("Shutting down the service for command '%s %s'", cron.Command.Path, strings.Join(cron.Command.Args, " "))
					}
				}
			}

		}(cron)
	}

	wg.Wait()
}

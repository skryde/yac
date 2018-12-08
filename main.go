package main

import (
	"fmt"
	"github.com/skryde/jsconf"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

const cronsFile = "crons.json"

var crons Crons

func init() {
	switch jsconf.Exist(cronsFile) {
	case jsconf.NotExist:
		crons = Crons{}
		crons = append(crons, Cron{Command: Cmd{Path: "/bin/bash", Args: []string{"~/bin/script.sh"}}, TimeLapse: 5, TimeUnit: time.Second})

		err := jsconf.SaveToFile(cronsFile, crons)
		if err != nil {
			panic(err)
		}

		fmt.Println("Created file crons.json.\nEdit it as your convenient and run yac again.")
		os.Exit(0)

	case jsconf.IsFile:
		err := jsconf.LoadFromFile(cronsFile, &crons)
		if err != nil {
			panic(err)
		}

	default:
		panic("Unhandled situation")
	}
}

func main() {
	wg := sync.WaitGroup{}

	for _, cron := range crons {
		wg.Add(1)

		go func() {
			for true {
				cmd := exec.Command(cron.Command.Path, cron.Command.Args...)
				err := cmd.Run()
				if err != nil {
					log.Fatalln(err.Error())
				}

				time.Sleep(cron.TimeLapse * cron.TimeUnit)
			}

			wg.Done()
		}()
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"github.com/skryde/jsconf"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

const cronsFile = "crons.json"

var crons Crons

func init() {
	var (
		cmdPath string
		cmdArg  string
	)

	if runtime.GOOS == "windows" {
		cmdPath = fmt.Sprintf("%s\\System32\\cmd.exe", os.Getenv("WINDIR"))
		cmdArg = fmt.Sprintf("%s\\bin\\script.bat", os.Getenv("USERPROFILE"))

	} else {
		cmdPath = "/bin/bash"
		cmdArg = "~/bin/script.sh"
	}

	switch jsconf.Exist(cronsFile) {
	case jsconf.NotExist:
		crons = Crons{}
		crons = append(crons, Cron{Command: Cmd{Path: cmdPath, Args: []string{cmdArg}}, TimeLapse: 5, TimeUnit: time.Minute})

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

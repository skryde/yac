package main

import (
	"fmt"
	"github.com/skryde/jsconf"
	"os"
	"runtime"
	"time"
)

func init() {
	var (
		cmdPath string
		cmdArg  []string
	)

	if runtime.GOOS == "windows" {
		cmdPath = fmt.Sprintf("%s\\System32\\cmd.exe", os.Getenv("WINDIR"))
		cmdArg = []string{"/C", fmt.Sprintf("%s\\bin\\script.bat", os.Getenv("USERPROFILE"))}

	} else {
		cmdPath = "/bin/bash"
		cmdArg = []string{fmt.Sprintf("%s/bin/script.sh", os.Getenv("HOME"))}
	}

	switch jsconf.Exist(cronsFile) {
	case jsconf.NotExist:
		crons = Crons{}
		crons = append(crons, Cron{Command: Cmd{Path: cmdPath, Args: cmdArg}, TimeLapse: 5, TimeUnit: time.Minute})

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

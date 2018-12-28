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

	// The 'crons.json' file should be inside home directory, this function
	// returns the full path to it.
	cronsFile = getCronsFilePath()

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

		fmt.Println("Created file crons.json.\nEdit it as your convenient and run yac again.\nFull path:", cronsFile)
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

func getCronsFilePath() string {

	var cronsFileDir string

	if runtime.GOOS == "windows" {
		cronsFileDir = fmt.Sprintf("%s\\yac", os.Getenv("APPDATA"))

	} else {
		cronsFileDir = fmt.Sprintf("%s/.config/yac", os.Getenv("HOME"))
	}

	switch jsconf.Exist(cronsFileDir) {
	case jsconf.NotExist:
		if err := os.MkdirAll(cronsFileDir, 0750); err != nil {
			panic(err)
		}

	case jsconf.IsFile:
		panic("It's a file... That doesn't seems right.")
	}

	return fmt.Sprintf("%s%c%s", cronsFileDir, os.PathSeparator, cronsFileName)
}

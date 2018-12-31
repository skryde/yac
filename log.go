package main

import (
	"fmt"
	"github.com/skryde/jsconf"
	"log"
	"os"
	"runtime"
)

func logConfig() (*os.File, error) {

	var cronsLogFileDir string

	if runtime.GOOS == "windows" {
		cronsLogFileDir = fmt.Sprintf("%s\\yac\\logs", os.Getenv("APPDATA"))

	} else {
		cronsLogFileDir = "/var/log"
	}

	switch jsconf.Exist(cronsLogFileDir) {
	case jsconf.NotExist:
		if err := os.MkdirAll(cronsLogFileDir, 0750); err != nil {
			panic(err)
		}

	case jsconf.IsFile:
		panic("It's a file... That doesn't seems right.")
	}

	// Idea from https://stackoverflow.com/a/19966217/3508426
	file, err := os.OpenFile(fmt.Sprintf("%s%c%s", cronsLogFileDir, os.PathSeparator, cronsLogFileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	return file, nil
}

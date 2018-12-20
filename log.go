package main

import (
	"log"
	"os"
)

func logConfig() (*os.File, error) {
	// Idea from https://stackoverflow.com/a/19966217/3508426
	file, err := os.OpenFile("yac.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	return file, nil
}

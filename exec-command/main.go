package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	destinationDirPath, err := os.Getwd()

	log := filepath.Join(destinationDirPath, "nohup2.out")
	cmd := exec.Command("nohup", "sh", "-c", "/Users/afdagli/Desktop/Study/docker-slim/bin/mac/docker-slim-backend")

	f, err := os.Create(log)
	if err != nil {
		// handle error
	}

	// redirect both stdout and stderr to the log file
	cmd.Stdout = f
	cmd.Stderr = f

	// start command (and let it run in the background)
	err = cmd.Start()
	if err != nil {
		// handle error
	}

}

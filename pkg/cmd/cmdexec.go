package cmd

import (
	"log"
	"os/exec"
	"strings"
)

func CmdExec(command, dir string, panicOnError bool) ([]byte, error) {
	log.Printf("running %s\n", command)
	commandAndArgs := strings.Split(command, " ")
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	if dir != "" {
		cmd.Dir = dir
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		if panicOnError {
			log.Fatal(err) //let's call this "set -e"
		}
		return out, err
	}
	return out, nil
}

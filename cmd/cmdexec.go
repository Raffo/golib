package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func cmdExec(command string, panicOnError bool) ([]byte, error) {
	fmt.Printf("running %s\n", command)
	commandAndArgs := strings.Split(command, " ")
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)

	outReader, _ := cmd.StdoutPipe()
	errReader, _ := cmd.StderrPipe()

	cmd.Start()

	cmdReader := io.MultiReader(outReader, errReader)

	reader := bufio.NewReader(cmdReader)
	line, err := reader.ReadString('\n')
	for err == nil {
		fmt.Println(line)
		line, err = reader.ReadString('\n')
	}
	cmd.Wait()
	return nil, nil
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	action := os.Args[1]
	args := os.Args[2:]

	var cmd *exec.Cmd
	var output []byte
	var err error

	for {
		cmd = exec.Command(action, args...)

		output, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Failed to run", action, args)
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(output))
		time.Sleep(2 * time.Second)

	}
}

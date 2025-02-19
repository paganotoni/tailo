package main

import (
	"fmt"
	"os"
	"os/exec"
)

// runs the Tailwind CSS binary with the given arguments
func run(binary string, args []string) error {
	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Running:", cmd)
	return cmd.Run()
}

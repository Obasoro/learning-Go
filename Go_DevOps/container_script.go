package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const container_name = "kunle_learning"

func isContainerRunning(name string) bool {
	cmd:= exec.Command("docker", "inspect", "-f", "{{.State.Running}}",
	container_name)

	out, err := cmd.Output()
	if err != nil {
		fmt.Print("Failed to inspect container: %v\n", err)
	return false
	}

	return strings.TrimSpace(string(out)) == "true"
}

func reStartContainer() {
	cmd:= exec.Command("docker", "restart", container_name)
	err:= cmd.Run()
	if err != nil {
		fmt.Print("Failed to restart container: %v\n", err)
	}
}

func main() {
	for {
		if !isContainerRunning() {
			fmt.Printf("Container %s is not running. Restarting...\n", container_name)
			reStartContainer()
		}
		time.Sleep(10*time.Second)
	}

}
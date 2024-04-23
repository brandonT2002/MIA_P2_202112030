package main

import (
	"bufio"
	"fmt"
	callparser "mia/CallParser"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var clearCmd *exec.Cmd
	if os.Getenv("GOOS") == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		if strings.ToLower(input) != "exit" {
			call := callparser.CallParser{}
			call.ExecutionParser(input)
			println()
			continue
		}
		break
	}
}

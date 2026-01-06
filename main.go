package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: fountain <file>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	firstline, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
		return
	}

	if strings.HasPrefix(firstline, "#!") {
		shebang := strings.TrimSpace(firstline[2:])
		var cmd *exec.Cmd

		switch shebang {
		case "/usr/local/bin/fountain":
			cmd = exec.Command("./core/f1", filename)

			// changing main process and exiting launcher
			args := []string{"./core/f1", filename}
			syscall.Exec("./core/f1", args, os.Environ())

		case "/usr/local/bin/ff":
			cmd = exec.Command("./core/f1", filename)

			// changing main process and exiting launcher
			args := []string{"./core/f1", filename}
			syscall.Exec("./core/f1", args, os.Environ())

		default:
			cmd = exec.Command(shebang, filename)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("error")
		}
	} else {
		fmt.Println("error")
	}
}

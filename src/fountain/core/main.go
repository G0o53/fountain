package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("oh noes")
        return
    }

    filename := os.Args[1]
    fmt.Println("script:", filename)
    fmt.Println("-----------")
    fmt.Println("")

    // execute via bash (will change to execute via own shell)
    cmd := exec.Command("bash", filename)

    // stdout pipe
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }

    // stderr pipe
    stderr, err := cmd.StderrPipe()
    if err != nil {
        log.Fatal(err)
    }

    // start the command
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    // stream output
    go io.Copy(os.Stdout, stdout)
    go io.Copy(os.Stderr, stderr)

    // wait for command to finish
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
}

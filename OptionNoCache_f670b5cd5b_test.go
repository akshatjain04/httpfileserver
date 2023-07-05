package main

import (
    "os/exec"
    "log"
)

func main() {
    cmd := exec.Command("apt-get", "install", "gcc")
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

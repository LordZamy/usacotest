package main

import (
    "os"
    "os/exec"
    "github.com/fatih/color"
)

var (
    progPath string
    testDir string
)

func main() {
    initVars()

    cmd := exec.Command(progPath)
    err := cmd.Run()
    if err != nil {
        color.Red(err.Error())
        os.Exit(0)
    }
}

func initVars() {
    numArgs := len(os.Args)
    if numArgs >= 2 {
        progPath = os.Args[1]
    } else {
        color.Red("Please specify exectable as first argument")
        os.Exit(0)
    }

    if numArgs >= 3 {
        testDir = os.Args[2]
    } else {
        x, err := os.Getwd()
        if err != nil {
            color.Red("Error getting current working directory")
            os.Exit(0)
        }
        testDir = x
    }
}

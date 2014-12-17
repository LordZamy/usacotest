package main

import (
    "os"
    "github.com/fatih/color"
)

var (
    progPath string
    testDir string
)

func main() {
    initVars()
}

func initVars() {
    numArgs := len(os.Args)
    if numArgs >= 2 {
        progPath = os.Args[1]
    } else {
        color.Red("Please specify exectable as first argument")
    }

    if numArgs >= 3 {
        testDir = os.Args[2]
    } else {
        x, err := os.Getwd()
        if err != nil {
            color.Red("Error getting current working directory")
        }
        testDir = x
    }
    color.Cyan(testDir)
}

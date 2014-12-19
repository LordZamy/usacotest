package main

import (
    "io/ioutil"
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
        color.Red("Error running program:\n" + err.Error())
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
            color.Red("Error getting current working directory:\n" + err.Error())
            os.Exit(0)
        }
        testDir = x
    }
}

func copyTestData(fileName string) {
    b, err := ioutil.ReadFile(fileName + ".in")
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(progPath + ".in", b, 0644)
    if err != nil {
        panic(err)
    }
}

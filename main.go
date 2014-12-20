package main

import (
    "io/ioutil"
    "os"
    "os/exec"
    "strconv"
    "github.com/fatih/color"
)

var (
    progPath string
    testDir string
    ioStyle int
)

func main() {
    initVars()

    // chdir to directory containg test cases
    os.Chdir(testDir)

    setIOStyle()

    i := 1

    if ioStyle == 1 {
        for {
            if _, err := os.Stat(strconv.Itoa(i) + ".in"); os.IsNotExist(err) {
                // TODO: Add number of total test cases passed
                color.White("Finished testing")
                os.Exit(0)
            }

            copyTestData(strconv.Itoa(i) + ".in")

            runProgram()

            a := readOutput(progPath + ".out")
            b := readOutput(strconv.Itoa(i) + ".out")

            if compareOutput(a, b) != -1 {
                color.Red("Wrong output at character " + strconv.Itoa(i) + ":\nExpected: " + string(a[i]) + "\tFound: " + string(b[i]))
            } else {
                color.Green("Test case " + strconv.Itoa(i) + " passed");
            }
            i++
        }
    } else if ioStyle == 2 {
        for {
            if _, err := os.Stat("I." + strconv.Itoa(i)); os.IsNotExist(err) {
                // TODO: Add number of total test cases passed
                color.White("Finished testing")
                os.Exit(0)
            }

            copyTestData("I." + strconv.Itoa(i))

            runProgram()

            a := readOutput(progPath + ".out")
            b := readOutput("O." + strconv.Itoa(i))

            if compareOutput(a, b) != -1 {
                color.Red("Wrong output at character " + strconv.Itoa(i) + ":\nExpected: " + string(a[i]) + "\tFound: " + string(b[i]))
            } else {
                color.Green("Test case " + strconv.Itoa(i) + " passed");
            }
            i++
        }

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
    b, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(progPath + ".in", b, 0644)
    if err != nil {
        panic(err)
    }
}

// if outputs are not same, return position of nonequivalence; otherwise return -1
func compareOutput(a, b []byte) int {
    fileLen := len(a)
    for i := 0; i < fileLen; i++ {
        if a[i] != b[i] {
            return i
        }
    }
    return -1
}

func readOutput(fileName string) []byte {
    b, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    return b
}

// ioStyle 1 is of form "X.in"; ioStyle 2 is of form "I.X";
// TODO: Change function to check if list of files in "testDir" contains "X.in" or "I.X" files using regex
func setIOStyle() {
    if _, err := os.Stat("1.in"); err == nil {
        ioStyle = 1
        return
    } else if _, err = os.Stat("I.1"); err == nil {
        ioStyle = 2
        return
    }
    color.Red("Could not find test input/ouput files in " + testDir)
    os.Exit(0)
}

func runProgram() {
    cmd := exec.Command(progPath)
    err := cmd.Run()
    if err != nil {
        color.Red("Error running program:\n" + err.Error())
        os.Exit(0)
    }
}

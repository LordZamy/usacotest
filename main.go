package main

import (
    "io/ioutil"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "time"
    "github.com/fatih/color"
)

var (
    progPath string
    progName string
    testDir string
    ioStyle int
    isPython bool
)

func main() {
    initVars()

    // chdir to directory containing test cases
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

            timeElapsed := runProgram()

            a := readOutput(progName + ".out")
            b := readOutput(strconv.Itoa(i) + ".out")

            if compareOutput(a, b) != -1 {
                color.Red("Wrong output for test case %d:\nExpected:\n%sFound:\n%s\n", i, string(b[:]), string(a[:]))
            } else {
                color.Green("Test case %d passed in %.3f seconds", i, timeElapsed.Seconds());
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

            timeElapsed := runProgram()

            a := readOutput(progName + ".out")
            b := readOutput("O." + strconv.Itoa(i))

            if compareOutput(a, b) != -1 {
                color.Red("Wrong output for test case %d:\nExpected:\n%sFound:\n%s\n", i, string(b[:]), string(a[:]))
            } else {
                color.Green("Test case %d passed in %.3f seconds", i, timeElapsed.Seconds());
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
        color.Red("Please specify program path as first argument")
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

    progName = getProgName()

    if progName[len(progName) - 3:] == ".py" {
        isPython = true
        progName = progName[:len(progName) - 3]
    }
}

func getProgName() string {
    pos := strings.LastIndex(progPath, "/")
    if pos == -1 {
        pos = strings.LastIndex(progPath, "\\")
    }
    return progPath[pos + 1:]
}

func copyTestData(fileName string) {
    b, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(progName + ".in", b, 0644)
    if err != nil {
        panic(err)
    }
}

// if outputs are not same, return position of nonequivalence; otherwise return -1
func compareOutput(a, b []byte) int {
    fileLen := len(a)
    bLen := len(b)
    if fileLen != bLen {
        if fileLen > bLen {
            return fileLen - bLen
        }
        return bLen - fileLen
    }
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

func runProgram() time.Duration {
    cmd := exec.Command(progPath)
    if isPython {
        // TODO: Think of fix for python command name on different OSes (python2, python3)
        cmd = exec.Command("python", progPath)
    }
    timeStart := time.Now()
    err := cmd.Run()
    timeEnd := time.Now()
    if err != nil {
        color.Red("Error running program:\n" + err.Error())
        os.Exit(0)
    }
    return timeEnd.Sub(timeStart)
}

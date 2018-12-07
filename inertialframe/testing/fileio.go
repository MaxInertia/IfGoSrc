package main

import (
	"io/ioutil"
	"fmt"
	"runtime"
)


var local1 = "fileio.go"
var local2 = "../collection/chain.go"

// Windows Only Absolute Paths

var absoluteWin = "C:/Users/Kraken/Go/Src/local/IfGoSrc/inertialframe/testing/fileio.go"

// Linux only Absolute Paths

var absoluteLin = "/mnt/c/Users/Kraken/Go/Src/local/IfGoSrc/inertialframe/testing/fileio.go"

func confirmAccess(pathname string) []byte {
	dat, err := ioutil.ReadFile(pathname)
	if err != nil {
		panic(err)
	}
	return dat
}

func main() {
	//fmt.Printf("OS: %s", runtime.GOOS)

	confirmAccess(local1)
	confirmAccess(local2)

	if runtime.GOOS == "windows" {
		fmt.Println("Hello from Windows")
		confirmAccess(absoluteWin)
	}
	if runtime.GOOS == "linux" {
		fmt.Println("Hello from Linux")
		confirmAccess(absoluteLin)
	}
}
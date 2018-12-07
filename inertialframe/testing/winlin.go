package main

import (
	"os/exec"
	"os"
)

func createGoCmd(path string) *exec.Cmd {
	cmd := new(exec.Cmd)
	cmd.Path = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func main() {
	args := os.Args

	if len(args) != 1 {

	}

	linPath, linErr := exec.LookPath("go") // Linux version of Golang
	winPath, winErr := exec.LookPath("go.exe") // Windows version of Golang

	if linErr != nil {
		print(linErr)
	} else {
		linGo := createGoCmd(linPath)
		linGo.Run()
	}

	if winErr != nil {
		print(winErr)
	} else {
		winGo := createGoCmd(winPath)
		winGo.Run()
	}

}
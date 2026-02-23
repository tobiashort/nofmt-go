package main

import (
	"os"
	"os/exec"

	"github.com/tobiashort/clog-go"
)

func main() {
	install()
	readme()
}

func install() {
	clog.Info("install...")
	cmd := exec.Command("go", "install", "./nofmt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func readme() {
	clog.Info("readme...")
	cmd := exec.Command("go", "run", "./nofmt", "-h")
	out, err := cmd.CombinedOutput()
	if err != nil {
		clog.Error(string(out))
		os.Exit(1)
	} else {
		os.WriteFile("README.md", out, 0644)
	}
}

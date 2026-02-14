package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tobiashort/clap-go"
	"github.com/tobiashort/utils-go/must"
)

type Args struct {
	File string `clap:"positional,mandatory,description='the file to format'"`
}

func main() {
	args := Args{}
	clap.Parse(&args)

	var replacements [][]string
	var replacement []string
	var enabled bool

	file := must.Do2(os.Open(args.File))
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "//nofmt:enable" {
			enabled = true
		} else if strings.TrimSpace(line) == "//nofmt:disable" {
			enabled = false
			replacements = append(replacements, replacement)
			replacement = make([]string, 0)
		} else {
			if enabled {
				replacement = append(replacement, line)
			}
		}
	}
	replacements = append(replacements, replacement)
	replacement = make([]string, 0)

	cmd := exec.Command("goimports", args.File)
	out := string(must.Do2(cmd.CombinedOutput()))

	enabled = false
	replacementIndex := 0
	scanner = bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "//nofmt:enable" {
			fmt.Println(line)
			enabled = true
			replacement := replacements[replacementIndex]
			for _, replacementLine := range replacement {
				fmt.Println(replacementLine)
			}
			replacementIndex++
		} else if strings.TrimSpace(line) == "//nofmt:disable" {
			fmt.Println(line)
			enabled = false
		} else {
			if enabled {
				continue
			} else {
				fmt.Println(line)
			}
		}
	}
}

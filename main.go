package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	goRootBasePath = "/opt/homebrew/Cellar/go/"
	symLinkPath    = "bin/go"
	symLinkFmtPath = "bin/gofmt"
	goBinary       = "/usr/local/bin/go"
	goFmt          = "/usr/local/bin/gofmt"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println(`
use: list
	to print all of golang version.
use: sw [version]
	to change golang version.
		`)
		return
	}
	files, err := ioutil.ReadDir(goRootBasePath)
	if err != nil {
		fmt.Println("read path err :", goRootBasePath, err)
		return
	}

	switch args[1] {
	case "list":
		{
			for _, f := range files {
				fmt.Println(f.Name())
			}
		}
	case "sw":
		{
			if len(args) < 3 {
				fmt.Println(`
use: list
	to print all of golang version.
use: sw [version]
	to change golang version.
		`)
				return
			}
			ver := args[2]

			for _, f := range files {
				if ver == f.Name() {
					symLink := goRootBasePath + f.Name() + "/" + symLinkPath
					symLinkFmt := goRootBasePath + f.Name() + "/" + symLinkFmtPath
					err := os.Remove(goBinary)
					if err != nil {
						fmt.Println("delete link error", err)
						return
					}
					err = os.Symlink(symLink, goBinary)
					if err != nil {
						fmt.Println("set link error", err)
						return
					}
					err = os.Remove(goFmt)
					if err != nil {
						fmt.Println("delete link error", err)
						return
					}
					err = os.Symlink(symLinkFmt, goFmt)
					if err != nil {
						fmt.Println("set link error", err)
						return
					}
				}
			}
		}
	default:
		fmt.Println(`
use: list
	to print all of golang version.
use: sw [version]
	to change golang version.
		`)
	}

}

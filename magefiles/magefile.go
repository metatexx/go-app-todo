package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
)

const appName = "todo"

// Build runs go mod download and then builds the binary and wasm parts of the app.
func Build() error {
	fmt.Println("Building...")
	if err := sh.RunV("go", "mod", "download"); err != nil {
		return err
	}
	if err := sh.RunWithV(map[string]string{"GOOS": "js", "GOARCH": "wasm"},
		"go", "build", "-o", "./web/app.wasm", "./server/"); err != nil {
		return err
	}
	if err := sh.RunV("go", "build", "-o", "./bin/"+appName, "./server/"); err != nil {
		return err
	}
	fmt.Println("Done")
	return nil
}

func Run() error {
	mg.Deps(Build)
	fmt.Println("Running...")
	return sh.RunV("./bin/" + appName)
}

func Clean() error {
	// we do not check errors because the files will not exist all the time
	fmt.Println("Cleaning...")
	_ = os.Remove("./bin/" + appName)
	_ = os.Remove("./web/app.wasm")
	fmt.Println("Done")
	return nil
}

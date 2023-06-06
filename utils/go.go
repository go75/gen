package utils

import (
	"os"
	"os/exec"
)

func WriteGoFile(filepath, content string) error {
	err := os.WriteFile(filepath, []byte(content), 0666)
	if err != nil {
		return err
	}

	err = GoFmt(filepath)
	
	return err
}

func GoFmt(filepath string) error {
	return exec.Command("gofmt", "-w", filepath).Run()
}

func GoGet(packageName string) error {
	cmd := exec.Command("go", "get", packageName)
	return cmd.Run()
}

func GoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	return cmd.Run()
}
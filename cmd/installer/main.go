package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	println("Installing Rice Paper")
	cacheDir, err := getCacheDir()
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	println(cacheDir)
	err = linkXresourceFile(cacheDir)
}

func getCacheDir() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return dir, nil
}

func linkXresourceFile(cacheDir string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fileName := filepath.Join(homeDir, ".Xresources")
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	templateFileName := filepath.Join(cacheDir, "rice-paper/rice-paper.Xresources")
	includeString := "#include " + templateFileName
	_, err = file.WriteString(includeString)
	if err != nil {
		return err
	}
	return nil
}

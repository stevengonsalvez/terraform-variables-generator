package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func getAllFiles(ext string) ([]string, error) {
	dir, err := os.Getwd()
	checkError(err)
	var files []string
	log.Infof("Finding files in %q directory", dir)
	files, err = filepath.Glob(tfFileExt)
	checkError(err)

	if len(files) == 0 {
		log.Infof("No files with .tf extensions found in %q", dir)
		os.Exit(0)
	}
	return files, nil
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	}
	return false
}

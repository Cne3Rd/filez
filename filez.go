package filez

// A Simple package Showing You File Infomation,
// A tiny Wrapper Around The OS Stat Method

import (
	"os"
	"time"
)

func FileSize(file string) (int, error) {

	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	fsize := int(f.Size())
	return fsize, nil
}

func ModTime(file string) (time.Time, error) {
	f, err := os.Stat(file)
	if err != nil {
		return time.Time{}, err
	}

	return f.ModTime(), nil
}

func FileMode(file string) (os.FileMode, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return f.Mode(), nil
}

func FileName(file string) (string, error) {
	fname, err := os.Stat(file)
	if err != nil {
		return "", err
	}
	name := fname.Name()
	return name, nil
}

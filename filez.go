package filez

// A Simple package Showing You File Infomation,
// A tiny Wrapper Around The OS Stat Method

import (
	"os"
	"time"
)

func FileSize(file string) int {

	f, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	fsize := int(f.Size())
	return fsize
}

func ModTime(file string) time.Time {
	f, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	return f.ModTime()
}

func FileMode(file string) os.FileMode {
	f, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	return f.Mode()
}

func FileName(file string) string {
	fname, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	name := fname.Name()
	return name
}

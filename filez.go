package filez

// A Simple package Showing You File Infomation,
// A tiny Wrapper Around The OS Stat Method

import (
	"os"
	"time"
)

// Return file size(length in bytes for regular files; system-dependent for others) if no error,
// Return 0 and error if there is error
func FileSize(file string) (int, error) {

	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	fsize := int(f.Size())
	return fsize, nil
}

// Return file Modification time if no error,
// Return The zero value of type Time(January 1, year 1, 00:00:00.000000000 UTC) if there is an error

func ModTime(file string) (time.Time, error) {
	f, err := os.Stat(file)
	if err != nil {
		return time.Time{}, err
	}

	return f.ModTime(), nil
}

// Return file mode if no error,
// Return 0 and error if there is error
func FileMode(file string) (os.FileMode, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return f.Mode(), nil
}

// Return file name if no error,
// return empty string and error if there is error
func FileName(file string) (string, error) {
	fname, err := os.Stat(file)
	if err != nil {
		return "", err
	}
	name := fname.Name()
	return name, nil
}

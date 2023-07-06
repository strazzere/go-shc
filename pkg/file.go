package goshc

import (
	"os"
	"runtime"
)

func Open(data []byte) (*os.File, error) {
	if runtime.GOOS == "android" {
		file, err := os.CreateTemp("", "tmp")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		err = file.Chmod(0755)
		if err != nil {
			return nil, err
		}

		_, err = file.Write(data)
		if err != nil {
			return nil, err
		}

		return file, nil
	}

	return MemOpen(data)
}

func Clean(file *os.File) {
	if runtime.GOOS == "android" {
		os.Remove(file.Name())
	} else {
		file.Close()
	}
}

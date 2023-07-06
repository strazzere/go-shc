package goshc

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func MemOpen(data []byte) (*os.File, error) {
	fileDescriptor, err := unix.MemfdCreate("", unix.MFD_CLOEXEC)
	if err != nil {
		return nil, err
	}

	file := os.NewFile(uintptr(fileDescriptor), fmt.Sprintf("/proc/%d/fd/%d", os.Getpid(), fileDescriptor))

	// Write data
	_, err = file.Write(data)
	if err != nil {
		_ = file.Close()
		return nil, err
	}

	// Seek to beginning so it is a "clean" file for whomever needed it
	_, err = file.Seek(0, 0)
	if err != nil {
		_ = file.Close()
		return nil, err
	}

	return file, nil
}

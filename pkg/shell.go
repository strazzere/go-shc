package goshc

import (
	"fmt"
	"os/exec"
)

func Execute(script []byte, direct bool, interpreter string) error {
	if direct {
		if interpreter == "" {
			interpreter = "sh"
		}
		cmd, err := exec.Command(interpreter, "-c", string(script)).Output()
		if err != nil {
			return fmt.Errorf("direct execute error %s", err)
		}
		output := string(cmd)

		fmt.Printf("Direct Output: \n<<<\n%v\n>>>", output)

		return nil
	}

	file, err := Open(script)
	if err != nil {
		return fmt.Errorf("open error %s", err)
	}

	cmd, err := exec.Command(file.Name()).Output()
	if err != nil {
		return fmt.Errorf("execute error %s", err)
	}
	output := string(cmd)

	fmt.Printf("Output: \n<<<\n%v\n>>>", output)
	Clean(file)

	return nil
}

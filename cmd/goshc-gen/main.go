package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	goshc "github.com/strazzere/go-shc/pkg"
)

var (
	archFlag        string
	osFlag          string
	useGarbleFlag   bool
	directFlag      bool
	interpreterFlag string
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: go-shc SCRIPT

Options:
`)
		flag.PrintDefaults()
	}
	flag.StringVar(&archFlag, "arch", "arm64",
		"`arch` to compile the script to, defaults to arm64",
	)
	flag.StringVar(&osFlag, "os", "android",
		"`arch` to compile the script to, defaults to arm64",
	)
	flag.BoolVar(&useGarbleFlag, "garble", false,
		"use garble instead of go to build the file, defaults to false",
	)
	flag.BoolVar(&directFlag, "direct", false,
		"directly pipe the script to an interpreter (no file used), defaults to false",
	)
	flag.StringVar(&interpreterFlag, "interpreter", "sh",
		"if `direct` was used, attempt to use this specific interpreter, defaults to sh with no direct path",
	)
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	tmpFile, err := buildLoaderTemplate(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(127)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	err = compile(tmpFile.Name(), fmt.Sprintf("%s.x", flag.Arg(0)), useGarbleFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(127)
	}
}

func compile(input string, outname string, garble bool) error {
	goCmd := "go"

	var args, prebuildFlags, postBuildFlags []string
	if garble {
		goCmd = "garble"
		prebuildFlags = append(prebuildFlags, "-tiny")
		prebuildFlags = append(prebuildFlags, "-literals")
	} else {
		// None of the three below approaches work for some reason, unsure why

		// postBuildFlags = []string{`-ldflags "-s"`}

		// postBuildFlags = append(postBuildFlags, "-ldflags \"-s\"")

		// postBuildFlags = append(postBuildFlags, "-ldflags")
		// postBuildFlags = append(postBuildFlags, `"-s"``)
	}
	if len(prebuildFlags) > 0 {
		args = append(args, prebuildFlags...)
	}
	args = append(args, "build")
	args = append(args, "-v")
	args = append(args, "-o")
	args = append(args, outname)
	if len(postBuildFlags) > 0 {
		args = append(args, postBuildFlags...)
	}
	args = append(args, input)
	cmd := exec.Command(goCmd, args...) //strings.Join(args, " "))
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOOS=%s", osFlag))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOARCH=%s", archFlag))
	var errb bytes.Buffer
	cmd.Stderr = &errb
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error using %v\n%v", cmd, errb.String())
		return err
	}
	fmt.Println(cmd)
	fmt.Print(string(out))

	return nil
}

func buildLoaderTemplate(script string) (*os.File, error) {
	newpath := filepath.Join(".", "gen")
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	output, err := os.Create("./gen/test.go")
	if err != nil {
		return nil, err
	}

	payloadBytes, err := os.ReadFile(script)
	if err != nil {
		fmt.Printf("Error loading payload : %v", err)
		return nil, err
	}

	key := goshc.GenerateKey()

	goshc.Crypt(payloadBytes, key)

	template := "// Code generated by goshc-gen. DO NOT EDIT." + "\n" +
		"package main" + "\n" +
		"" + "\n" +
		"import (" + "\n" +
		`	"fmt"` + "\n" +
		"" + "\n" +
		`	goshc "github.com/strazzere/go-shc/pkg"` + "\n" +
		")" + "\n" +
		"" + "\n" +
		"var (" + "\n" +
		"	payload = []byte{" + goshc.ToGoString(payloadBytes) + " }" + "\n" +
		"	key     = []byte{" + goshc.ToGoString(key) + "}" + "\n" +
		")" + "\n" +
		"" + "\n" +
		"//garble:controlflow" + "\n" +
		"func main() {" + "\n" +
		"	err := goshc.Crypt(payload, key)" + "\n" +
		"	if err != nil {" + "\n" +
		`		fmt.Printf("error %s", err)` + "\n" +
		"		return" + "\n" +
		"	}" + "\n" +
		"" + "\n" +
		`	err = goshc.Execute(payload, ` + strconv.FormatBool(directFlag) + `, "` + interpreterFlag + `")` + "\n" +
		"	if err != nil {" + "\n" +
		`		fmt.Printf("execute error %s", err)` + "\n" +
		"	}" + "\n" +
		"}"

	_, err = output.Write([]byte(template))
	if err != nil {
		return nil, err
	}

	return output, nil
}

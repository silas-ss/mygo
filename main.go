package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	switch strings.Trim(os.Args[1], "") {
	case "init":
		must(initialize())
	case "install":
		must(install())
	case "remove":
		fmt.Println("remove package")
	default:
		fmt.Println("mygo init <url> - inicializa um projeto")
		fmt.Println("mygo install <url> - instala um novo pacote")
		fmt.Println("mygo remove <url>  - remove um pacote instalado")
	}
}

func initialize() error {
	var err error
	if len(os.Args) < 3 {
		return errors.New("Informe o endereço do pacote")
	}

	goExecutablePath, err := exec.LookPath("go")
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path:   goExecutablePath,
		Args:   []string{goExecutablePath, "mod", "init", os.Args[2]},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func install() error {
	var err error
	if len(os.Args) < 3 {
		return errors.New("Informe o endereço do pacote")
	}

	goExecutablePath, err := exec.LookPath("go")
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path:   goExecutablePath,
		Args:   []string{goExecutablePath, "get", "-v", os.Args[2]},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

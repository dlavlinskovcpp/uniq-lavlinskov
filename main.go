package main

import (
	"fmt"
	"os"

	"github.com/dlavlinskovcpp/uniq-lavlinskov/internal/uniq"
)

func main() {
	opts, args, err := uniq.ParseFlags()
	if err != nil {
		fatal("ошибка флагов:", err)
	}

	if err := opts.Validate(); err != nil {
		fatal("ошибка опций:", err)
	}

	var input, output *os.File
	switch len(args) {
	case 0:
		input, output = os.Stdin, os.Stdout
	case 1:
		input = mustOpen(args[0], os.O_RDONLY, 0)
		output = os.Stdout
	case 2:
		input = mustOpen(args[0], os.O_RDONLY, 0)
		output = mustOpen(args[1], os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	default:
		fatal("слишком много аргументов")
	}
	defer input.Close()
	if output != os.Stdout {
		defer output.Close()
	}

	if err := uniq.Run(input, output, opts); err != nil {
		fatal("ошибка:", err)
	}
}

func mustOpen(name string, flag int, perm os.FileMode) *os.File {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		fatal("не удалось открыть файл:", name, err)
	}
	return f
}

func fatal(args ...interface{}) {
	fmt.Fprint(os.Stderr, "uniq: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

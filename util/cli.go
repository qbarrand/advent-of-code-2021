package util

import "flag"

type CommandLine struct {
	InputFile string
}

func ParseCommandLine() *CommandLine {
	cli := CommandLine{}

	flag.StringVar(&cli.InputFile, "i", "input.txt", "name of the input file")

	flag.Parse()

	return &cli
}

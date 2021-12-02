package util

import "flag"

type CommandLine struct {
	InputFile string
}

func ParseCommandLine(programName string, args []string) (*CommandLine, error) {
	cli := CommandLine{}

	fs := flag.NewFlagSet(programName, flag.ContinueOnError)

	fs.StringVar(&cli.InputFile, "i", "input.txt", "name of the input file")

	return &cli, fs.Parse(args)
}

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type bitColumn struct {
	ones   int
	zeroes int
}

func (bc bitColumn) leastCommon() int {
	if bc.ones > bc.zeroes {
		return 0
	}

	return 1
}

func (bc bitColumn) mostCommon() int {
	if bc.ones > bc.zeroes {
		return 1
	}

	return 0
}

func parseLine(cols []bitColumn, line string) {
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '0':
			cols[i].zeroes++
		case '1':
			cols[i].ones++
		default:
			log.Fatalf("%c: invalid character", line[i])
		}
	}
}

func pow(a, b int) int {
	ret := 1

	for i := 1; i <= b; i++ {
		ret *= a
	}

	return ret
}

func main() {
	cl, err := util.ParseCommandLine(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("Could not parse the command line: %v", err)
	}

	fd, err := os.Open(cl.InputFile)
	if err != nil {
		log.Fatalf("Could not open the input file: %v", err)
	}
	defer fd.Close()

	var line string

	if _, err := fmt.Fscanf(fd, "%s", &line); err != nil {
		log.Fatalf("Could not read the first line: %v", err)
	}

	L := len(line)

	cols := make([]bitColumn, L)

	parseLine(cols, line)

	for i := 2; ; i++ {
		if _, err := fmt.Fscanf(fd, "%s", &line); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not read line %d: %v", i, err)
		}

		parseLine(cols, line)
	}

	var gamma, epsilon int

	j := 0

	for i := L - 1; i >= 0; i-- {
		gamma += cols[i].mostCommon() * pow(2, j)
		epsilon += cols[i].leastCommon() * pow(2, j)

		j++
	}

	log.Printf("Part 1: %d", gamma*epsilon)

	if _, err := fd.Seek(0, 0); err != nil {
		log.Fatalf("Could not seek(0,0): %v", err)
	}
}

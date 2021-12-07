package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/qbarrand/advent-of-code-2021/util"
)

func scanCommaSeparatedInts(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}

		if data[i] == '\n' {
			return i + 1, data[:i], bufio.ErrFinalToken
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	// There is one final token to be delivered, which may be the empty string.
	// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
	// but does not trigger an error to be returned from Scan itself.
	return 0, data, bufio.ErrFinalToken
}

func main() {
	cl, err := util.ParseCommandLine(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("Error while parsing the command line: %v", err)
	}

	fd, err := os.Open(cl.InputFile)
	if err != nil {
		log.Fatalf("Could not open the input file: %v", err)
	}
	defer fd.Close()

	fishes := make([]int, 0)

	s := bufio.NewScanner(fd)
	s.Split(scanCommaSeparatedInts)

	for s.Scan() {
		if err = s.Err(); err != nil {
			log.Fatalf("Error while scanning: %v", err)
		}

		t := s.Text()

		n, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			log.Fatalf("Could not parse %q as an integer: %v", t, err)
		}

		fishes = append(fishes, int(n))
	}

	for i := 0; i < 80; i++ {
		var newFishes []int

		for j := 0; j < len(fishes); j++ {
			if fishes[j] == 0 {
				fishes[j] = 6
				newFishes = append(newFishes, 8)
			} else {
				fishes[j]--
			}
		}

		fishes = append(fishes, newFishes...)
	}

	log.Printf("Part 1: %d", len(fishes))
}

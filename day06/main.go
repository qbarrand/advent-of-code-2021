package main

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type cacheKey struct {
	rounds int
	start  int
}

var cache = make(map[cacheKey]int)

func run(start int, rounds int) (res int) {
	ck := cacheKey{
		rounds: rounds,
		start:  start,
	}

	defer func() {
		if _, ok := cache[ck]; !ok {
			cache[ck] = res
		}
	}()

	if v, ok := cache[ck]; ok {
		res = v
		return
	}

	if rounds == 0 {
		res = 1
		return
	}

	if start == 0 {
		res = run(6, rounds-1) + run(8, rounds-1)
		return
	}

	res = run(start-1, rounds-1)
	return
}

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

	part1 := 0
	part2 := 0

	for _, i := range fishes {
		part1 += run(i, 80)
		part2 += run(i, 256)
	}

	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}

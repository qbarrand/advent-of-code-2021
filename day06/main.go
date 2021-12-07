package main

import (
	"bufio"
	"log"
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

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	part1 := 0
	part2 := 0

	s := bufio.NewScanner(fd)
	s.Split(util.ScanCommaSeparatedInts)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatalf("Error while scanning: %v", err)
		}

		t := s.Text()

		n, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			log.Fatalf("Could not parse %q as an integer: %v", t, err)
		}

		part1 += run(int(n), 80)
		part2 += run(int(n), 256)
	}

	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}

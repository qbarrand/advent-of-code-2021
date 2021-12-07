package main

import (
	"bufio"
	"log"
	"math"
	"strconv"

	"github.com/qbarrand/advent-of-code-2021/util"
)

func getFuel(pos []int, midPoint int, moveCost func(int) int) int {
	res := 0

	for _, p := range pos {
		distance := util.MaxInt(p, midPoint) - util.MinInt(p, midPoint)
		res += moveCost(distance)
	}

	return res
}

func constantCost(n int) int {
	return n
}

func sumCost(n int) int {
	return (n * (n + 1)) / 2
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	s := bufio.NewScanner(fd)
	s.Split(util.ScanCommaSeparatedInts)

	min := math.MaxInt32
	max := 0

	pos := make([]int, 0)

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatalf("Error while scanning: %v", err)
		}

		t := s.Text()

		n, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			log.Fatalf("Could not parse %q as an integer: %v", t, err)
		}

		i := int(n)

		pos = append(pos, i)

		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	part1 := math.MaxInt32
	part2 := math.MaxInt32

	for i := min; i <= max; i++ {
		if fuel := getFuel(pos, i, constantCost); fuel < part1 {
			part1 = fuel
		}

		if fuel := getFuel(pos, i, sumCost); fuel < part2 {
			part2 = fuel
		}
	}

	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}

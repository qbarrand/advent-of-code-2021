package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type posInGrid struct {
	grid  int
	line  int
	index int
}

type number struct {
	i      int
	marked bool
}

type grid struct {
	won  bool
	nums [][]number
}

func newGrid() *grid {
	nums := make([][]number, 5)

	for i := 0; i < 5; i++ {
		nums[i] = make([]number, 5)
	}

	return &grid{nums: nums}
}

func (g grid) markAndGetScore(line, index int) int {
	g.nums[line][index].marked = true
	row := true
	column := true

	for i := 0; i < 5; i++ {
		if !g.nums[line][i].marked {
			row = false
		}

		if !g.nums[i][index].marked {
			column = false
		}
	}

	if !row && !column {
		return -1
	}

	// compute score
	sum := 0

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if c := g.nums[i][j]; !c.marked {
				sum += c.i
			}
		}
	}

	return g.nums[line][index].i * sum
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var firstLine string

	if _, err := fmt.Fscanf(fd, "%s", &firstLine); err != nil {
		log.Fatalf("Could not read the first line: %v", err)
	}

	nums := make([]int, 0)

	for _, s := range strings.Split(firstLine, ",") {
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Fatalf("Could not parse %q as an integer: %v", s, err)
		}

		nums = append(nums, int(i))
	}

	s := bufio.NewScanner(fd)
	s.Split(bufio.ScanWords)

	grids := make([]*grid, 0)
	numMembership := make(map[int][]posInGrid)

	for i := 0; s.Scan(); i++ {
		if err := s.Err(); err != nil {
			log.Fatalf("Error while reading grids: %v", err)
		}

		gridIndex := i / 25
		lineIndex := (i - gridIndex*25) / 5
		index := i - gridIndex*25 - lineIndex*5

		if i%25 == 0 {
			grids = append(grids, newGrid())
		}

		t := s.Text()

		c, err := strconv.ParseInt(t, 10, 32)
		if err != nil {
			log.Fatalf("Could not parse %q as an integer: %v", t, err)
		}

		curr := int(c)

		grids[gridIndex].nums[lineIndex][index] = number{i: curr}

		pos := posInGrid{
			grid:  gridIndex,
			line:  lineIndex,
			index: index,
		}

		numMembership[curr] = append(numMembership[curr], pos)
	}

	var (
		part1 = -1
		part2 int
	)

	// Now, let's play
	for _, n := range nums {
		for _, pos := range numMembership[n] {
			if g := grids[pos.grid]; !g.won {
				if score := g.markAndGetScore(pos.line, pos.index); score > 0 {
					if part1 == -1 {
						part1 = score
					}

					part2 = score

					g.won = true
				}

			}
		}
	}

	// part1:
	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}

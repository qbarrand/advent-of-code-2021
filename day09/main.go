package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"sort"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type coord struct{ x, y int }

type basinFinder struct {
	grid    grid
	visited map[coord]bool
}

func newBasinFinder(g grid) *basinFinder {
	return &basinFinder{
		grid:    g,
		visited: make(map[coord]bool),
	}
}

func (bf *basinFinder) getSize(x, y int) int {
	c := coord{x: x, y: y}

	v := bf.grid[y][x]

	if bf.visited[c] || v == 9 {
		return 0
	}

	bf.visited[c] = true

	size := 1

	if x > 0 && bf.grid[y][x-1] > v {
		size += bf.getSize(x-1, y)
	}

	if x < len(bf.grid[y])-1 && bf.grid[y][x+1] > v {
		size += bf.getSize(x+1, y)
	}

	if y > 0 && bf.grid[y-1][x] > v {
		size += bf.getSize(x, y-1)
	}

	if y < len(bf.grid)-1 && bf.grid[y+1][x] > v {
		size += bf.getSize(x, y+1)
	}

	return size
}

type grid [][]int

func (g grid) isLowPoint(x, y int) bool {
	v := g[y][x]

	if x > 0 && v >= g[y][x-1] {
		return false
	}

	if x < len(g[y])-1 && v >= g[y][x+1] {
		return false
	}

	if y > 0 && v >= g[y-1][x] {
		return false
	}

	if y < len(g)-1 && v >= g[y+1][x] {
		return false
	}

	return true
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var (
		g          = make(grid, 0)
		line       = make([]int, 0)
		lineLength = 0
		part1      = 0
		basinSizes = make([]int, 0)
		r          = bufio.NewReader(fd)
	)

	for i := 0; ; i++ {
		i, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not read a byte: %v", err)
		}

		if i == '\n' {
			g = append(g, line)
			line = make([]int, lineLength)

			continue
		}

		if i == 0 {
			lineLength++
		}

		line = append(line, int(i-'0'))
	}

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g.isLowPoint(x, y) {
				part1 += g[y][x] + 1

				basinSizes = append(
					basinSizes,
					newBasinFinder(g).getSize(x, y),
				)
			}
		}
	}

	sort.Ints(basinSizes)

	part2 := basinSizes[len(basinSizes)-1]

	for i := 0; i < 2; i++ {
		part2 *= basinSizes[len(basinSizes)-2-i]
	}

	log.Print("Part 1: ", part1)
	log.Print("Part 2: ", part2)
}

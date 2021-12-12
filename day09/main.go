package main

import (
	"container/ring"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type grid [][]int

func (g grid) risks() int {
	s := 0

	for idx, i := range g[1] {
		if idx > 0 && g[1][idx] >= g[1][idx-1] {
			continue
		}

		if idx < len(g[1])-1 && g[1][idx] >= g[1][idx+1] {
			continue
		}

		if g[0] != nil && i >= g[0][idx] {
			continue
		}

		if g[2] != nil && i >= g[2][idx] {
			continue
		}

		s += i + 1
	}

	return s
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	const bufferSize = 3

	var (
		g     = make(grid, bufferSize)
		ints  []int
		line  string
		r     = ring.New(bufferSize)
		rInit = false
		s     = 0
	)

	for i := 0; ; i++ {
		if _, err := fmt.Fscanf(fd, "%s", &line); err != nil {
			if errors.Is(err, io.EOF) {
				g[0] = r.Prev().Prev().Value.([]int)
				g[1] = r.Prev().Value.([]int)
				g[2] = nil

				s += g.risks()

				break
			}

			log.Fatalf("Could not read line %d: %v", i+1, err)
		}

		if !rInit {
			for i := 0; i < bufferSize; i++ {
				r.Value = make([]int, len(line))
				r = r.Next()
			}

			rInit = true
		}

		ints = r.Value.([]int)

		for idx, c := range line {
			ints[idx] = int(c - '0')
		}

		if i == 1 {
			g[0] = nil
			g[1] = r.Prev().Value.([]int)

		} else if i > 1 {
			g[0] = r.Prev().Prev().Value.([]int)
			g[1] = r.Prev().Value.([]int)
		}

		g[2] = ints

		s += g.risks()

		r.Value = ints
		r = r.Next()
	}

	log.Printf("Part 1: %d", s)
}

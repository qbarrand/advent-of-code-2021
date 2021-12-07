package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type point struct {
	x int
	y int
}

func max(i, j int) int {
	if i >= j {
		return i
	}

	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}

	return j
}

func main() {
	cl, err := util.ParseCommandLine(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("Could not parse the command-line: %v", err)
	}

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var (
		verticalHorizontal = make(map[point]int)
		diagonal           = make(map[point]int)
		p1                 point
		p2                 point
	)

	for i := 1; ; i++ {
		if _, err := fmt.Fscanf(fd, "%d,%d -> %d,%d", &p1.x, &p1.y, &p2.x, &p2.y); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Error reading line %d: %v", i, err)
		}

		if p1.x == p2.x {
			for i := min(p1.y, p2.y); i <= max(p1.y, p2.y); i++ {
				p := point{p1.x, i}

				verticalHorizontal[p]++
				diagonal[p]++
			}
		}

		if p1.y == p2.y {
			for i := min(p1.x, p2.x); i <= max(p1.x, p2.x); i++ {
				p := point{i, p1.y}

				verticalHorizontal[p]++
				diagonal[p]++
			}
		}

		if p2.x != p1.x {
			slope := (p2.y - p1.y) / (p2.x - p1.x)

			if slope != 1 && slope != -1 {
				continue
			}

			var src, dst point

			if p1.x <= p2.x {
				src = p1
				dst = p2
			} else {
				src = p2
				dst = p1
			}

			inc := src.y <= dst.y

			for i := 0; src.x+i <= dst.x; i++ {
				y := src.y

				if inc {
					y += i
				} else {
					y -= i
				}

				diagonal[point{src.x + i, y}]++
			}
		}
	}

	part1 := 0
	part2 := 0

	for _, v := range verticalHorizontal {
		if v >= 2 {
			part1++
		}
	}

	for _, v := range diagonal {
		if v >= 2 {
			part2++
		}
	}

	log.Printf("Part 1: %d", part1)
	log.Printf("Part 2: %d", part2)
}
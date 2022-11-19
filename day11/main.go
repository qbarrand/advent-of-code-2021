package main

import (
	"bufio"
	"errors"
	"github.com/qbarrand/advent-of-code-2021/util"
	"io"
	"log"
)

const (
	height = 10
	width  = 10
)

type octopus struct {
	energyLevel int
	hasFlashed  bool
}

var matrix = [height][width]*octopus{}

func flashIfNeeded(h, w int) int {
	flashes := 0

	o := matrix[h][w]

	if o.hasFlashed {
		return flashes
	}

	if o.energyLevel > 9 {
		flashes++

		o.energyLevel = 0
		o.hasFlashed = true

		for y := h - 1; y <= h+1; y++ {
			for x := w - 1; x <= w+1; x++ {
				if y < 0 || y >= height || x < 0 || x >= width || (y == h && x == w) {
					continue
				}

				n := matrix[y][x]

				if n.hasFlashed {
					continue
				}

				n.energyLevel++

				flashes += flashIfNeeded(y, x)
			}
		}
	}

	return flashes
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)

	rd := bufio.NewReader(fd)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			r, _, err := rd.ReadRune()
			if err != nil {
				log.Fatalf("Error while reading: %v", err)
			}

			matrix[h][w] = &octopus{
				energyLevel: int(r - '0'),
			}
		}

		_, _, err := rd.ReadRune()
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatalf("Error while reading: %v", err)
		}
	}

	if err := fd.Close(); err != nil {
		log.Fatalf("Could not close the input file: %v", err)
	}

	const part1Steps = 100

	var totalFlashes = 0

	for i := 1; ; i++ {
		flashes := 0

		// First pass: increase everyone by 1
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				matrix[h][w].energyLevel++
				matrix[h][w].hasFlashed = false
			}
		}

		// Second pass: recursively flash octopuses if needed
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				flashes += flashIfNeeded(h, w)
			}
		}

		totalFlashes += flashes

		if i == part1Steps {
			log.Printf("Part 1: %d", totalFlashes)
		}

		if flashes == height*width {
			log.Printf("Part 2: %d", i)

			if i > part1Steps {
				return
			}
		}
	}
}

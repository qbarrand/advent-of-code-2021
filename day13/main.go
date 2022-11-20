package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type coord struct {
	x, y int
}

func readCoordinates(s string) (*coord, error) {
	elems := strings.Split(strings.TrimSuffix(s, "\n"), ",")

	x, err := strconv.Atoi(elems[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer: %v", elems[0], err)
	}

	y, err := strconv.Atoi(elems[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse %q as integer: %v", elems[1], err)
	}

	return &coord{x: x, y: y}, nil
}

func readFoldInstructions(s string) (*coord, error) {
	bsr := bufio.NewReader(strings.NewReader(s))

	beforeEqual, err := bsr.ReadString('=')
	if err != nil {
		log.Fatalf("Could not read the first part of the fold line: %v", err)
	}

	var val int

	if _, err = fmt.Fscanf(bsr, "%d\n", &val); err != nil {
		return nil, fmt.Errorf("could not read the value: %v", err)
	}

	switch d := beforeEqual[len(beforeEqual)-2]; d {
	case 'x':
		return &coord{x: val}, nil
	case 'y':
		return &coord{y: val}, nil
	default:
		return nil, fmt.Errorf("%q: invalid dimension", d)
	}
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)

	r := bufio.NewReader(fd)

	coords := make(map[coord]struct{})
	instructions := make([]*coord, 0)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Error reading the file: %v", err)
		}

		switch {
		case strings.Contains(s, ","): // coordinates
			c, err := readCoordinates(s)
			if err != nil {
				log.Fatalf("Could not read coordinates: %v", err)
			}

			coords[*c] = struct{}{}
		case len(s) <= 1: // separation line between coordinates and fold instructions, or last line
		case strings.HasPrefix(s, "fold"):
			c, err := readFoldInstructions(s)
			if err != nil {
				log.Fatalf("could not read fold instructions: %v", err)
			}

			instructions = append(instructions, c)
		default:
			log.Fatalf("%q: could not process line", s)
		}
	}

	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]

		for c := range coords {
			if f := inst.x; f != 0 && c.x >= f {
				newCoord := coord{x: f - (c.x - f), y: c.y}

				coords[newCoord] = struct{}{}
				delete(coords, c)
				continue
			}

			if f := inst.y; f != 0 && c.y >= f {
				newCoord := coord{x: c.x, y: f - (c.y - f)}

				coords[newCoord] = struct{}{}
				delete(coords, c)
				continue
			}
		}

		if i == 0 {
			log.Println("Part 1:", len(coords))
		}
	}

	var maxX, maxY int

	for k := range coords {
		if k.x > maxX {
			maxX = k.x
		}

		if k.y > maxY {
			maxY = k.y
		}
	}

	log.Println("Part 2:")

	w := bufio.NewWriter(log.Writer())

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, ok := coords[coord{x: x, y: y}]; ok {
				w.WriteByte('#')
			} else {
				w.WriteByte(' ')
			}
		}

		w.WriteByte('\n')
	}

	w.Flush()
}

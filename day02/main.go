package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/qbarrand/advent-of-code-2021/util"
)

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var (
		aim       int
		direction string
		distance  int
		p1depth   int
		p2depth   int
		pos       int
	)

	for i := 1; ; i++ {
		_, err := fmt.Fscanf(fd, "%s %d", &direction, &distance)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not read line %d: %v", i, err)
		}

		switch direction {
		case "forward":
			pos += distance
			p2depth += aim * distance
		case "up":
			p1depth -= distance
			aim -= distance
		case "down":
			p1depth += distance
			aim += distance
		}
	}

	log.Printf("Part 1: %d", p1depth*pos)
	log.Printf("Part 2: %d", p2depth*pos)
}

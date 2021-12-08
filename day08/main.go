package main

import (
	"bufio"
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

	var w0, w1, w2, w3 string

	r := bufio.NewReader(fd)

	part1 := 0

	for {
		if _, err := r.ReadString('|'); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not consume the data prior to pipe: %v", err)
		}

		if _, err := fmt.Fscanf(r, "%s %s %s %s", &w0, &w1, &w2, &w3); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not scan: %v", err)
		}

		for _, w := range []string{w0, w1, w2, w3} {
			L := len(w)

			if L == 2 || L == 3 || L == 4 || L == 7 {
				part1 += 1
			}
		}
	}

	log.Print("Part 1: ", part1)
}

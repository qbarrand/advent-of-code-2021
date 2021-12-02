package main

import (
	"container/ring"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/qbarrand/advent-of-code-2021/util"
)

func main() {
	cli, err := util.ParseCommandLine(os.Args[0], os.Args[1:])
	if err != nil {
		os.Exit(1)
		log.Fatalf("Could not parse the command line: %v", err)
	}

	fd, err := os.Open(cli.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	var (
		inc  int
		r    = ring.New(4)
		v    int
		wInc int
	)

	for i := 1; ; i++ {
		_, err = fmt.Fscanf(fd, "%d", &v)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Failed while reading line %d: %v", i, err)
		}

		r.Value = v

		if i > 1 && r.Prev().Value.(int) < v {
			inc++
		}

		if i >= 4 {
			pcw := r
			cw := v

			ppw := r.Prev()
			pw := ppw.Value.(int)

			for j := 0; j < 2; j++ {
				pcw = pcw.Prev()
				cw += pcw.Value.(int)

				ppw = ppw.Prev()
				pw += ppw.Value.(int)
			}

			if cw > pw {
				wInc++
			}
		}

		r = r.Next()
	}

	log.Printf("Part 1: %d", inc)
	log.Printf("Part 2: %d", wInc)
}

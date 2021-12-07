package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type bitColumn struct {
	ones   int
	zeroes int
}

func (bc bitColumn) leastCommon() int {
	if bc.ones > bc.zeroes {
		return 0
	}

	return 1
}

func (bc bitColumn) mostCommon() int {
	if bc.ones > bc.zeroes {
		return 1
	}

	return 0
}

type node struct {
	branchCount int
	childZero   *node
	childOne    *node
}

func (n *node) addNumber(num, path string) {
	n.branchCount++

	if len(path) == 0 {
		return
	}

	if path[0] == '0' {
		if n.childZero == nil {
			n.childZero = &node{}
		}

		n.childZero.addNumber(num, path[1:])
	} else {
		if n.childOne == nil {
			n.childOne = &node{}
		}

		n.childOne.addNumber(num, path[1:])
	}
}

func parseLine(cols []bitColumn, line string) {
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '0':
			cols[i].zeroes++
		case '1':
			cols[i].ones++
		default:
			log.Fatalf("%c: invalid character", line[i])
		}
	}
}

func pow(a, b int) int {
	ret := 1

	for i := 1; i <= b; i++ {
		ret *= a
	}

	return ret
}

func main() {
	cl, err := util.ParseCommandLine(os.Args[0], os.Args[1:])
	if err != nil {
		log.Fatalf("Could not parse the command line: %v", err)
	}

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var (
		line string
		tree = &node{}
	)

	if _, err = fmt.Fscanf(fd, "%s", &line); err != nil {
		log.Fatalf("Could not read the first line: %v", err)
	}

	tree.addNumber(line, line)

	L := len(line)

	cols := make([]bitColumn, L)

	parseLine(cols, line)

	for i := 2; ; i++ {
		if _, err = fmt.Fscanf(fd, "%s", &line); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Could not read line %d: %v", i, err)
		}

		parseLine(cols, line)
		tree.addNumber(line, line)
	}

	var gamma, epsilon int

	j := 0

	for i := L - 1; i >= 0; i-- {
		gamma += cols[i].mostCommon() * pow(2, j)
		epsilon += cols[i].leastCommon() * pow(2, j)

		j++
	}

	log.Printf("Part 1: %d", gamma*epsilon)

	var (
		oString, co2String string
		n                  = tree
	)

	for {
		nZero := 0
		nOne := 0

		if n.childZero != nil {
			nZero = n.childZero.branchCount
		}

		if n.childOne != nil {
			nOne = n.childOne.branchCount
		}

		if nZero == nOne {
			if nZero == 0 {
				break
			}

			oString += "1"
			n = n.childOne
		} else if nZero > nOne {
			oString += "0"
			n = n.childZero
		} else {
			oString += "1"
			n = n.childOne
		}
	}

	o, err := strconv.ParseInt(oString, 2, 32)
	if err != nil {
		log.Fatalf("Could not parse %s as binary: %v", oString, err)
	}

	n = tree

	for {
		nZero := math.MaxInt32
		nOne := math.MaxInt32

		if n.childZero != nil {
			nZero = n.childZero.branchCount
		}

		if n.childOne != nil {
			nOne = n.childOne.branchCount
		}

		if nZero == nOne {
			if nZero == math.MaxInt32 {
				break
			}

			co2String += "0"
			n = n.childZero
		} else if nZero > nOne {
			co2String += "1"
			n = n.childOne
		} else {
			co2String += "0"
			n = n.childZero
		}
	}

	co2, err := strconv.ParseInt(co2String, 2, 32)
	if err != nil {
		log.Fatalf("Could not parse %s as binary: %v", co2String, err)
	}

	log.Printf("Part 2: %d", o*co2)
}

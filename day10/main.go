package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"sort"

	"github.com/qbarrand/advent-of-code-2021/util"
)

var (
	pairs = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	mapPart1Scores = map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	mapPart2Scores = map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

type node struct {
	b    byte
	next *node
}

type stack struct {
	top *node
}

func (s *stack) clear() {
	s.top = nil
}

func (s *stack) pop() (byte, error) {
	if s.top == nil {
		return 0, errors.New("stack empty")
	}

	top := s.top.b
	s.top = s.top.next

	return top, nil
}

func (s *stack) put(b byte) {
	s.top = &node{b: b, next: s.top}
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)
	defer fd.Close()

	var (
		r           = bufio.NewReader(fd)
		part1       = 0
		part2Scores = make([]int, 0)
		st          = &stack{}
	)

	for {
		b, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Error while reading: %v", err)
		}

		if b == '\n' {
			score := 0

			for {
				b, err = st.pop()
				if err != nil {
					break
				}

				score *= 5
				score += mapPart2Scores[b]
			}

			part2Scores = append(part2Scores, score)

			st.clear()
			continue
		}

		if pairs[b] != 0 {
			st.put(b)
		} else {
			open, err := st.pop()
			if err != nil {
				log.Fatalf("could not pop() from the stack: %v", err)
			}

			if pairs[open] != b {
				part1 += mapPart1Scores[b]

				if _, err := r.ReadBytes('\n'); err != nil {
					log.Fatalf("could not read till the next line: %v", err)
				}

				st.clear()
				continue
			}
		}
	}

	sort.Ints(part2Scores)

	log.Print("Part 1: ", part1)
	log.Print("Part 2: ", part2Scores[len(part2Scores)/2])
}

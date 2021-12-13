package main

import (
	"bufio"
	"errors"
	"io"
	"log"

	"github.com/qbarrand/advent-of-code-2021/util"
)

var pairs = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

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
		r     = bufio.NewReader(fd)
		score = 0
		st    = &stack{}
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
				score += scores[b]

				_, err := r.ReadBytes('\n')
				if err != nil {
					log.Fatalf("could not read till the next line: %v", err)
				}
			}
		}
	}

	log.Print("Part 1: ", score)
}

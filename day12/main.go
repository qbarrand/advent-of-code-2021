package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strings"
	"unicode"

	"github.com/qbarrand/advent-of-code-2021/util"
)

type listNode struct {
	name string
	prev *listNode
}

func (l *listNode) containsNode(name string) bool {
	ln := l

	for {
		if ln == nil {
			break
		}

		if ln.name == name {
			return true
		}

		ln = ln.prev
	}

	return false
}

type node = map[string]struct{}

var nodes = make(map[string]node)

const (
	end   = "end"
	start = "start"
)

func findAllPaths(this string, prev *listNode, remainingDoubleVisits int) int {
	if this == end {
		return 1
	}

	if unicode.IsLower(rune(this[0])) && prev.containsNode(this) {
		if this == start || remainingDoubleVisits == 0 {
			return 0
		}

		remainingDoubleVisits--
	}

	currentList := &listNode{
		name: this,
		prev: prev,
	}

	paths := 0

	for k := range nodes[this] {
		paths += findAllPaths(k, currentList, remainingDoubleVisits)
	}

	return paths
}

func main() {
	cl := util.ParseCommandLine()

	fd := util.MustOpen(cl.InputFile)

	s := bufio.NewReader(fd)

	for {
		left, err := s.ReadString('-')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			log.Fatalf("Error while reading the first token: %v", err)
		}

		left = strings.TrimSuffix(left, "-")

		right, err := s.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatalf("Error while reading the second token: %v", err)
		}

		right = strings.TrimSuffix(right, "\n")

		if _, ok := nodes[left]; !ok {
			nodes[left] = make(node)
		}

		nodes[left][right] = struct{}{}

		if _, ok := nodes[right]; !ok {
			nodes[right] = make(node)
		}

		nodes[right][left] = struct{}{}
	}

	log.Println(
		"Part 1:",
		findAllPaths(start, nil, 0),
	)

	log.Println(
		"Part 2:",
		findAllPaths(start, nil, 1),
	)
}

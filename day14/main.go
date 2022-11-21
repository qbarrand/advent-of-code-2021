package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type cacheKey struct {
	name string
	iter int
}

type stat = map[rune]int

var (
	cache = make(map[cacheKey]stat)
	rules = make(map[string]rune)
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	if !s.Scan() {
		log.Fatalf("Could not scane the template: %v", s.Err())
	}

	template := s.Text()

	s.Scan()

	for s.Scan() {
		elems := strings.Split(s.Text(), " -> ")

		rules[elems[0]] = rune(elems[1][0])
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while scanning: %v", err)
	}

	m := getElementMap(template, 10)

	log.Println("Part 1:", findResult(m))

	m = getElementMap(template, 40)

	log.Println("Part 2:", findResult(m))
}

func findResult(m stat) int {
	var (
		max = 0
		min *int
	)

	for _, v := range m {
		if v > max {
			max = v
		}

		if min == nil || v < *min {
			v := v
			min = &v
		}
	}

	return max - *min
}

func getElementMap(template string, steps int) stat {
	key := cacheKey{
		name: template,
		iter: steps,
	}

	if m := cache[key]; m != nil {
		return m
	}

	m := make(stat)

	if steps == 0 {
		for _, c := range template {
			m[c]++
		}

		return m
	}

	for i := 1; i < len(template); i++ {
		pair := template[i-1 : i+1]
		triple := []rune{rune(pair[0]), rules[pair], rune(pair[1])}

		// Do not count doubles
		if i != len(template)-1 {
			m[rune(pair[1])]--
		}

		for k, v := range getElementMap(string(triple), steps-1) {
			m[k] += v
		}
	}

	cache[key] = m

	return m
}

package main

import (
	"bufio"
	"container/heap"
	"errors"
	"io"
	"log"
	"math"
	"os"
)

type coord struct{ x, y int }

type vertex struct {
	c         coord
	dist      int
	index     int
	riskLevel int
	visited   bool
}

type djikstra struct {
	g  graph
	pq *PQ
}

var neighCoords = [4]coord{}

func (d *djikstra) getLowestRisk() int {
	d.g[0][0].dist = 0

	for d.pq.Len() > 0 {
		v := heap.Pop(d.pq).(*vertex)
		v.visited = true

		neighCoords[0] = coord{x: v.c.x, y: v.c.y - 1} // top
		neighCoords[1] = coord{x: v.c.x - 1, y: v.c.y} // left
		neighCoords[2] = coord{x: v.c.x + 1, y: v.c.y} // right
		neighCoords[3] = coord{x: v.c.x, y: v.c.y + 1} // bottom

		for _, nc := range neighCoords {
			if n := d.g.getUnvisitedVertex(nc.x, nc.y); n != nil {
				if newDist := v.dist + n.riskLevel; newDist < n.dist {
					d.pq.update(n, newDist)
				}
			}
		}
	}

	return d.g[d.g.height()-1][d.g.width()-1].dist
}

func newDjikstra(orig [][]int, factor int) *djikstra {
	height := len(orig) * factor
	width := len(orig[0]) * factor

	g := make([][]*vertex, 0, height)
	pq := make(PQ, 0, height*width)

	i := 0

	for y := 0; y < height; y++ {
		line := make([]*vertex, 0, width)

		for x := 0; x < width; x++ {
			v := vertex{
				c:         coord{x: x, y: y},
				dist:      math.MaxInt32,
				index:     i,
				riskLevel: getRiskLevel(orig, x, y),
			}

			line = append(line, &v)
			heap.Push(&pq, &v)

			i++
		}

		g = append(g, line)
	}

	heap.Init(&pq)

	return &djikstra{g: g, pq: &pq}
}

type graph [][]*vertex

func (g graph) bottomRight() *vertex {
	return g[g.height()-1][g.width()-1]
}

func (g graph) height() int {
	return len(g)
}

func (g graph) getUnvisitedVertex(x, y int) *vertex {
	if x < 0 || y < 0 || x >= g.width() || y >= g.height() {
		return nil
	}

	v := g[y][x]

	if v.visited {
		return nil
	}

	return v
}

func (g graph) width() int {
	return len(g[0])
}

func getRiskLevel(orig [][]int, x, y int) int {
	origHeight := len(orig)
	origWidth := len(orig[0])

	factorX := x / origWidth
	factorY := y / origHeight

	risk := orig[y%origHeight][x%origWidth]

	risk += factorX + factorY

	if risk > 9 {
		risk %= 9
	}

	return risk
}

type PQ []*vertex

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	// We want the smallest dist (= vertices with the lowest dist) at the end of the slice
	return pq[i].dist < pq[j].dist
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	item := x.(*vertex)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the distance of an item in the queue.
func (pq *PQ) update(v *vertex, dist int) {
	v.dist = dist
	heap.Fix(pq, v.index)
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	orig := make([][]int, 0)

	for y := 0; s.Scan(); y++ {
		str := s.Text()

		line := make([]int, 0)

		for _, c := range str {
			line = append(line, int(c)-'0')
		}

		orig = append(orig, line)
	}

	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error while reading: %v", err)
	}

	log.Println("Part 1:", newDjikstra(orig, 1).getLowestRisk())
	log.Println("Part 2:", newDjikstra(orig, 5).getLowestRisk())
}

package dijkstra

import (
	"container/heap"

	"github.com/kr/pretty"
)

const inf = 1 << 30

type pqItem struct {
	value      int
	graphIndex int
	pqIndex    int
	processed  bool
}
type priorityQueue []*pqItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value < pq[j].value
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].pqIndex = i
	pq[j].pqIndex = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqItem)
	item.pqIndex = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	pretty.Println(pq)
	old := *pq
	n := len(old)
	item := old[n-1]
	item.pqIndex = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds shortest paths from given nodes to the others in the graph
// Graph is represented like this:
// - index of array `graph` is an index of the node
// - each value of array `graph` is an array of neighbours of the node
// (index of the node is an index of `graph` array, remember?)
// - each neighbour is represented by 2-element array where 0-th element - index of neighbour node,
// 1-st element - weight of the edge to that neighbour
//
// returns array of path lengths from given node with index `start` to any other node
func Dijkstra(graph [][][2]int, start int) []int {
	pq := make(priorityQueue, len(graph))
	ref := make([]*pqItem, len(graph))
	for i := range pq {
		pq[i] = &pqItem{
			graphIndex: i,
			pqIndex:    i,
			value:      inf,
			processed:  false,
		}
	}
	pq[start].value = 0
	copy(ref, pq)
	heap.Init(&pq)

	for len(pq) > 0 {
		node := heap.Pop(&pq).(*pqItem)

		node.processed = true
		for _, edge := range graph[node.graphIndex] {
			neighbour := ref[edge[0]]
			if neighbour.processed {
				continue // protection against bi-directional edges
			}

			weight := node.value + edge[1]
			if weight < neighbour.value {
				neighbour.value = weight
				heap.Fix(&pq, neighbour.pqIndex)
			}
		}
	}

	result := make([]int, len(graph))
	for i := range ref {
		result[i] = ref[i].value
	}

	return result
}

package dijkstra

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := [][][2]int{
		{ // 0
			{1, 7},
			{2, 9},
			{5, 14},
		},
		{ // 1
			{0, 7},
			{2, 10},
			{3, 15},
		},
		{ // 2
			{0, 9},
			{1, 10},
			{5, 2},
			{3, 11},
		},
		{ // 3
			{1, 15},
			{4, 6},
			{2, 11},
		},
		{ // 4
			{5, 9},
			{3, 6},
		},
		{ // 5
			{0, 14},
			{2, 2},
			{4, 9},
		},
	}

	result := Dijkstra(graph, 0)
	expected := []int{0, 7, 9, 20, 20, 11}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected: %v\nGot: %v", expected, result)
	}

	result = Dijkstra(graph, 1)
	expected = []int{7, 0, 10, 15, 21, 12}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected: %v\nGot: %v", expected, result)
	}

}

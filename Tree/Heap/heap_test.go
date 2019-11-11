package Heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap(10)
	a := []int{5, 4, 1, 6, 7, 3, 2, 8, 9, 0}
	for i := range a {
		heap.inster(i)
	}
	fmt.Println(heap.base)

	heap.remove()
}

func TestHeap2(t *testing.T) {
	heap := NewHeap(10)
	a := []int{5, 4, 1, 6, 7, 3, 2, 8, 9, 0}
	for i := range a {
		heap.inster(a[i])

	}

	fmt.Println(heap.base)
	heap.remove()
	fmt.Println(heap.base)
}

func TestHeapBuild(t *testing.T) {
	a := []int{0, 5, 4, 1, 6, 7, 3, 2, 8, 9}
	heap := NewHeap(len(a))
	heap.base = a
	heap.cap = 9
	heap.buildHeap()
	fmt.Println(heap.base)
}

package main

import (
	"container/heap"
	"fmt"
)

type cursor struct {
	val, arr, idx int
}

type cursorHeap []cursor

func (h cursorHeap) Len() int           { return len(h) }
func (h cursorHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h cursorHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cursorHeap) Push(x any)        { *h = append(*h, x.(cursor)) }
func (h *cursorHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// MergeK merges k individually sorted slices into one sorted slice. A
// min-heap holds one cursor per array; popping the smallest advances only
// that array.
func MergeK(arrays [][]int) []int {
	h := &cursorHeap{}
	total := 0
	for i, a := range arrays {
		total += len(a)
		if len(a) > 0 {
			*h = append(*h, cursor{a[0], i, 0})
		}
	}
	heap.Init(h)
	out := make([]int, 0, total)
	for h.Len() > 0 {
		c := heap.Pop(h).(cursor)
		out = append(out, c.val)
		if next := c.idx + 1; next < len(arrays[c.arr]) {
			heap.Push(h, cursor{arrays[c.arr][next], c.arr, next})
		}
	}
	return out
}

func main() {
	fmt.Println(MergeK([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}))
}

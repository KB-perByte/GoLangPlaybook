package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.MaxHeapifyDown(0)

	return extracted
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30}
	for _, v := range buildHeap {
		m.Insert(v)
	}
}

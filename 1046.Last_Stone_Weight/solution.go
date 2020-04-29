package main

import "fmt"

type key int
type pri int

const (
	Priority_EQ pri = 0
	Priority_HI pri = 1
	Priority_LO pri = 2
)

type IPriority interface {
	Compare(a int, b int) pri
}

type Heap struct {
	priority IPriority
	size     uint
	store    []int
}

func (h *Heap) IsEmpty() bool {
	fmt.Println("size", h.size)
	return h.size == 0
}

func (h *Heap) Top() int {
	return h.store[0]
}

func (h *Heap) Insert(v int) {
	h.store = append(h.store, v)
	h.size++
	i := key(h.size - 1)
	isHi := h.priority.Compare(h.store[h.parent(i)], h.store[i]) != Priority_LO
	for i > 0 && !isHi {
		h.swap(h.parent(i), i)
		i = h.parent(i)
	}
}

func (h *Heap) ExtractTop() int {
	if h.size == 0 {
		return -1
	}
	top := h.store[0]
	h.store[0] = h.store[h.size-1]
	h.size--
	h.heapify(0)

	return top
}

func (h *Heap) heapify(hi key) {
	l := h.left(hi)
	r := h.right(hi)
	_hi := hi
	if h.priority.Compare(h.store[hi], h.store[l]) == Priority_LO {
		_hi = l
	}
	if h.priority.Compare(h.store[hi], h.store[r]) == Priority_LO {
		_hi = r
	}
	if _hi != hi {
		h.swap(_hi, hi)
		h.heapify(hi)
	}
}

func (h *Heap) parent(i key) key {
	p := (i - 1) >> 1
	if p >= 0 {
		return p
	}

	return 0
}

func (h *Heap) left(i key) key {
	return i<<1 + 1
}

func (h *Heap) right(i key) key {
	return h.left(i) + 1
}

func (h *Heap) swap(a key, b key) {
	tmp := h.store[a]
	h.store[a] = h.store[b]
	h.store[b] = tmp
}

func (h *Heap) initHeap(vals *[]int) {
	for _, v := range *vals {
		h.Insert(v)
	}
}

type MaxHeap struct {
	IPriority
}

func (h *MaxHeap) Compare(a int, b int) pri {
	if a > b {
		return Priority_HI
	} else if a == b {
		return Priority_EQ
	}

	return Priority_LO
}

func NewHeap(vals []int, priority IPriority) *Heap {
	h := &Heap{
		priority: priority,
		size:     0,
		store:    make([]int, 0, 1000),
	}
	for _, v := range vals {
		h.Insert(v)
	}

	return h
}

func lastStoneWeight(stones []int) int {
	maxHeap := NewHeap(stones, new(MaxHeap))
	for maxHeap.IsEmpty() == false {
		fmt.Println("Before Extract", maxHeap.store)
		h1 := maxHeap.ExtractTop()
		fmt.Println(h1, "Extracted", maxHeap.store)
		if maxHeap.IsEmpty() {
			return h1
		}
		h2 := maxHeap.ExtractTop()
		fmt.Println(h2, "Extracted", maxHeap.store)
		if h1 == h2 {
			continue
		} else {
			maxHeap.Insert(h1 - h2)
		}
	}
	return 0
}

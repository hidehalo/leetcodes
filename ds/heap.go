package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type key int
type val int
type pri int

const (
	Priority_EQ pri = 0
	Priority_HI pri = 1
	Priority_LO pri = 2
)

type IHeap interface {
	Compare(a val, b val) pri
	Top() val
	Insert(v val)
	ExtractTop()
	String()
}

type Heap struct {
	IHeap
	size  uint
	store []val
}

type MinHeap struct {
	Heap
}

type MaxHeap struct {
	Heap
}

func (h *Heap) Top() val {
	return h.store[0]
}

func (h *Heap) Insert(v val) {
	h.store = append(h.store, v)
	h.size++
	i := key(h.size - 1)
	fmt.Println(*h)
	isHi := h.Compare(h.store[h.parent(i)], h.store[i]) == Priority_HI
	for i > 0 && !isHi {
		h.swap(h.parent(i), i)
		i = h.parent(i)
	}
}

func (h *Heap) ExtractTop() val {
	if h.size >= 1 {
		top := h.store[0]
		h.store[0] = h.store[h.size-1]
		h.size--
		h.heapify(0)

		return top
	}

	return h.store[0]
}

func (h *Heap) heapify(hi key) {
	l := h.left(hi)
	r := h.right(hi)
	_hi := hi
	if h.Compare(h.store[hi], h.store[l]) == Priority_LO {
		_hi = l
	}
	if h.Compare(h.store[hi], h.store[r]) == Priority_LO {
		_hi = r
	}
	if _hi != hi {
		h.swap(_hi, hi)
		h.heapify(hi)
	}
}

func (h *Heap) parent(i key) key {
	return (i - 1) << 1
}

func (h *Heap) left(i key) key {
	return i>>1 + 1
}

func (h *Heap) right(i key) key {
	return h.left(i) + 1
}

func (h *Heap) String() string {
	var strBuf bytes.Buffer
	i := 0
	for i < int(h.size) {
		strBuf.WriteString(strconv.FormatInt(int64(h.store[i]), 32))
		i++
	}

	return strBuf.String()
}

func (h *Heap) swap(a key, b key) {
	tmp := h.store[a]
	h.store[a] = h.store[b]
	h.store[b] = tmp
}

func (h *Heap) initHeap(vals *[]val) {
	for _, v := range *vals {
		h.Insert(v)
	}
}

func NewHeap(vals []val) Heap {
	h := Heap{
		size:  0,
		store: make([]val, 0, 1000),
	}

	return h
}

func NewMinHeap(vals []val) *MinHeap {
	h := &MinHeap{
		Heap: NewHeap(vals),
	}
	h.initHeap(&vals)

	return h
}

func NewMaxHeap(vals []val) *MaxHeap {
	h := &MaxHeap{
		Heap: NewHeap(vals),
	}
	h.initHeap(&vals)

	return h
}

func (h *MinHeap) Compare(a val, b val) pri {
	if a < b {
		return Priority_HI
	}

	return Priority_LO
}

func (h *MaxHeap) Compare(a val, b val) pri {
	if a > b {
		return Priority_HI
	}

	return Priority_LO
}

func main() {
	vals := []val{1, 2, 3, 4, 5}
	maxHeap := NewMaxHeap(vals)
	minHeap := NewMinHeap(vals)
	fmt.Println(maxHeap)
	fmt.Println(minHeap)
}

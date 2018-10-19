package main

type key int
type val uint
type pri int

const (
	Priority_EQ pri = 0
	Priority_HI pri = 1
	Priority_LO pri = 2
)

type Heap struct {
	size  uint
	store []val
}

func (h *Heap) Top() val {
	return h.store[0]
}

func (h *Heap) Insert(v val) {
	h.size++
	h.store[h.size-1] = v
}

func (h *Heap) ExtractTop() val {
	// bottom to top
	// size--
	// heapify
	return h.store[0]
}

func (h *Heap) Compare(a val, b val) pri {
	return Priority_EQ
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

func (h *Heap) swap(a key, b key) {
	tmp := h.store[a]
	h.store[a] = h.store[b]
	h.store[b] = tmp
}

func main() {

}

package ds

const (
	PriorityEQ = 0
	PriorityHI = 1
	PriorityLO = 2
)

type IPriority interface {
	Compare(a int, b int) int
}

type Heap struct {
	priority IPriority
	size     int
	store    []int
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) Top() int {
	if h.size == 0 {
		return -1
	}
	return h.store[0]
}

func (h *Heap) Insert(v int) int {
	if h.size >= len(h.store) {
		h.store = append(h.store, v)
	} else {
		h.store[h.size] = v
	}
	i := h.size
	h.size++
	for i > 0 {
		isLow := h.priority.Compare(h.store[h.parent(i)], h.store[i]) == PriorityLO
		if !isLow {
			break
		}
		h.swap(h.parent(i), i)
		i = h.parent(i)
	}

	return i
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

func (h *Heap) heapify(hi int) {
	if hi >= h.size-1 {
		return
	}
	l := h.left(hi)
	r := h.right(hi)
	newHi := hi
	if l < h.size && h.priority.Compare(h.store[newHi], h.store[l]) == PriorityLO {
		newHi = l
	}
	if r < h.size && h.priority.Compare(h.store[newHi], h.store[r]) == PriorityLO {
		newHi = r
	}
	if newHi != hi {
		h.swap(newHi, hi)
		h.heapify(newHi)
	}
}

func (h *Heap) parent(i int) int {
	p := (i - 1) >> 1
	if p >= 0 {
		return p
	}

	return 0
}

func (h *Heap) left(i int) int {
	return h.right(i) - 1
}

func (h *Heap) right(i int) int {
	return (i + 1) << 1
}

func (h *Heap) swap(a int, b int) {
	tmp := h.store[a]
	h.store[a] = h.store[b]
	h.store[b] = tmp
}

func NewHeap(vals []int, priority IPriority) *Heap {
	h := &Heap{
		priority: priority,
		size:     len(vals),
		store:    make([]int, int(len(vals))),
	}
	copy(h.store, vals)
	for i := h.parent(h.size - 1); i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

type MaxHeap struct{ IPriority }

func (h *MaxHeap) Compare(a int, b int) int {
	if a > b {
		return PriorityHI
	} else if a == b {
		return PriorityEQ
	}

	return PriorityLO
}

type MinHeap struct{ IPriority }

func (h *MinHeap) Compare(a int, b int) int {
	if a < b {
		return PriorityHI
	} else if a == b {
		return PriorityEQ
	}

	return PriorityLO
}

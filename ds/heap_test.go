package ds

import (
	"reflect"
	"sort"
	"testing"
)

func TestHeap_ExtractTop(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				0,
				[]int{},
			},
			-1,
		},
		{
			"case2",
			fields{
				&MinHeap{},
				0,
				[]int{},
			},
			-1,
		},
		{
			"case3",
			fields{
				&MaxHeap{},
				3,
				[]int{3, 2, 1},
			},
			3,
		},
		{
			"case4",
			fields{
				&MinHeap{},
				3,
				[]int{1, 3, 2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			if got := h.ExtractTop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap.ExtractTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Insert(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				0,
				[]int{},
			},
			args{1},
			0,
		},
		{
			"case2",
			fields{
				&MaxHeap{},
				5,
				[]int{10, 10, 5, 7, 8},
			},
			args{4},
			4,
		},
		{
			"case3",
			fields{
				&MinHeap{},
				4,
				[]int{1, 3, 5, 7},
			},
			args{2},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			h.Insert(tt.args.v)
		})
	}
}

func TestHeap_Top(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				2,
				[]int{5, 1},
			},
			5,
		},
		{
			"case2",
			fields{
				&MaxHeap{},
				0,
				[]int{},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			if got := h.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_IsEmpty(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				0,
				nil,
			},
			true,
		},
		{
			"case2",
			fields{
				&MaxHeap{},
				1,
				[]int{1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("Heap.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeap_Compare(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0, 1},
			2,
		},
		{
			"case2",
			args{1, 1},
			0,
		},
		{
			"case3",
			args{1, 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MaxHeap{}
			if got := h.Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxHeap.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Compare(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0, 1},
			1,
		},
		{
			"case2",
			args{1, 1},
			0,
		},
		{
			"case3",
			args{1, 0},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{}
			if got := h.Compare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinHeap.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHeap(t *testing.T) {
	type args struct {
		vals     []int
		priority IPriority
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{
				[]int{6, 1, 3, 2, 8, 4, 6, 10, 9, 7},
				&MaxHeap{},
			},
			[]int{10, 9, 6, 6, 8, 4, 3, 2, 1, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeap(tt.args.vals, tt.args.priority).store; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapPersistence(t *testing.T) {
	A := []int{6, 1, 3, 2, 8, 4, 6, 10, 9, 7}
	B := make([]int, len(A))
	copy(B, A)
	sort.Ints(B)
	maxHeap := NewHeap(A, &MaxHeap{})
	i := 0
	for !maxHeap.IsEmpty() {
		max := maxHeap.ExtractTop()
		i++
		if max != B[len(B)-i] {
			t.Errorf("MaxHeap.ExtractTop() %d time = %v, want %v", i, max, B[len(B)-i])
		}
	}
	i = 0
	minHeap := NewHeap(A, &MinHeap{})
	for !minHeap.IsEmpty() {
		min := minHeap.ExtractTop()
		i++
		if min != B[i-1] {
			t.Errorf("MinHeap.ExtractTop() %d time = %v, want %v", i, min, B[i-1])
		}
	}
}

func TestHeap_swap(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				4,
				[]int{4, 3, 2, 1},
			},
			args{0, 3},
			[]int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			h.swap(tt.args.a, tt.args.b)
			if got := h.store; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap.parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_parent(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0},
			0,
		},
		{
			"case2",
			args{1},
			0,
		},
		{
			"case3",
			args{2},
			0,
		},
		{
			"case4",
			args{5},
			2,
		},
		{
			"case5",
			args{6},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{}
			if got := h.parent(tt.args.i); got != tt.want {
				t.Errorf("Heap.parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_left(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0},
			1,
		},
		{
			"case2",
			args{1},
			3,
		},
		{
			"case3",
			args{2},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{}
			if got := h.left(tt.args.i); got != tt.want {
				t.Errorf("Heap.left() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_right(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{0},
			2,
		},
		{
			"case2",
			args{2},
			6,
		},
		{
			"case3",
			args{1},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{}
			if got := h.right(tt.args.i); got != tt.want {
				t.Errorf("Heap.right() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_heapify(t *testing.T) {
	type fields struct {
		priority IPriority
		size     int
		store    []int
	}
	type args struct {
		hi int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			"case1",
			fields{
				&MaxHeap{},
				5,
				[]int{1, 2, 3, 4, 5},
			},
			args{1},
			[]int{1, 5, 3, 4, 2},
		},
		{
			"case2",
			fields{
				&MinHeap{},
				5,
				[]int{5, 4, 3, 2, 1},
			},
			args{1},
			[]int{5, 1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Heap{
				priority: tt.fields.priority,
				size:     tt.fields.size,
				store:    tt.fields.store,
			}
			h.heapify(tt.args.hi)
			if got := h.store; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap.heapify() make Heap.store = %v, want %v", got, tt.want)
			}
		})
	}
}

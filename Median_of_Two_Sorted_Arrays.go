package main

import (
	"fmt"
)

type Pointer struct {
	Arr   *[]int
	Index int
	Tail  int
}

func (p *Pointer) Value() int {
	if p.Index > p.Tail {
		return 2 << (32 - 2)
	}

	return (*p.Arr)[p.Index]
}

func rankValueOfTwoSortedArray(nums1 []int, nums2 []int, rank int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	p1 := &Pointer{&nums1, 0, size1 - 1}
	p2 := &Pointer{&nums2, 0, size2 - 1}
	ret := 0.0
	if rank == 1 {
		if size1 == 0 {
			ret = float64(p2.Value())
		} else if size2 == 0 {
			ret = float64(p1.Value())
		} else if size1 != 0 && size2 != 0 {
			if p1.Value() > p2.Value() {
				ret = float64(p2.Value())
			} else {
				ret = float64(p1.Value())
			}
		}

		return ret
	}
	for rank-1 > 0 {
		rank--
		if p1.Value() > p2.Value() {
			tmp := p1
			p1 = p2
			p2 = tmp
		}
		p1.Index++
		if p1.Value() < p2.Value() {
			ret = float64(p1.Value())
		} else {
			ret = float64(p2.Value())
		}
	}

	return ret
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	rank := (size1 + size2 + 1) / 2
	isEven := ((size1+size2)%2 == 0)
	medians := make([]float64, 2)
	k := 1.0
	medians[0] = rankValueOfTwoSortedArray(nums1, nums2, rank)
	if isEven == true {
		k = 0.5
		medians[1] = rankValueOfTwoSortedArray(nums1, nums2, rank+1)
	}
	var sum float64
	for _, v := range medians {
		sum += v
	}

	return k * sum
}

func main() {
	nums1 := []int{1}
	nums2 := []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

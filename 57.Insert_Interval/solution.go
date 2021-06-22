package main

import (
	"errors"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	// 1. sort by interval[0]
	// 2. merge intervals[0:]
	// 3. shrink intervals size if necessary
	intervals = append(intervals, newInterval)
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > intervals[len(intervals)-1][0] {
			intervals[i], intervals[len(intervals)-1] = intervals[len(intervals)-1], intervals[i]
		}
	}

	for i := 0; i < len(intervals)-1; i++ {
		merged, err := mergeIntervals(intervals[i], intervals[i+1])
		if err != nil {
			continue
		}
		intervals[i] = nil
		intervals[i+1] = merged
	}

	mergedIntervals := make([][]int, 0, len(intervals))
	for i := 0; i < len(intervals); i++ {
		if intervals[i] != nil {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}
	return mergedIntervals
}

// test [a, b] cross [m, n]
//  0: a equal b
//  1: a inside b
//  2: a left cross b only
//  3: a right cross b only
//  4: a contain b
// -1: no cross
func cross(a, b []int) int {
	// a equal b
	if a[0] == b[0] && a[1] == b[1] {
		return 0
	}
	// a inside b
	if a[0] >= b[0] && a[1] <= b[1] {
		return 1
	}
	// a left cross b only
	if a[0] >= b[0] && a[0] <= b[1] {
		return 2
	}
	// a right cross b only
	if a[1] >= b[0] && a[1] <= b[1] {
		return 3
	}
	// a contain b
	if a[0] < b[0] && a[1] > b[1] {
		return 4
	}
	return -1
}

func mergeIntervals(a, b []int) ([]int, error) {
	crossed := cross(a, b)
	if crossed == -1 {
		return []int{}, errors.New("can't merge")
	}
	if crossed == 0 || crossed == 1 {
		return b, nil
	}
	if crossed == 2 {
		return []int{b[0], a[1]}, nil
	}
	if crossed == 3 {
		return []int{a[0], b[1]}, nil
	}
	if crossed == 4 {
		return a, nil
	}
	return []int{}, errors.New("missing case")
}

package main

import (
	"fmt"
	"sort"
	"strconv"
)

type T []int

func (a *T) Len() int           { return len((*a)) }
func (a *T) Swap(i, j int)      { (*a)[i], (*a)[j] = (*a)[j], (*a)[i] }
func (a *T) Less(i, j int) bool { return (*a)[i] < (*a)[j] }

func threeSum(nums T) [][]int {
	hash := make(map[string]bool)
	sort.Sort(&nums)
	n := len(nums)
	ret := make([][]int, 0, n)
	for i := 0; i < n-2; i++ {
		a := nums[i]
		start := i + 1
		end := n - 1
		for start < end {
			b := nums[start]
			c := nums[end]
			if a+b+c == 0 {
				key := strconv.Itoa(a) + "," + strconv.Itoa(b) + "," + strconv.Itoa(c)
				if hash[key] != true {
					ret = append(ret, []int{a, b, c})
					hash[key] = true
				}
				// Continue search for all triplet combinations summing to zero.
				if b == nums[start+1] {
					start = start + 1
				} else {
					end = end - 1
				}
			} else if a+b+c > 0 {
				end = end - 1
			} else {
				start = start + 1
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(threeSum(T{0, 0, 0}))
}

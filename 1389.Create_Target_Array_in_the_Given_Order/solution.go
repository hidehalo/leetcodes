package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	bucket := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		if bucket[index[i]] == nil {
			bucket[index[i]] = make([]int, 0)
		}
		bucket[index[i]] = append(bucket[index[i]], nums[i])
	}
	ret := make([]int, 0)
	uniq := make(map[int]int)
	for i := 0; i < len(index); i++ {
		if _, ok := uniq[index[i]]; ok {
			continue
		}
		uniq[index[i]] = 1
		list := bucket[index[i]]
		for j := len(list) - 1; j >= 0; j-- {
			ret = append(ret, list[j])
		}
	}
	fmt.Println(bucket)
	return ret
}

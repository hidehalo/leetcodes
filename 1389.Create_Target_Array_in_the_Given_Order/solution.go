package main

func createTargetArray(nums []int, index []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(index); i++ {
		if len(ret) > index[i] {
			n := len(ret)
			ret = append(ret, 0)
			for j := n; j > index[i]; j-- {
				ret[j] = ret[j-1]
			}
			ret[index[i]] = nums[i]
		} else {
			ret = append(ret, nums[i])
		}
	}
	return ret
}

package main

type NumArray struct {
	originNums []int
	memoOfSums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{nums, make([]int, len(nums))}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		numArray.memoOfSums[i] = sum
	}
	return numArray
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 {
		i = 0
	}
	if j >= len(this.originNums) {
		j = len(this.originNums) - 1
	}
	if j < i {
		return 0
	}
	return this.originNums[i] + this.memoOfSums[j] - this.memoOfSums[i]
}

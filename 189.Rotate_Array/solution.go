package main

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	offset1 := 0
	offset2 := len(nums) - k
	// fixme: overlap situation
	for offset1 < offset2 {
		swapArray(nums, offset1, offset2, k)
		offset1 += k
	}
}

func swapArray(arr []int, offset1, offset2, length int) {
	for i := 0; i < length; i++ {
		if offset1+i >= len(arr) || offset2+i >= len(arr) {
			break
		}
		arr[offset1+i], arr[offset2+i] = arr[offset2+i], arr[offset1+i]
	}
}

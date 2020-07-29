package main

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	reverseArray(nums[0:])
	reverseArray(nums[0:k])
	reverseArray(nums[k:])
}

func reverseArray(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

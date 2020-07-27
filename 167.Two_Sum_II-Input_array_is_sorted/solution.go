package main

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > target {
			return []int{}
		}
		j := binarySearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func binarySearch(numbers []int, start, end, target int) int {
	if start > end {
		return -1
	}
	pivot := start + (end-start)/2
	if target == numbers[pivot] {
		return pivot
	} else if target > numbers[pivot] {
		return binarySearch(numbers, pivot+1, end, target)
	}
	return binarySearch(numbers, start, pivot-1, target)
}

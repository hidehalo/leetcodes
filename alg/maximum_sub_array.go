package main

import "fmt"

func findMaximumSubArray(arr []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, arr[low]
	}
	mid := (low + high) / 2
	leftMaxStart, leftMaxEnd, leftMaxSum := findMaximumSubArray(arr, low, mid)
	rightMaxStart, rightMaxEnd, rightMaxSum := findMaximumSubArray(arr, mid+1, high)
	crossMaxStart, crossMaxEnd, crossMaxSum := findMaxCrossingSubArray(arr, low, mid, high)
	if leftMaxSum > rightMaxSum && leftMaxSum > crossMaxSum {
		return leftMaxStart, leftMaxEnd, leftMaxSum
	} else if rightMaxSum > leftMaxSum && rightMaxSum > crossMaxSum {
		return rightMaxStart, rightMaxEnd, rightMaxSum
	}

	return crossMaxStart, crossMaxEnd, crossMaxSum
}

func findMaxCrossingSubArray(arr []int, low, mid, high int) (int, int, int) {
	var maxLeft, maxRight, leftSum, rightSum, sum int
	leftSum = arr[mid]
	maxLeft = mid
	sum = leftSum
	for i := mid - 1; i >= low; i-- {
		sum += arr[i]
		if sum >= leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum = arr[mid+1]
	maxRight = mid + 1
	sum = rightSum
	for j := mid + 2; j <= high; j++ {
		sum += arr[j]
		if sum >= rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	sum = leftSum + rightSum
	if maxLeft == mid && maxRight == mid+1 {
		max := arr[low]
		for k := low; k <= high; k++ {
			if arr[k] >= max {
				max = arr[k]
				maxLeft, maxRight = k, k
			}
		}
		sum = max
	}

	return maxLeft, maxRight, sum
}

func bruteForceFindMaximumSubArray(arr []int, low, high int) (int, int, int) {
	var maxStart, maxEnd, maxSum, sum int
	for i := low; i <= high; i++ {
		sum = 0
		for j := 0; j <= high-i; j++ {
			if i == low && j == 0 {
				maxStart, maxEnd = i, i+j
				maxSum = arr[i]
				sum = arr[i]
				continue
			}
			sum += arr[i+j]
			if sum > maxSum {
				maxSum = sum
				maxStart, maxEnd = i, i+j
			}
		}
	}

	return maxStart, maxEnd, maxSum
}

func optimalFindMaximumSubArray(arr []int, low, high int) (int, int, int) {
	max := 0
	for k := low; k <= high; k++ {
		max += arr[k]
	}
	i, j := low, high
	leftSum, rightSum := max, max
	for i <= j {
		leftSum -= arr[i]
		rightSum -= arr[j]
		if max > leftSum && max > rightSum {
			break
		}
		// FIXIME: figure out how to find new max value
		if leftSum > max && leftSum > rightSum {
			max -= arr[i]
			i++
		}
		if rightSum > max && rightSum > leftSum {
			max -= arr[j]
			j--
		}
	}

	return i, j, max
}

func main() {
	arr := []int{-1, -2, -3, 4, 5, 1, 2, 3, -1, -3, -1, -6, 7, 8, 9, 10, 12, -10, -1, 9}
	fmt.Println(findMaximumSubArray(arr, 0, len(arr)-1))
	fmt.Println(bruteForceFindMaximumSubArray(arr, 0, len(arr)-1))
	fmt.Println(optimalFindMaximumSubArray(arr, 0, len(arr)-1))
}

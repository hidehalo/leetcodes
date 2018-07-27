package main

import (
	"fmt"
)

const INT32_MAX = 2147483647
const INT32_MIN = -2147483648

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		var num int
		// define has a k divide nums to 2 sorted sets Ak&Bk, guessing k is mid
		// PS: Ak := nums[0:k] and Bk := nums[k+1:], A'k := nums[0:mid] and B'k = [mid+1:]
		// if nums[mid] and target both in A'k or Bk', means the target in a sorted set
		// all right, then do binary search as normal
		// otherwise, target and nums[mid] belongs to different set
		// make nums be a sorted set (that's tricky) for binary search
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
		} else {
			// at first, judge the target in which set(A'k or B'k)
			// HOW: compare target with the smallest item of A'k (A0 => nums[0])
			// if target belongs to B'k, and A'mid > target (nums rotated)
			// think about try to convert A'k < B'k (Ans: pretend A'k := {-INF...})
			// otherwise target belongs to A'k, but idk if B'mid > target, so pretend B'k := {+INF...}
			if target < nums[0] {
				num = INT32_MIN
			} else {
				num = INT32_MAX
			}
		}
		if num < target {
			lo = mid + 1
		} else if num > target {
			hi = mid
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{0, 1, 2, 4, 5, 6, 7}
	fmt.Println(search(arr, 1))
}

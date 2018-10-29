package main

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return [][]int{}
	}
}

func twoSum(nums []int, a int, t int) bool {
	for _, v := range nums {
		if v+a == t {
			return true
		}
	}

	return false
}

func recu() {

}

func main() {

}

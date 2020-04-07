package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func replaceElements(arr []int) []int {
	c := make([]int, len(arr)+1)
	c[len(arr)] = -1 // maximum of arr[i:]
	for i := len(arr) - 1; i >= 0; i-- {
		c[i] = max(c[i+1], arr[i])
	}

	return c[1:]
}

package main
// import "fmt"
func heightChecker(heights []int) int {
	ret := 0
    for i :=1;i<len(heights);i++{
		// fmt.Printf("arr[%d]: %d, arr[%d]: %d\n", i-1, heights[i-1],i, heights[i])
		if heights[i-1]>heights[i] {
			ret++
			// find minimum & swap
			// optimization：could use priority queue
			min := heights[i]
			minIndex := i
			// fmt.Println(heights)
			for j:=i+1;j<len(heights);j++ {
				if heights[j] < min{
					minIndex = j
					min = heights[j]
				}
			}
			// fmt.Printf("Swap arr[%d]:%d, arr[%d]:%d\n", i-1,heights[i-1],minIndex,heights[minIndex])
			heights[i-1],heights[minIndex] = heights[minIndex],heights[i-1]
		}
	}
	return ret
}

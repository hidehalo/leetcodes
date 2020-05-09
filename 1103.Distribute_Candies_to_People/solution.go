package main

func distributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	giveNum := 1
	nIdx := 0
	for candies > 0 {
		if giveNum > candies {
			ret[nIdx] += candies
			break
		}
		ret[nIdx] += giveNum
		candies -= giveNum
		giveNum++
		nIdx = (nIdx + 1) % num_people
	}

	return ret
}

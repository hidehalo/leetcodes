package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}
	max := candies[0]
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	ret := make([]bool, 0)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies < max {
			ret = append(ret, false)
		} else {
			ret = append(ret, true)
		}
	}

	return ret
}

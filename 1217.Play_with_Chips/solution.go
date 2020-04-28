package main

func minCostToMoveChips(chips []int) int {
	even := 0
	odd := 0
	// even chip convert to any even chip is zero cost
	// odd chip convert to any odd chip is zero cost
	// any even chip convert to any odd chip only cost 1 unit
	// so, the min cost to move chips is min(numOfEvens, numOfOdds)
	for i := 0; i < len(chips); i++ {
		if chips[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even > odd {
		return odd
	}
	return even
}

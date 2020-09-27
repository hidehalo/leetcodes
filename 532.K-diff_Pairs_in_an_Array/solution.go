package main

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	distincKeyMap := make(map[int]int)
	usedKeyMap := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		distincKeyMap[nums[i]]++
	}
	ret := 0
	for key := range distincKeyMap {
		// fmt.Println("Test key=", key, "diff=", k)
		if k == 0 {
			// fmt.Println("Find target=", key, "(", distincKeyMap[key], ")")
			if distincKeyMap[key] >= 2 {
				ret++
			}
			continue
		}
		targets := []int{key - k, key + k}
		for _, target := range targets {
			if _, ok := distincKeyMap[target]; ok == true {
				duplicated := false
				if _, ok := usedKeyMap[target]; ok == true {
					for j := 0; j < len(usedKeyMap[target]); j++ {
						if usedKeyMap[target][j] == key {
							// fmt.Println("Pair(", key, ",", key-k, ")", "already counted")
							duplicated = true
							break
						}
					}
				}
				if !duplicated {
					// fmt.Println("Find target=", target, "(", distincKeyMap[target], ")")
					usedKeyMap[key] = append(usedKeyMap[key], target)
					usedKeyMap[target] = append(usedKeyMap[target], key)
					ret++
				}
			}
		}
	}
	return ret
}

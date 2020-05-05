package main

func distributeCandies(candies []int) int {
	// greedy algo
	// make a count for every kind of candies
	// sis take one from every kind
	// test sis's candies number less equal bro's
	// yes: return number of kind
	// no: return n(sis) - (n(sis) - n(bro) / 2)
	counts := make(map[int]int) // hash table is too slow on big case, intger array is better
	for i := 0; i < len(candies); i++ {
		counts[candies[i]]++
	}
	countOfSis := len(counts)
	countOfBro := len(candies) - len(counts)
	if countOfSis <= countOfBro {
		return countOfSis
	}
	return countOfSis - (countOfSis-countOfBro)/2
}

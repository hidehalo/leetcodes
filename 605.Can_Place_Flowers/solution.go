package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	} else if n > len(flowerbed) {
		return false
	} else if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 0 {
			if i > 0 && i < len(flowerbed)-1 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			} else if i == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			} else if i == len(flowerbed)-1 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}

package main

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	if len(heaters) == 0 {
		return 2147483647
	} else if len(houses) == 0 {
		return 0
	}
	sort.Ints(houses)  // O(n) = lgn
	sort.Ints(heaters) // O(n) = lgm
	min := houses[0]
	max := houses[len(houses)-1]
	atLeast := -1
	outsideRadius := 2147483647

	// outside_left [border_left...indise...border_right] outside_right
	// d = max(max(inside-border_left, border_right-inside), min(border_left - min(outside_left)), min(outside_right-border_right))
	for i := 0; i < len(heaters); i++ { // O(n) = mlgn
		if min > heaters[i] && max-heaters[i] < outsideRadius { // outside left
			outsideRadius = max - heaters[i]
		} else if max < heaters[i] && heaters[i]-min < outsideRadius { // outeside right
			outsideRadius = heaters[i] - min
		} else if heaters[i] >= min && heaters[i] <= max { // inside
			s := succesor(houses, heaters[i])
			var p int
			if s > 0 {
				p = precessor(houses[:s], heaters[i])
			} else {
				p = precessor(houses, heaters[i])
			}
			if s < 0 && p < 0 {
				if max-heaters[i] > atLeast {
					atLeast = max - heaters[i]
				}
				if heaters[i]-min > atLeast {
					atLeast = heaters[i] - min
				}
			} else {
				if s > 0 && houses[s]-heaters[i] > atLeast {
					atLeast = houses[s] - heaters[i]
				}
				if p >= 0 && heaters[i]-houses[p] > atLeast {
					atLeast = heaters[i] - houses[p]
				}
			}
		}
	}
	if atLeast >= 0 {
		return atLeast
	}
	return outsideRadius
}

func succesor(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] > key {
		return succesor(arr[:mid], key)
	} else if arr[mid] < key {
		return succesor(arr[mid+1:], key)
	}
	if mid+1 >= len(arr) {
		return -1
	}
	return mid + 1
}

func precessor(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] > key {
		return precessor(arr[:mid], key)
	} else if arr[mid] < key {
		return precessor(arr[mid+1:], key)
	}
	if mid-1 < 0 {
		return -1
	}
	return mid - 1
}

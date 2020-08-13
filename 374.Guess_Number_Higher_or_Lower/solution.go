package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	s, e := 1, n
	pivot := s + (e-s)/2

	for e >= s {
		guessResult := guess(pivot)
		if guessResult == 0 {
			return pivot
		} else if guessResult == -1 {
			e = pivot - 1
		} else {
			s = pivot + 1
		}
		pivot = s + (e-s)/2
	}
	return -1
}

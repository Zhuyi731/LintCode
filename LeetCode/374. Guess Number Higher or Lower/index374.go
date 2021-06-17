package main

func main() {
	guessNumber(2)
	guessNumber(1)
	guessNumber(10)
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(n int) int {
	if n == 2 {
		return 0
	}
	return 1
}

func guessNumber(n int) int {
	l := 1
	r := n
	m := (l + r) / 2
	for l < r {
		g := guess(m)
		if g == 0 {
			return m
		} else if g == -1 {
			r = m
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}
	return m
}

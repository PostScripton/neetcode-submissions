func hammingWeight(n int) int {
	var result int

	for n != 0 {
		if n % 2 == 0 {
			n /= 2
			continue
		}

		n -= 1
		n /= 2
		result++
	}

	return result
}

// Time:  O(n)
// Space: O(1)
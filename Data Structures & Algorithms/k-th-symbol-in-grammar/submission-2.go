func kthGrammar(n int, k int) int {
	var curr int
	left, right := 1, 1 << (n - 1) // 2^(n-1)

	for range n - 1 {
		var mid int = (left + right) / 2

		if k <= mid {
			right = mid
		} else {
			left = mid + 1

			if curr == 0 {
				curr = 1
			} else {
				curr = 0
			}
		}
	}

	return curr
}

// Time:  O(n)
// Space: O(1)
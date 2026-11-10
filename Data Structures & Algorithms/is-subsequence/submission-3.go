func isSubsequence(s string, t string) bool {
	i := 0

	for j := 0; i < len(s) && j < len(t); j++ {
		if s[i] == t[j] {
			i++
			continue
		}
	}

	return i == len(s)
}

// Time:  O(n)
// Space: O(1)
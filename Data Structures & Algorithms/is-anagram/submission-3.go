func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var zero, anagram [26]int

	for i := range s {
		anagram[s[i]-'a']++
		anagram[t[i]-'a']--
	}

	return zero == anagram
}

// Time:  O(n)
// Space: O(n)

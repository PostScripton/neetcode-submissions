func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var strCharFreqs [26]int

		for _, char := range str {
			strCharFreqs[char-'a']++
		}

		groups[strCharFreqs] = append(groups[strCharFreqs], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

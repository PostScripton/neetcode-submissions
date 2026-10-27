func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqs {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets)-1; i >= 0; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

// Time:  O(n)
// Space: O(n)
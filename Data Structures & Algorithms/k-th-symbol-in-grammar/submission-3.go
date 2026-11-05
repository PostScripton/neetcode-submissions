func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}


	parent := kthGrammar(n-1, (k+1)/2)
	if k%2 == 1 {
		return parent
	}

	return 1 - parent
}

// Explaination:
// (4, 5)
// -> (3, 3), because (5 + 1) / 2 = 3
// -> (2, 2), because (3 + 1) / 2 = 2
// -> (1, 1), because (2 + 1) / 2 = 1
// 
// Now return upward:
// k = 2: even, flip 0 to 1
// k = 3: odd, keep 1
// k = 5: odd, keep 1
//
// Therefore, answer is 1

// Time:  O(n)
// Space: O(n) // recursion stack
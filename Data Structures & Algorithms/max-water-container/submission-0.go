func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	max := math.MinInt64

	for l < r {
		width := r - l
		height := min(heights[l], heights[r])

		water := width * height

		if water > max {
			max = water
		}

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

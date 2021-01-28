package longestsubstringwithoutrepeatingcharacters_3

func lengthOfLongestSubstring(s string) int {
	var max = 0

	for i := 0; i < len(s); i++ {
		m := make(map[byte]bool)
		step := 0
		for idx := i; idx < len(s); idx++ {
			if !m[s[idx]] {
				step++
				m[s[idx]] = true
			} else {
				break
			}
		}

		if max < step {
			max = step
		}
	}

	return max
}

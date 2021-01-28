package longestpalindromicsubstring_5

func LongestPalindrome(s string) string {
	length := len(s)
	m := make([][]bool, length)
	for i := 0; i < length; i++ {
		m[i] = make([]bool, length)
		m[i][i] = true
	}

	var start, end, max = 0, 0, 0

	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] && (i - j < 2 || m[j+1][i-1]) {
				delta := i - j + 1
				m[j][i] = true
				if max < delta {
					max = delta
					start = j
					end = i
				}
			}
		}
	}

	return s[start:end+1]
}

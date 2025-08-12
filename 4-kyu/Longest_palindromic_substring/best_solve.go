package longestpalindromicsubstring

func subStrHelper(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i, j = i-1, j+1
	}

	return s[i+1 : j]
}

func getMaxLength(strList ...string) string {
	var result string
	for _, str := range strList {
		if len(str) > len(result) {
			result = str
		}
	}

	return result
}

func LongestPalindrome_s(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		result = getMaxLength(result, subStrHelper(s, i, i), subStrHelper(s, i, i+1))
	}

	return result
}

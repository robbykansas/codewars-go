package longestpalindromicsubstring

import (
	"fmt"
	"strings"
)

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// Transform string to handle even/odd palindromes
	t := "#" + strings.Join(strings.Split(s, ""), "#") + "#"
	fmt.Println(t)
	n := len(t)

	// Array to store palindrome radii
	p := make([]int, n)
	center, right := 0, 0
	maxLen, maxCenter := 0, 0

	for i := 0; i < n; i++ {
		mirror := 2*center - i
		// Check if current index is within right boundary
		if i < right {
			p[i] = min(right-i, p[mirror])
		}

		// Expand around current center
		a, b := i+1+p[i], i-1-p[i]
		for a < n && b >= 0 && t[a] == t[b] {
			p[i]++
			a++
			b--
		}

		// Update center and right boundary
		if i+p[i] > right {
			center = i
			right = i + p[i]
		}

		// Track max palindrome
		if p[i] > maxLen {
			maxLen = p[i]
			maxCenter = i
		}
	}

	// Extract palindrome substring
	start := (maxCenter - maxLen) / 2
	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

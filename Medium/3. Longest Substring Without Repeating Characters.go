package main

import "fmt"

/*
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	lastIndex := make([]int, 128)

	start := 0
	for end := 0; end < n; end++ {
		currentChar := s[end]
		if lastIndex[currentChar] > start {
			start = lastIndex[currentChar]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIndex[currentChar] = end + 1
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	length := lengthOfLongestSubstring(s)

	fmt.Println(length)
}

package main

import "fmt"

/*
You are given an array of strings words. Each element of words consists of two lowercase English letters.
Create the longest possible palindrome by selecting some elements from words and concatenating them in any order. Each element can be selected at most once.
Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome, return 0.
A palindrome is a string that reads the same forward and backward.


Input: words = ["lc","cl","gg"]
Output: 6
Explanation: One longest palindrome is "lc" + "gg" + "cl" = "lcggcl", of length 6.
Note that "clgglc" is another longest palindrome that can be created.
*/

func longestPalindrome(words []string) int {
	count, same := 0, 0
	wordMap := make(map[string]int)

	for i, word := range words {
		if word[0] == word[1] {
			same++
		}
		r := []rune(words[i])
		r[0], r[1] = r[1], r[0]
		if wordMap[string(r)] > 0 {
			count += 4
			wordMap[string(r)]--
			if r[0] == r[1] {
				same -= 2
			}
		} else {
			wordMap[word]++
		}
	}

	if same > 0 {
		count += 2
	}
	return count
}

func main() {
	words := []string{"dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"}
	length := longestPalindrome(words)
	fmt.Println(length)

}

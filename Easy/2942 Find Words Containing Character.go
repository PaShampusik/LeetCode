package main

/*
You are given a 0-indexed array of strings words and a character x.
Return an array of indices representing the words that contain the character x.
Note that the returned array may be in any order.
*/

func findWordsContaining(words []string, x byte) []int {
	result := make([]int, 0, len(words))
	for i, val := range words {
		for k := 0; k < len(val); k++ {
			if val[k] == x {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

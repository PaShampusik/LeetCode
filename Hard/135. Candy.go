package main

/*
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.
You are giving candies to these children subjected to the following requirements:
Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
Return the minimum number of candies you need to have to distribute the candies to the children.

Example 1:
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:
Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.
*/

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := range n {
		if ratings[i-1] < ratings[i] {
			c[i] = c[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && c[i] < c[i+1]+1 {
			c[i] = c[i+1] + 1
		}
	}
	sum := 0
	for _, v := range c {
		sum += v
	}
	return sum
}

package main

/*
Given a circular array nums, find the maximum absolute difference between adjacent elements.
Note: In a circular array, the first and last elements are adjacent.

Example 1:
Input: nums = [1,2,4]
Output: 3
Explanation:
Because nums is circular, nums[0] and nums[2] are adjacent. They have the maximum absolute difference of |4 - 1| = 3.

Example 2:
Input: nums = [-5,-10,-5]
Output: 5
Explanation:
The adjacent elements nums[0] and nums[1] have the maximum absolute difference of |-5 - (-10)| = 5.
*/
func maxAdjacentDistance(nums []int) int {
	maxDistance := 0
	for i := 0; i < len(nums)-1; i++ {
		maxDistance = max(maxDistance, absDiff(nums[i], nums[i+1]))
	}
	maxDistance = max(maxDistance, absDiff(nums[len(nums)-1], nums[0]))
	return maxDistance
}

func absDiff(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

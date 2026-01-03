package main

import "fmt"

/*
You are given an integer array nums of length n.

A partition is defined as an index i where 0 <= i < n - 1, splitting the array into two non-empty subarrays such that:

Left subarray contains indices [0, i].
Right subarray contains indices [i + 1, n - 1].
Return the number of partitions where the difference between the sum of the left and right subarrays is even.

Example 1:

Input: nums = [10,10,3,7,6]

Output: 4

Explanation:

The 4 partitions are:

[10], [10, 3, 7, 6] with a sum difference of 10 - 26 = -16, which is even.
[10, 10], [3, 7, 6] with a sum difference of 20 - 16 = 4, which is even.
[10, 10, 3], [7, 6] with a sum difference of 23 - 13 = 10, which is even.
[10, 10, 3, 7], [6] with a sum difference of 30 - 6 = 24, which is even.
Example 2:

Input: nums = [1,2,2]

Output: 0

Explanation:

No partition results in an even sum difference.

Example 3:

Input: nums = [2,4,6,8]

Output: 3

Explanation:

All partitions result in an even sum difference.

Constraints:

2 <= n == nums.length <= 100
1 <= nums[i] <= 100
*/
func sumArr(arr []int) int {
	var result int

	for _, v := range arr {
		result += v
	}

	return result
}

func countPartitions(nums []int) int {
	n := len(nums)
	var result int
	for i := 0; i < n-1; i++ {
		left := sumArr(nums[0 : i+1])
		right := sumArr(nums[i+1:])

		if (left-right)%2 == 0 {
			result++
		}

	}
	return result
}

func main() {
	test1 := []int{10, 10, 3, 7, 6}
	test2 := []int{1, 2, 2}
	test3 := []int{2, 4, 6, 8}
	fmt.Println(countPartitions(test1))
	fmt.Println(countPartitions(test2))
	fmt.Println(countPartitions(test3))
}

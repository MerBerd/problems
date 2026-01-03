/*
The beauty of a string is the difference in frequencies between the most frequent and least frequent characters.


For example, the beauty of "abaacc" is 3 - 1 = 2.
Given a string s, return the sum of beauty of all of its substrings.



Example 1:

Input: s = "aabcb"
Output: 5
Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.
Example 2:

Input: s = "aabcbaa"
Output: 17


Constraints:

1 <= s.length <= 500
s consists of only lowercase English letters.

*/

package main

import "fmt"

func beautySum(s string) int {
	var result int

	n := len(s)

	for i := 0; i < n; i++ {
		freq := make([]int, 26)
		for j := i; j < n; j++ {
			freq[s[j]-'a']++

			minVal, maxVal := n+1, freq[0]

			for k := 0; k < 26; k++ {
				if freq[k] == 0 {
					continue
				}
				if freq[k] < minVal {
					minVal = freq[k]
				}
				if freq[k] > maxVal {
					maxVal = freq[k]
				}
			}

			if minVal == 0 {
				minVal = maxVal
			}

			result += (maxVal - minVal)

		}

	}
	return result
}

func main() {
	fmt.Println(beautySum("aab"))
}

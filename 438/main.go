package main

import "fmt"

func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	n := len(s)
	m := len(p)

	freqP := make([]int, 26)

	for i := 0; i < m; i++ {
		freqP[p[i]-'a']++
	}

	for i := 0; i < n; i++ {
		freqS := make([]int, 26)
		for j := i; j < n; j++ {
			freqS[s[j]-'a']++

			fmt.Println(freqS)
			fmt.Println(freqP)
			match := true
			for k := 0; k < 26; k++ {
				if freqS[k] != freqP[k] {
					match = false
					break
				}
			}

			if match {
				result = append(result, i)
			}
		}

	}
	fmt.Println(result)
	return result
}

func main() {
	//findAnagrams("abab", "ab")
	//findA("abab", "ab")
	//findAnagrams("aa", "bb")
	findAnagrams("aa", "bb")
}

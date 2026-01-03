package main

import "fmt"

func checkOnesSegment(s string) bool {


	var segment, segment_count int
		if len(s) == 1 {
		return true
	}

	
	for _, v := range s {

		if v == '1' {
			segment += 1
		} else {
			if segment != 0 {
				segment_count += 1
			}
			segment = 0
		}

	}

	if segment >= 1 {
		segment_count += 1
	}

	if segment_count == 1 {
		return true
	} else {
		return false
	}

}


func main() {


	fmt.Println(checkOnesSegment("111000"))
	fmt.Println(checkOnesSegment("10"))
	fmt.Println(checkOnesSegment("1"))
	fmt.Println(checkOnesSegment("101"))

	fmt.Println(checkOnesSegment("10101"))
	fmt.Println(checkOnesSegment("1001"))
	fmt.Println(checkOnesSegment("10101010"))
	fmt.Println(checkOnesSegment("10111"))



}
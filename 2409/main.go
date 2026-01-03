package main

import (
	"fmt"
	"strconv"
)

func DateToDay(date string) int {
	days := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	month := date[:2]
	day := date[3:]

	month_i, err1 := strconv.Atoi(month)
	day_i, err2 := strconv.Atoi(day)

	if err1 != nil || err2 != nil {
		return 0
	}


	for i:=0; i < month_i-1; i++ {
		day_i += days[i]
	}

	return day_i
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {

	x1_1 := DateToDay(arriveAlice)
	x1_2 := DateToDay(leaveAlice)

	x2_1 := DateToDay(arriveBob)
	x2_2 := DateToDay(leaveBob)

	var x, y int

	if x2_1 >= x1_1 {
		x = x2_1
	} else {
		x = x1_1
	}

	if x1_2 <= x2_2 {
		y = x1_2
	} else {
		y = x2_2
	}
	 
	res := (y - x) + 1
	if res < 0  || res > 365{
        return 0
    } else {
        return res
    }
    
}

func main() { 
	str1 := "10-01"
	str2 := "10-31"
	str3 := "11-01"
	str4 := "12-31"

	fmt.Println(countDaysTogether(str1, str2, str3, str4))


	
}
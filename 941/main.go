package main

import "fmt"


func validMountainArray(arr []int) bool {
    mount := true
   
	var peak int

    n := len(arr)

    if n < 3 {
        return false
    }

    for i:=0; i<n-1; i++ {
        if arr[i] >= arr[i+1]{
			peak = i
			fmt.Printf("The peak is: %v -> %v\n", peak, arr[peak])
			break
        } 
    }

	if peak+1 == n-1 {
		if arr[peak+1] == arr[peak] {
			return false
		} else {
			return true
		}
	} else if peak == 0 {
		return false
	}

	for i:=peak+1; i<n-1; i++ {
		if arr[i] <= arr[i+1]{
            mount = false
			break
        }
	}
    return mount
}


func main() {
	test4 := []int{9,8,7,6,5,4,3,2,1,0}

	fmt.Println(validMountainArray(test4))

}
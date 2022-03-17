package main

import (
	"fmt"
)


/*
  first []dTypes, second: len, third: capacity
  a := make([]int, 5)  // len(a)=5
  
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5

*/
type data_arr []int

// 1,2,3

func smallerNumbersThanCurrent(nums []int) []int {
	var ans = make([]int, len(nums))
	var countarray = make([]int, 10)
	var count int

	for i := 0; i < len(nums); i++ {
		countarray[nums[i]]++
	}
	fmt.Println("countarray: ", countarray)

	for i := 0; i < len(nums); i++ {
		count = 0
		for j := 0; j < nums[i]; j++ {
			count += countarray[j]
			fmt.Println("j: ", j, countarray[j])
		}
		fmt.Println("broder -----")
		ans[i] = count
	}
	return ans
}

func main() {

	values := data_arr{4, 3, 2, 6, 4}
	result := smallerNumbersThanCurrent(values)
	fmt.Println(result)

}

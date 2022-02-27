package main

import (
	"fmt"
)

// go function take another function as parameter

/*
should understand this well
bcs further we'll encounter this form of function oftenly
*/
type func_param func(int) int
type arr_param []int

func apply(nums arr_param, f func_param) {
	for i, v := range nums {
		nums[i] = f(v)
	}
	fmt.Println(nums)

}
func main() {
	nums := []int{10, 20}
	apply(nums, func(i int) int { return i * 2 })

}

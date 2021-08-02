package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("addition of two slice")
	//two numbers 
	num1 := []int{1, 2, 3, 4, 5, 6}
	num2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(Add(num1, num2))

}
func Add(num1 []int, num2 []int) []int {
	c := make([]int, len(num1))
//Sorted slice and revese the second number  
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for i, _ := range num1 {
		
		c[i] = num1[i] + num2[i]
	}
	return c
}

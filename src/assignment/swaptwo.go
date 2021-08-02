package main

import "fmt"

func main() {
	x := 5
	y := 6

	fmt.Printf("Before Swaping value of x is %d and y is %d  ", x, y)
	swap(&x, &y)
	fmt.Printf("after swapping value of x is %d and y is %d ", x, y)

}
func swap(a, b *int) {
	tempx := *a
	tempy := *b
	*a = tempy
	*b = tempx

}

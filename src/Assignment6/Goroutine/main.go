package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrintEvenNumbers() {
	defer wg.Done()
	fmt.Println("Even Numbers :(")
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func PrintOddNumbers() {
	defer wg.Done()
	fmt.Println("Odd Numbers :")
	for i := 1; i < 20; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Get The main funtion..")
	fmt.Println("Hello Golang")
	wg.Add(3)
	go func() {
		fmt.Println("Annonymous Function")
		wg.Done()
	}()
	go PrintEvenNumbers()
	go PrintOddNumbers()
	fmt.Println("get out main Function")
	wg.Wait()

}

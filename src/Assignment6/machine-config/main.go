package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("logical processor :", runtime.NumCPU())
	fmt.Println("operating  system:", runtime.GOOS)
	fmt.Println("Sysmet Architecture :", runtime.GOARCH)
	fmt.Println("max processor:", runtime.GOMAXPROCS(14))
}

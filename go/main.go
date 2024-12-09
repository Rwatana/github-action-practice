package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Hello, World!")
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Go version: ", runtime.Version())
}

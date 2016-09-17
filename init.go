package main 

import (
    "fmt"
)

var foo int = 32

func startClient() {
	foo++
	fmt.Println(foo)
	fmt.Printf("This is startClient \n")
}

func startCache () {
	foo += 2
	fmt.Println(foo)
}
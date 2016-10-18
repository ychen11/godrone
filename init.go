package main 

import (
    "fmt"
)

var foo int = 32
var cache map[string]string

func startClient() {
	foo++
	fmt.Println(foo)
	fmt.Printf("This is startClient \n")
}

func startCache () {
	cache = make(map[string]string)
	fmt.Println(foo)
}

func putData (key string, value string) {
	cache[key] = value
	fmt.Printf(cache[key])
}
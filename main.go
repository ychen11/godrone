package main 

import (
         "fmt"
)


func start () {
	startClient()
}


func startLocalCache () {
	startCache()
}

func main() {
	start()
	startLocalCache()
	fmt.Printf("This is a main file. \n")
}
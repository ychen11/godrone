package main 

import (
    "fmt"
)

var foo int = 32
var cache map[string]string


/**
 * Init a pre-allocated map.
 */
func initCache () {
  cache = make(map[string]string)
  fmt.Printf("Cache init'd. \n")
}

/**
 * Returns true if this operation succeed.
 */
func cachePutOne (key string, value string) (bool){
  cache[key] = value
  return true
}

/**
 * Returns true 
 */
func cacheGetOne (key string) (string, bool) {
  value := cache[key]
  if value == "" {
    return "", false
  } else {
    return value, true
  } 
}
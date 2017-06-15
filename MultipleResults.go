package main

import "fmt"

func fSwap(iString1, iString2 string) (string, string) {
  return iString2, iString1
}

func main() {
  aWord1, aWord2 := fSwap("hello","world")
  fmt.Println(aWord1, aWord2)
}

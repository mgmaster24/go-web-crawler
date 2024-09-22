package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	val, _ := NormalizeURL("https://www.boot.dev/lessons/98ac1f38-22dd-4682-b114-8638a0625567")
	fmt.Println(val)

}

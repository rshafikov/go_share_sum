package main

import "fmt"

func Add(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {
	fmt.Println(Add(2, 2))
}

package main

import "fmt"

func main() {
	println(Insert[int](1, 12, []int{1, 2}))
	fmt.Printf("%v", Insert[int](0, 12, []int{1, 2}))
}

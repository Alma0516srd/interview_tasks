package main

import "fmt"

func insert(a []int) {
	a = append(a, 4)
}

func main() {
	a := make([]int, 0, 4)
	a = append(a, []int{1, 2, 3}...)
	insert(a)
	fmt.Println(a)
}

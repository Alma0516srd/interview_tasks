package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b = append(b, 4)
	c := b
	b[0] = 0
	e := append(c, 5)
	b[2] = 7

	fmt.Println(a, b, c, e) // что отобразится после вызова?
}

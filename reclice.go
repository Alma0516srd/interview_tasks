package main

import "fmt"

func testSlices1() {
	a := []string{"a", "b", "c"}
	b := a[1:2]
	b[0] = "q"

	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

func testSlices2() {
	a := []byte{'a', 'b', 'c'}
	b := append(a[1:2], 'd')
	b[0] = 'z'

	fmt.Printf("%s\n", a) // что отобразится после вызова?
}

func main() {
	testSlices2()
}

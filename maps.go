package main

import "fmt"

func main() {
	var m map[string]int

	fmt.Println(m["foo"])

	m["foo"] = 42 // panic - map не инициализирован

	fmt.Println(m["foo"])
}

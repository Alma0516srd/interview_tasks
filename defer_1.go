package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	panic("error") // выполняется после всех defer-ов
}

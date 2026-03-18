package main

import "fmt"

func main() {
	v := 5
	p := &v
	fmt.Println(*p) // что отобразится после вызова?

	changePointer(p)
	fmt.Println(*p) // что отобразится после вызова?
}

func changePointer(p *int) { // лок копия указателя
	v := 3
	p = &v // передается копия указателя
}

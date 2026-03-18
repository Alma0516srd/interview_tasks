package main

import "fmt"

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{
		Name: "Olga",
	} // сюда передается копия указателя. После выхода из функции копия уничтожается
}

func main() {
	person := &Person{
		Name: "Eugene",
	}
	fmt.Println(person.Name) // что отобразится после вызова?
	changeName(person)
	fmt.Println(person.Name) // что отобразится после вызова?
}

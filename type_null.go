package main

import "fmt"

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1) // что отобразится после вызова?

	var e *errorString
	checkErr(e) // что отобразится после вызова?

	e = &errorString{}
	checkErr(e) // что отобразится после вызова?

	e = nil
	checkErr(e) // что отобразится после вызова?
}

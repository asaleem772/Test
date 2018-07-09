package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ali := person{"Ali", "Saleem", contactInfo{"ali.saleem@gmail.com", 12345}}
	aliPtr := &ali
	aliPtr.updateFirstName("Aly")
	ali.print()
}
func (ptrToPerson *person) updateFirstName(firstName string) {
	(ptrToPerson).firstName = firstName
	fmt.Println(*ptrToPerson)
	fmt.Println("****")
	fmt.Println(ptrToPerson)
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

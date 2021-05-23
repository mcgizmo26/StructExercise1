package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Anderson",
	// }

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)

	// The %+v allows us to print out the structs properties and
	// values
	fmt.Printf("%+v", alex)

}

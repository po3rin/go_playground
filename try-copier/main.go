package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Role string
	Age  int32
}

type Employee struct {
	Name      string
	Role      string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}


func main() {
	var (
		user      = User{}
		employee  = Employee{Name: "Jinzhu",
		Role: "Admin", 
		Age: 18, 
		EmployeId: 23,
		SuperRule: "admin"}
	)

	copier.Copy(&user, &employee)
	fmt.Printf("%#v \n", user)
}
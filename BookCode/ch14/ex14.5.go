package main

import "fmt"

type User struct {
	Name string
	Age int
}

func NewUser(name string, age int) *User{
//var u = User(name, age)
//return &u


	return &User{name, age}

//var a = User(name, age)
//return &a
}

func main() {

	userPointer := NewUser("AAA",23)
	fmt.Println(userPointer)

}
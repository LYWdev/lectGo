package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age	 int
}

func ( s Student )String() string {
	return fmt.Sprintf("Hello, %d, %s", s.Age, s.Name)

}

func main() {

	Student := Student{" 철수",12}
	var stringer Stringer

	stringer =Student
	fmt.Printf("%s", stringer.String())
}
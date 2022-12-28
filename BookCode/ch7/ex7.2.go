package main

import "fmt"

func main() {
	math 	:=80
	eng 	:=74
	history	:=95 

	fmt.Println("age :",(math+eng+history)/3)

	math 	=20
	eng 	=74
	history	=95 

	fmt.Println("age :",(math+eng+history)/3)

	math 	=80
	eng 	=74
	history =95 

	fmt.Println("age :",(math+eng+history)/3)
}


package main

import "fmt"

type opFunc func(a,b int) int

func getOperator(op string) opFunc{

	switch op {
		case "+":

			return func(a,b int)int {
			//Go에서 함수 리터럴(익명함수, 람다 등)
				return a+b
			}
		case "*":
			return func(a,b int)int {
				return a*b
			}
			default :
				return nil
	} 
}

func main() {
	fn := getOperator("*")

	var result = fn(3,4)
	fmt.Println(result)
}

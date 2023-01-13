package main

import "fmt"

type account struct {
	balance  int
}


func withdrawFunc( a *account, amount int){ 		//일반  함수
	a.balance -=amount
	fmt.Println("Func",a.balance)
}

func (a *account)withdrawMethod(amount int){ 		//일반  함수
	a.balance -=amount
	fmt.Println("Method",a.balance)
}

func main(){
	a:=&account{100}

	withdrawFunc(a,30)
	a.withdrawMethod(20)

	fmt.Printf("%d\n",a.balance)
}

package main

import "fmt"


type account struct{
	balance 		int
	firstName		 string
	lastName	 		string
} 		// 사용자 지정 별칭 타입

func (num1 *account)withdrawPointer(amount int){ 		//일반  함수
	num1.balance -= amount
}

func (num2 account)withdrawValue(amount int){ 		//일반  함수
	num2.balance -= amount
}
func (num2 account)withdrawValue2(amount int)account{ 		//일반  함수
	num2.balance -= amount
	return num2
}

func (num3 account)withdrawReturnValue(amount int) account{ 		//일반  함수
	num3.balance -= amount
	return num3
}

func main(){
	var mainA *account = &account{100, "Joe","Park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(2)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	mainA.withdrawValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
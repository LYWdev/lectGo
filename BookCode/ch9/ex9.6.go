package main

import "fmt"

func HasFriendRich() bool {
	return false
}

func GetFrendsCount() int {
	return 3
}
func Init() int {
	return 50000
}
func main() {
	

	if price :=Init() ; price >= 50000{
		if HasFriendRich(){
			fmt.Println("50000보다  HFR이 참인 값")
		}else {
			fmt.Println("50000보다 작고 HFR이 거짓 값")
		}
	}else if price >=30000 {
		if GetFrendsCount()>3{
			fmt.Println("30000보다 크고  HFR이 참인 값")
		}else{
			fmt.Println("30000보다 크고  HFR이 거짓 값")
		}
	}
}
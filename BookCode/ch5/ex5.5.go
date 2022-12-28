package main

import "fmt"

func main() {
	var a int
	var b int

	//var 와 int타입 생략
	//이 부분이 이해가지 않을 수 도 있다.
	//func Scan(a ...interface{})(n int err error)
	//기본적으로 Scan 함수의 리턴 ㄱ밧이 성공적으로 입력한 값 개수와 입력 실패 시 에러를 반환한다.
	s,e := fmt.Scan(&a, &b)

	if e!= nil{
		// 에러값이 빈값이 아니면 == 에러가 발생하면
		fmt.Println(s,e)
	}else{
		// 에러값이 빈 값이면 = 에러가 발생하지 않으면 
		fmt.Println(s,a,b)
	}
}
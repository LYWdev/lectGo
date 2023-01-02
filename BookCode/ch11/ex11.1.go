package main

import "fmt"

func main() {
	// 초기문; 조건문; 후처리;로 구성된다.
	for i := 0; i<10; i++ {
		fmt.Print(i,"")
	}
	//이미 이곳에서는 i 변수가 해제되었기 때문에 사용할수 없음.
	//go언어는 {}를 기준으로 변수의 범위를 관리한다.
}
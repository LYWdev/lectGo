package main

import (
	"fmt"
)
//레이블을 이용한 반복문 종료
func main() {
	a:=1
	b:=1
//레이블 정의
OuterFor:
	for ; a<=9; a++{
		for b=1; b<=9; b++{
			if a*b==45{
				break OuterFor
			}
		}
	} 
	fmt.Printf("%d * %d =%d\n",a,b,a*b)
}
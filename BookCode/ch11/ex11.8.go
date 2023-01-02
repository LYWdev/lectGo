package main

import (
	"fmt"
)
// flag(found)를 활용한 분기점 특정하기
func main() {
	a := 1
	b := 1
	OuterFor:
	for ; a <=9; a++{
		for b=1; b <=9; b++{
			//안쪽 루프에서 찾았나 검사 
			if a*b == 45{
				break OuterFor
			}
		}
	}
	fmt.Printf("%d=%d=%d\n", a, b, a*b)
}
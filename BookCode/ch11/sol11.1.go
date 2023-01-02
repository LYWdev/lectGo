package main

import (
	"fmt"
)

//되도록 플래그를 사용하고 레이브은 사용하지 않아야 클린 코드를 지향할 수 있다.
//중첩과 플래그, 레이블이 없는 깔끔한 코드로 수정하면 아래와 같다.

func find45(a int) (int,bool){//45가 되는 값을 찾는 함수
	for b:=1; b<=9; b++{
		if a*b == 45{
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 0
	for ; a <=9; a++{
		var found bool
		if b, found = find45(a);found{
			break
		}
	}
	fmt.Printf("%d*%d=%d\n", a, b, a*b)
}
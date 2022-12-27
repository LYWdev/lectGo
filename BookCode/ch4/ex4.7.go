// 4.7 정수 표현

//2바이트 부호 있는 정수에서 양수 음수를 구분하기위한 1비트를 제외하면 15비트를 이용해서 숫자를 표현한다.

package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a*b 
	var d float32 = c*3 

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
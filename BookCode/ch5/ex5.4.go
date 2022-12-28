package main

import "fmt"

func main() {
	str := "Hello\tGo\ttworld\n\"Go\"is Awesome!\n"

	fmt.Print(str)// 문자열을 기본 수식으로 출력
	fmt.Printf("%s",str)//  %s 수식으로 출력
	fmt.Printf("%q",str)// %q 모든 특수 문자가 기능을 잃고 문자 자체로 동작한다.

}

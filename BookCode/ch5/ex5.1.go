// 표준 입출력과 fmt패키지
// 입출력(프린터, 키보드 등)응 경우에 따라 직접 만드는 것 보다 표준 입출력 스티림을 이용하염 보다 효율적으로 처리할 수 있다.
// 표준 입출력을 사용하면, 목적지 상관없이 간편하게 입출력을 처리할 수 있다.
package main

import "fmt"

func main(){
	var a int = 10
	var b int = 20
	var f float32 = 32799438743.8297

	//서식에 맞추어서 출력
	fmt.Print("Print ", "a :", a, "b:", b,"\n")
	//서식에 맞추어 출력 + 개행
	fmt.Println("Println a :", a, "b:", b, "f:",f,"\n" )
	//Printf는 주어진 사용자 서식에 맞추어 입력값을 출력한다.
	fmt.Printf("Print : %d b: %d f:%f\n", a,b,f )
	//

}
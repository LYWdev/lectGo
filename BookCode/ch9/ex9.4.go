package main

import "fmt"

	var cnt int = 1

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()",cnt)
	cnt ++
	return cnt
}

func main() {
	//쇼트 서킷
	//값만 비교해야한다.
	// 함수는 값을 비교만하라, 함수 내부의 지역변수는 변해도 되지만 그 외값은 조작하면 안된다.

	if cnt == 0  && IncreaseAndReturn()<5{
		fmt.Println("1cnt 1증가")
	}else if cnt == 1 && IncreaseAndReturn()<5{
		fmt.Println(" && 기준 좌측이 거짓이라서 실행안함 흥!")
	}


//if cnt == 0  || IncreaseAndReturn()<5{
//	fmt.Println(" || 왼쪽거 맞으면 그냥 넘어가는 경우 있음 이게 소트서킷")
//}else if cnt == 1 || IncreaseAndReturn()<5{
//	fmt.Println(" && 기준 좌측이 거짓이라서 실행안함 흥!")
//}


}
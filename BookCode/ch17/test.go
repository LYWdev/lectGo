package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	가진돈은 1000으로 시작
	1~5사이의 값을 입력받는다.
	1~5사이 랜덤 값을 선택한다.
	입력값과 랜덤값이 같으면 가진돈에 500추가 및 잔액표시
	다를 경우 100마이너스, 아쉽다 및 잔액 표시
	0 이하 혹은 5000이상인 경우 게임종료
*/
func main(){
	rand.Seed(time.Now().UnixNano())
	r:=rand.Intn(5)+1//랜덤 값 생성
	for {
				time.Sleep(2*time.Second)
				fmt.Println("랜덤값 :",r)
		}
}
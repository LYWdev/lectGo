package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

const (
	//초기 잔액
	Balacne 			= 1000

	//정답일 때 잃는 양
	EarnPoint			= 500

	//오답 때 잃는 양
	LosePoint			= 100

	//게임 승리 포인트
	VictoryPoint 	= 5000

	//초기 잔액
	GameoverPoint	= 0
)  
var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int,error){
	var n int
	//int 타입값을 입력받음
	_, err := fmt.Scanln(&n)
	//에러 발생시 입력 스트림 비우기
	if err != nil	{
		stdin.ReadString('\n')
	}
	return n,err
}

func main()	{
	rand.Seed(time.Now().UnixNano())
	balance :=Balacne
	for {
		fmt.Println(" 1~5사이 값을 입력")
		n, err := InputIntValue()
		if err != nil{
			fmt.Println("숫자만 입력하세요")
		}else if n<1 || n>5 {
			fmt.Println("1~5사이의 값만 입력하세요")
		}else {
			r := rand.Intn(5)+1
			if n==r {
				balance +=EarnPoint
				fmt.Println("축하합니다. 정답 : ",balance)
				if balance >=VictoryPoint {
					fmt.Println("게임승리")
					break
				}
			}else{
				balance -=LosePoint
				fmt.Println(" 틀림 : ", balance)
				if balance <= GameoverPoint {
					fmt.Println(" 게임오버 : ", balance)
					break
				}
			}
		}
	}
}

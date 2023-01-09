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

const InitMoney int = 1000

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error){
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil{
		//input stream clear
		stdin.ReadString('\n')
	}
	return n, err
}

func main(){
	rand.Seed(time.Now().UnixNano())
	
	r:=rand.Intn(6)//랜덤 값 생성
	AddPoint := 500
	MinPoint := 100
	CurrentMoney := InitMoney

	for {
		fmt.Printf("type your number")

		n, err :=InputIntValue()

		if err != nil || n > 5 || n < 1{
			fmt.Println("1~5사이의 숫자 입력")

		}else{
			//에러가 아닌 정상적인 입력값을 받았을 때,
			//입력 값과 랜덤 값이 같으면 가진 돈에 500추가
			if n==r{
				CurrentMoney += AddPoint	
				fmt.Println("축하 잔액 :",CurrentMoney)
			//입력 값과 랜덤 값이 같지 않으면 -500
			}else if n!=r{
				CurrentMoney -= MinPoint

				//게임 종료
				if CurrentMoney <= 0{ 
					fmt.Println("게임종료 g잔액 :",CurrentMoney)
					break
				}else{
					fmt.Println("n>r","잔액 : ",CurrentMoney)
				}
	}
}
}
}
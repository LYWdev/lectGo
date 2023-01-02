package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("입력")
		var number int
		_, err :=fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			
			//키보드 버퍼 비우기
			stdin.ReadString('\n')
			// 에러 발생시 이 조건문안에서 키보드 버퍼를 비우고 다시 초기 루프문으로 돌아감.
			continue
		}

		fmt.Printf("입력하신 숫자는%d입니다. \n",number)

		if number%2 == 0{
		fmt.Println("짝수입니다 나갑니다")
		break
		}
	}
fmt.Println("for문이 종료되었습니다")
}
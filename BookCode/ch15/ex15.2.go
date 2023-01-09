package main

import "fmt"

func main() {
	// ``로 묶으면 특수 문자가 동작하지 않는다.
	poet0 := `If you use an Back quart, special characters do not work.\n markdown is very easy`

	// 백쿼트에선는 여러 줄 표현에 특수문자가 필요없다.
	poet1 := `use an Back quart, special characters do not work.\n markdown is very easy
	and Back quart is not nedd special characters  
	so is  easy,
	`

	// 큰따옴표에서 여러 줄을 표현하려면 \n을 사용해야한다.
	poet2 := "안녕 고"
	
	fmt.Println(poet0)
	fmt.Println(poet1)
	fmt.Println(poet2)
	//len()으로 문자열 크기 알아내기

	fmt.Println("UTF-8에서 한글은 글자당 3바이트 차지 ",len(poet0))
	fmt.Println("UTF-8에서 영어는 글자당 1바이트 차지 ",len(poet1))

	poet3 := "Hello"
	runes := []rune(poet3)

	fmt.Println("글자수를 알아내는 방법으로 []rune 슬라이스 타입이 있음. ",len(poet3))
	fmt.Println(" []rune이 있음. Poet2의 글자수: 실제로는 배열의 요소 개수 리턴",len(runes))

	for i:=0; i<len(runes); i++{
		fmt.Printf("[%d]ruens :타입:%T 값%d 문자값:%c\n",i,runes[i],runes[i],runes[i])
		fmt.Printf("[%d]poet3 : 타입:%T 값%d 문자값:%c\n",i,poet3[i],poet3[i],poet3[i])
	}

	for i, v:= range	poet3{
		fmt.Printf("USER range [%d]poet3 : 타입:%T 값%d 문자값:%c\n",i,v,v,v)
	}
}

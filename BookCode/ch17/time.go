package main

import (
	"fmt"
	"time"
)

func main(){
	//타임존을 나타냄
	loc, _:=time.LoadLocation("Asia/Seoul")

	//상수로 월 일, 년 시간에 대한 형식을 지정함
	const longForm ="Jan 2, 2006 at 3:04pm"

	//포맷, 실제 시간값, 타임존 변수로 입력
	t1, _ := time.ParseInLocation(longForm,"Jun14, 2021 at 10:00pm",loc)
	// t1을 출력한다. 타임존, 그리치천문대 기준
	fmt.Println(t1, t1.Location,t1.UTC())

	const shortForm ="2006-Jan-02"
	t2, _ := time.Parse(shortForm,"2021-Jun-14")
	fmt.Println(t2, t2.Location())

	t3, _ := time.Parse("2021-Jun-02","2021-Jun-14")
	fmt.Println(t3, t3.Location())

}
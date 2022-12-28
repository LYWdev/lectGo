package main

import "fmt"

	func PrintAvgScore(name string, math int, eng int, history int){
		total := math + eng + history
		avg := total/3
		fmt.Println(name, "의 avg :", avg)
	}
func main() {
	PrintAvgScore("1번", 80, 22, 25)
	PrintAvgScore("2번", 70, 20, 95)
	PrintAvgScore("3번", 60, 22, 95)
	PrintAvgScore("4번", 50, 10, 85)
	PrintAvgScore("5번", 40, 20, 95)
}

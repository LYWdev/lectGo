package main

import "fmt"

// 타입을 만들 수 잇음.
type ColorType int
const (
	Red ColorType = iota //0
	Blue 		//1	
	Green 		//2	 
	Yello 		//3	
)
func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "RED"
	case Blue:
		return "BLUE"
	case Green:
		return "Green"
	case Yello:
		return "Yello"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	fmt.Println("My Favorite color is ",colorToString(getMyFavoriteColor()))
	fmt.Println("My Favorite color is ",colorToString(4))
	fmt.Println("My Favorite color is ",colorToString(3))
	fmt.Println("My Favorite color is ",colorToString(2))
	fmt.Println("My Favorite color is ",colorToString(1))
	fmt.Println("My Favorite color is ",colorToString(0))
}



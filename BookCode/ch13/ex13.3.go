package main

import "fmt"

//구조제 정의
type User struct {
	Name string
	ID string
	Age int
}

type VIPUser struct {
	User
	VIPLevel int
	Price int
}

func main() {
	user := User{"송하나", "hana", 23}

	vip := VIPUser{
		User{"송둘", " two", 23},
		3,
		250,
	}
	fmt.Printf("유저 %s ID: %s 나이:%d\n",user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨 : %d VIP 가격: %d만 원\n",
	vip.Name,
	vip.ID,
	vip.Age,
	vip.VIPLevel,
	vip.Price,
)

}
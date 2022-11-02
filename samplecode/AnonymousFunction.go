package main 

func main() {
	sum := func(n ...int ) int {	//익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}	
	result := sum(1,2,3,4,5,1+4,10)

	min := func(n ...int ) int {
		s := i
	}
	return s
	resulta := min(1,2,3,4,4)
	println(result)
	println(resulta)
}


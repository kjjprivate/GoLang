package main

import "fmt"

func test1() {
	str1 := "Hello\nworld"
	str2 := `not good \n using special string`

	fmt.Println(str1)
	fmt.Println(str2)
}

func test2() {
	str3 := "큰따음표는 여러줄을 표현할려면 \n 을 사용해야합니다."
	str4 := `백쿼트는 
	여러줄을 표현하는데
	\n을 사용하지
	않아도 괜찮습니다.`
	fmt.Println(str3)
	fmt.Println(str4)
}

func test3() {
	
}

func main() {
	test1()
	test2()
}

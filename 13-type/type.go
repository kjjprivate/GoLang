package main

import "fmt"

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func test1() {
	var house House
	house.Address = "서울시 강동구..."
	house.Size = 28
	house.Price = 9.8
	house.Type = "아파트"

	fmt.Println(house.Address)
	fmt.Println(house.Size)
	fmt.Println(house.Price)
	fmt.Println(house.Type)

	//	var house House = House{"test", 10, 1.9, "아파트아니다"}
	fmt.Println(house)
}

func setUp() {
	type User struct {
		Name string
		ID   string
		Age  int
	}
	type VIPUser struct {
		UserInfo User
		VIPLevel int
		price    int
	}

	user := User{"송혜빈", "혜삔", 26}
	vip := VIPUser{
		User{"혜빈", "hyebin", 26},
		5,
		100000000,
	}
	fmt.Printf("유저: %s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID:%s 나이:%d viplevel:%d vipPrice:%d\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.price)
}

func setUp2() {
	type User struct {
		Name string
		ID   string
		Age  int
	}
	type VIPUser struct {
		User
		VIPLevel int
		price    int
	}

	user := User{"송혜빈", "혜삔", 26}
	vip := VIPUser{
		User{"혜빈", "hyebin", 26},
		5,
		100000000,
	}
	fmt.Printf("유저: %s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID:%s 나이:%d viplevel:%d vipPrice:%d\n", vip.Name, vip.ID, vip.Age, vip.VIPLevel, vip.price)
}

func main() {
	test1()
	setUp()
	setUp2()
}

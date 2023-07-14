package main

import "fmt"

func test1() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}

func test2() {
	var nums [5]int
	days := [3]string{"day1", "day2", "day3"}
	var temps [5]float64 = [5]float64{24.3, 26.7}
	var s = [5]int{1: 10, 3: 30}
	x := [...]int{10, 20, 30}
	fmt.Println(nums)
	fmt.Println(days)
	fmt.Println(temps)
	fmt.Println(s)
	fmt.Println(x)
}

// func test3() {
// 	const Y int = 3
// 	x := 5
// 	a := [x]int{1, 2, 3, 4, 5}
// 	b := [y]int{1, 2, 3}
// 	var c [10]int
// }

func test3() {
	t := [...]float64{24.0, 25.9, 27.8, 26.8, 21.2}
	for i, v := range t {
		fmt.Println(i, v)
	}
	for _, v := range t {
		fmt.Println(v)
	}
}

func test4() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 600, 700, 800, 900}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	fmt.Println()
	for i, v := range b {
		fmt.Printf("a[%d] = %d\n", i, v)
	}
	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d]=%d\n", i, v)
	}
}

func test5() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println()
	test1()
	test2()
	test3()
	test4()
	test5()
}

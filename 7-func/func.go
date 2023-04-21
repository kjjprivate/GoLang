package main

import "fmt"

func Add(a int, b int) int {
	return a + b

}

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// 재귀호출
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)

	result, success := Divide(9, 3)
	fmt.Println(result, success)
	result2, success2 := Divide(9, 0)
	fmt.Println(result2, success2)

	printNo(5)
}

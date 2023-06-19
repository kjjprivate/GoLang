package main

import "fmt"

var cnt = 0

// short circuit
func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

////////////////////////////////////////////////////////////

func main() {

}

package main

import "fmt"

func getMyAge() int { return 22 }

func main() {
	// switch age := getMyAge(); age {
	// case 10:
	// 	fmt.Println("Teenage")
	// case 33:
	// 	fmt.Println("Pair 3")
	// default:
	// 	fmt.Println("My age is", age)
	// }

	switch age := getMyAge(); true {
	case age < 10:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("10s")
	case age < 30:
		fmt.Println("20s")
	default:
		fmt.Println("MY age is", age)
	}

}

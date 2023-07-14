package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		return
	}
}

func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func test2() string {
	for i := 10; i > 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d=%d\n", i, j, i*j)
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Printf("%d * %d = %d\n", i, i, i*i)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j > 5-i; j++ {
			fmt.Print("*")
		}
	}

	return "END"
}

func main() {

	//stdin := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Println("입력하세요")
	// 	var number int
	// 	_, err := fmt.Scanln(&number)
	// 	if err != nil {
	// 		fmt.Println("숫자를 입력하세요.")
	// 		stdin.ReadString('\n')
	// 		continue
	// 	}

	// 	fmt.Printf("입력하신 숫자는 %d입니다.", number)
	// 	if number%2 == 0 {
	// 		break
	// 	}
	// }
	// fmt.Println("for 문이 종료되었습니다.")

	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// dan := 2
	// for {
	// 	for {
	// 		fmt.Printf("%d *%d = %d\n", dan, b, dan*b)
	// 		b++
	// 		if b == 10 {
	// 			break
	// 		}
	// 	}
	// 	b = 1
	// 	dan++
	// 	if dan == 10 {
	// 		break
	// 	}
	// }

	// a := 1
	// b := 1

	test2()
	// found := false
	// for a = 1; a <= 9; a++ {
	// 	for b = 1; b <= 9; b++ {
	// 		if a*b == 45 {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	if found {
	// 		break
	// 	}
	// }
	// fmt.Printf("%d * %d = %d\n", a, b, a*b)

	// outerFor:
	//
	//	for ; a <= 9; a++ {
	//		for b = 1; b <= 9; b++ {
	//			if a*b == 45 {
	//				break outerFor
	//			}
	//		}
	//	}
	//	fmt.Printf("%d * %d =%d\n", a, b, a*b)

	// for ; a <= 9; a++ {
	// 	var found bool
	// 	if b, found = find45(a); found {
	// 		break
	// 	}
	// }
	// fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

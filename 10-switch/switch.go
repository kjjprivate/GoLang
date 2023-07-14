package main

import "fmt"

func getMyAge() int { return 22 }

// switch 초기문; 비굣값 {
// case 값1:
// case 값2:
// default:
// }

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func test1() string {
	switch age := getMyAge(); age {
	case 10:
		return "Teenage"
	case 33:
		return "Pair 3"
	default:
		return "My age is"
	}

}

func test2() string {
	switch age := getMyAge(); true {
	case age < 10:
		return "child"
	case age < 20:
		return "10"
	case age < 30:
		return "20"
	default:
		return "MY age is ${age}"
	}
}

func test3(da string) string {
	day := da
	switch day {
	case "monday", "tuesday":
		return "월,화 싫어"
	case "wednesday":
		return "수요일 혜빈 만난다~"
	case "thursday", "friday":
		return "목금 빨리 지나가라"
	default:
		return "기본"
	}
}

func test4(temp int) string {
	switch true {
	case temp < 10, temp > 30:
		return "바깥 활동하기 좋은 날씨가 아닙니다."
	case temp >= 10 && temp < 20:
		return "약간 추울 수 있으니 가벼운 겉옷을 준비하세요."
	case temp >= 15 && temp < 25:
		return "야외 활동하기 좋은 날씨입니다."
	default:
		return "따듯합니다."
	}
}

func test5(color ColorType) string {

	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyfavoriteColor() ColorType {
	return Blue
}

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func getDirection(angle float64) Direction {

	switch angle := angle; true {
	case angle >= 315:
		return North
	case angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 135:
		return East
	case angle > 135 && angle < 225:
		return South
	case angle >= 225 && angle < 315:
		return West
	default:
		return None
	}
}

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3("wednesday"))
	fmt.Println(test4(12))
	fmt.Println("내가 제일 좋아하는 색은", test5(getMyfavoriteColor()))
	fmt.Println("-------------------")
	fmt.Println(DirectionToString(getDirection(38.3)))
	fmt.Println(DirectionToString(getDirection(200)))

}

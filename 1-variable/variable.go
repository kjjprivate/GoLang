package main

import "fmt"

func main() {
	fmt.Println("Hello Go World")
	{
		var a int = 10
		var msg string = "Hello variable"
		fmt.Println(a, msg)
		a = 20
		msg = "goodmorngsn"
		fmt.Println(msg, a)
	}
	{
		var minwage int = 10
		var hour int = 8

		var income int = minwage * hour
		fmt.Println(minwage, hour, income)
	}
	{
		var a int = 3
		var b int // 초깃 값 생략 시 타입별 기본값으로 대체
		var c int = 4
		d := 5

		fmt.Println(a, b, c, d)
	}
	{
		var a = 3.123
		b := 345
		c := "test"
		d := 3.1234
		fmt.Println(a, b, c, d)
	}
	{
		/*
			a := 3
			var b float64 = 3.5
			var c int = b
			d := a * b
			var e int64 = 7
			f := a * e
			fmt.Println(a, b, c, d, f)
		*/
	}

	{
		a := 3
		var b float64 = 3.5

		var c int = int(b)
		d := float64(a * c)

		var e int64 = 7
		f := int64(d) * e

		var g int = int(b * 3)
		var h int = int(b) * 3
		fmt.Println(a, b, c, d, e, f, g, h)

	}

	{
		var a int16 = 346
		var b int8 = int8(a)
		fmt.Println(a)
		fmt.Println(b)
	}
	{
		var a float32 = 1234.532
		var b float32 = 3512.2145
		var c float32 = a * b
		var d float32 = c * 3

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	}
}

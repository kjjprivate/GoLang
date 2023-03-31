package main

import "fmt"
func main() {
	{
	var a int  = 10
	var b int = 20 
	var f float64 = 34514215312.3432
	
	fmt.Print("a:",a,"b:",b)
	fmt.Println("a:", a,"b:", b,"f:", f)
	fmt.Printf("a: %d b:%d f:%f\n",a,b,f)
	}
	{
		var a = 123
		var b = 456
		var c = 123456

		fmt.Printf("%5d,%5d\n",a,b)
		fmt.Printf("%05d,%05d\n",a,b)
		fmt.Printf("%-5d,%-05d\n",a,b)

		fmt.Printf("%5d,%5d\n",c,c)
		fmt.Printf("%05d,%05d\n",c,c)
	}
	{
		var a = 324.13455
		var c = 3.14

		fmt.Printf("%08.2f\n",a)
		fmt.Printf("%08.2g\n",a)
		fmt.Printf("%8.5g\n",a)
		fmt.Printf("%f\n",c)
	}
	{
		fmt.Println("----------------------------")

		str := "Hello\tGo\t\tWorld\n\"go\"isAwesome!\n"
		fmt.Print(str)
		fmt.Printf("%s",str)
		fmt.Printf("%q",str)
	}
}

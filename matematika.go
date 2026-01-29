package main

import "fmt"

func main() {
	var a = 20
	var b = 15
	var c = a + b
	var d = a * b
	var e = a - b
	var f = a / b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	var i = 10

	i += 10 // i = i + 10

	fmt.Println(i)

	i += 100 // i = i + 100

	fmt.Println(i)
}

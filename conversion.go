package main

import "fmt"

func main() {
	var nilai64 int64 = 32767
	var nilai32 int32 = int32(nilai64)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai32)
	fmt.Println(nilai16)

	var name = "Najib"
	var e uint8 = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}

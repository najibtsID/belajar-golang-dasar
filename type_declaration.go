package main

import "fmt"

func main() {
	type noKTP string

	var ktpNajib noKTP = "111111"

	var contoh string = "22222"

	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpNajib)
	fmt.Println(contohKtp)
}

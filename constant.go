package main

import "fmt"

func main() {
	const (
		firstName string = "Najib"
		lastName         = "Baik"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//error
	//firstName = "Najib"
	//lastName = "Tidak Baik"
}

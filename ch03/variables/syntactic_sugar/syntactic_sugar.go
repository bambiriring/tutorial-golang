package main

import "fmt"

func main() {
	var a, b, c int32 = 1, 2, 3
	d, e, f := 4, int(5), int32(6)

	fmt.Println(a, b, c, d, e, f)

	var (
		name          = "Gil Dong, Hong"
		age     int32 = 30
		address       = "Seoul, Republic of Korea"
	)

	fmt.Println(name, age, address)
}

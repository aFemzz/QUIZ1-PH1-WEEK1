package main

import (
	"fmt"
)

func main() {
	// Quiz1()
	Quiz2()
}

func Quiz1() {
	var nama string = "nama"
	var umur int = 22

	fmt.Println(nama)
	fmt.Println(umur)
}

func Quiz2() {
	var buah map[string]int
	buah = map[string]int{}

	buah["Apel"] = 10
	buah["Nanas"] = 20
	buah["Pisang"] = 30

	fmt.Println("Apel", buah["Apel"])
	fmt.Println("Nanas", buah["Nanas"])
	fmt.Println("Pisang", buah["Pisang"])
}

// func Quiz3() {
// 	var number int
// }

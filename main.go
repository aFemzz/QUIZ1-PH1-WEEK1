package main

import (
	"fmt"
)

func main() {
	// Quiz1()
	// Quiz2()
	// Quiz3()
	Quiz4()
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

func Quiz3() {
	var number int
	number = 10
	if number%2 == 0 {
		fmt.Printf("Nomor berikut %d adalah Genap.\n", number)
	} else {
		fmt.Printf("Nomor berikut %d adalah Ganjil.\n", number)
	}
}

func Quiz4() {
	fmt.Print("Isi angka bulan : ")
	var bulan int
	fmt.Scanln(&bulan)

	switch {
	case bulan == 1:
		fmt.Println("Januari")
	case bulan == 2:
		fmt.Println("Februari")
	case bulan == 3:
		fmt.Println("Maret")
	case bulan == 4:
		fmt.Println("April")
	case bulan == 5:
		fmt.Println("Mei")
	case bulan == 6:
		fmt.Println("Juni")
	case bulan == 7:
		fmt.Println("Juli")
	case bulan == 8:
		fmt.Println("Agustus")
	case bulan == 9:
		fmt.Println("September")
	case bulan == 10:
		fmt.Println("October")
	case bulan == 11:
		fmt.Println("November")
	case bulan == 12:
		fmt.Println("Desember")
	}

}

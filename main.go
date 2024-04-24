package main

import (
	"fmt"
)

func main() {
	// Quiz1()
	// Quiz2()
	// Quiz3()
	// Quiz4()
	// Quiz5()
	// Quiz6()
	// Quiz7()
	// Quiz8()
	Quiz9()
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

func Quiz5() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func Quiz6() {
	fmt.Print("Masukan Deret : ")
	var n int
	fmt.Scanln(&n)

	var temp int = 1

	for i := 1; i <= n; i++ {
		temp = temp * 2

		fmt.Println(temp)
	}
}

func Quiz7() {
	// Gunakan perulangan dan pengkondisian untuk mencetak angka dari 1 hingga 20,
	// tetapi cetak "Fizz" untuk angka yang habis dibagi 3 dan "Buzz"
	// untuk angka yang habis dibagi 5.

	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("Nomor Sekarang :", i)
			fmt.Println("fiizzzzz")
		}
		if i%5 == 0 {
			fmt.Println("Nomor Sekarang :", i)
			fmt.Println("buuzzzzz")
		}
	}
}

func Quiz8() {
	// Tuliskan sebuah program yang mencetak bilangan dari 1 sampai N,
	// tetapi untuk kelipatan 3 cetak "Fizz", untuk kelipatan 5 cetak "Buzz",
	// dan untuk kelipatan keduanya cetak "FizzBuzz".

	fmt.Print("Isi angka bulan : ")
	var angka int
	fmt.Scanln(&angka)

	for i := 1; i <= angka; i++ {
		if i%3 == 0 {
			fmt.Println("Nomor Sekarang :", i)
			fmt.Println("fiizzzzz")
		}
		if i%5 == 0 {
			fmt.Println("Nomor Sekarang :", i)
			fmt.Println("buuzzzzz")
		}
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("Nomor Sekarang :", i)
			fmt.Println("fiizzzzz buuzzzzz")
		}
	}
}

func Quiz9() {
	// Buatlah sebuah program yang mencetak pola kotak dengan lapisan luar terdiri dari karakter *,
	// sedangkan bagian dalamnya kosong/hollow (menggunakan karakter spasi), dengan ukuran N x N.
	fmt.Print("Isi dimensi : ")
	var dimensi int
	fmt.Scanln(&dimensi)

	for i := 1; i <= dimensi; i++ {
		fmt.Println("*")
	}
}

// Print("*")x5
// Print(" ")

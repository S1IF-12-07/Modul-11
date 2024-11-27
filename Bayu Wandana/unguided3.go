package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&input)
	switch {
	case input%2 != 0 && input < 25:
		hasil := input + (input + 1)
		fmt.Printf("Kategori: bilangan ganjil\nHasilpenjumlahan dengan bilangan berikutnya %d + %d = %d\n", input, input+1, hasil)
	case input%2 == 0 && input%5 != 0 && input%10 != 0:
		hasil := input * (input + 1)
		fmt.Printf("Kategori: bilangan genap\nHasil perkaliandengan bilangan berikutnya %d * %d = %d\n", input, input+1, hasil)
	case input%5 == 0 && input%10 != 0:
		Hasil := (input * input)
		fmt.Printf("Hasil kuadrat dari 5\nHasil kuadrat dari %d^2 == %d", input, Hasil)
	case input%10 == 0:
		hasil := input / 10
		fmt.Printf("Kategori bilangan kelipatan 10 \nHasil pembagian antara %d / %d = %d", input, input/10, hasil)
	}
}

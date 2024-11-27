package main

import "fmt"

func main() {
	var integer int
	fmt.Print("integer  : ")
	fmt.Scan(&integer)

	keSepuluh := integer%10 == 0 && integer > 10
	keLima := integer%5 == 0 && integer > 6
	genap := integer%2 == 0
	ganjil := integer%2 != 0

	switch {
	case keSepuluh:
		hasil := integer / 10
		fmt.Printf("Kategori : bilangan kelipatan 10\nHasil pembagian antara %d / 10 = %d\n", integer, hasil)
	case keLima:
		hasil := integer * integer
		fmt.Printf("Kategori : bilangan kelipatan 5\nHasil kuadrat dari %d ^ 2 = %d\n", integer, hasil)
	case genap:
		integer2 := integer + 1
		hasil := integer * integer2
		fmt.Printf("Kategori : bilangan genap\nHasil perkalian dengan bilangan berikutnya %d * %d = %d\n", integer, integer2, hasil)
	case ganjil:
		integer2 := integer + 1
		hasil := integer + integer2
		fmt.Printf("Kategori : bilangan ganjil\nHasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", integer, integer2, hasil)
	}
}

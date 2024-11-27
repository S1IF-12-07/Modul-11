package main

import "fmt"

func main() {
	var angka int
	fmt.Print("masukkan angka : ")
	fmt.Scan(&angka)

	kel10 := angka%10 == 0 && angka > 10
	kel5 := angka%5 == 0 && angka > 5
	ganjil := angka%2 != 0
	genap := angka%2 == 0

	switch {
	case kel10:
		hasil := angka / 10
		fmt.Println("Kategori = Bilangan kelipatan 10")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d / 10 = %d", angka, hasil)
	case kel5:
		hasil := angka * angka
		fmt.Println("Kategori = Bilangan kelipatan 5")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d^2 = %d", angka, hasil)
	case ganjil:
		angka2 := angka + 1
		hasil := angka + angka2
		fmt.Println("Kategori = Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d + %d = %d", angka, angka2, hasil)
	case genap:
		angka2 := angka + 1
		hasil := angka * angka2
		fmt.Println("Kategori = Bilangan genap")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d * %d = %d", angka, angka2, hasil)
	}
}

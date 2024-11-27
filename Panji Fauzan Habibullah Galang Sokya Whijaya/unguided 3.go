package main

import "fmt"

func main() {
	var bil, hasil int
	fmt.Scan(&bil)
	switch {
	case bil%10 == 0 && bil > 10:
		hasil = bil / 10
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Print("Hasil pembagian antara ", bil, " / ", 10, " = ", hasil)
	case bil%2 == 0:
		hasil = (bil + 1) * bil
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Print("Hasil perkalian dengan bilangan berikutnya ", bil, " * ", bil+1, " = ", hasil)
	case bil%5 == 0 && bil > 5:
		hasil = bil * bil
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Print("Hasil kuadrat dari ", bil, " ^2 = ", hasil)
	default:
		hasil = (bil + 1) + bil
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Print("Hasil penjumlahan dengan bilangan berikutnya ", bil, " + ", bil+1, " = ", hasil)
	}
}

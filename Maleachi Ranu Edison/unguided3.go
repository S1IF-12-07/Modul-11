package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a%10 == 0:
		fmt.Println("kategori: Bilangan Kelipatan 10")
		perhitungan := a / 10
		fmt.Printf("Hasil pembagian antara %d / 10 = %d", a, perhitungan)
		break
	case a%5 == 0 && a != 5:
		fmt.Println("kategori: Bilangan Kelipatan 5")
		perhitungan := a * a
		fmt.Printf("Hasil perpangkatan antara %d ^ 2 = %d", a, perhitungan)
		break
	case a%2 == 0:
		fmt.Println("kategori: Bilangan Genap")
		bilanganberikutnya := a + 1
		perhitungan := a * bilanganberikutnya
		fmt.Printf("Hasil dengan bilangan berikutnya %d * %d = %d", a, bilanganberikutnya, perhitungan)
		break
	default:
		fmt.Println("kategori: Bilangan Ganjil")
		bilanganberikutnya := a + 1
		perhitungan := a + bilanganberikutnya
		fmt.Printf("Hasil  penjumlahan dengan bilangan berikutnya %d + %d = %d", a, bilanganberikutnya, perhitungan)

	}
}

package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan n : ")
	fmt.Scan(&n)

	bagi := n%10 == 0 && n > 10
	kuadrat := n%5 == 0 && n > 5
	ganjil := n%2 != 0
	genap := n%2 == 0

	switch {
	case bagi:
		hasil := n / 10
		fmt.Println("Kategori = Bilangan kelipatan 10")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d / 10 = %d", n, hasil)
	case kuadrat:
		hasil := n * n
		fmt.Println("Kategori = Bilangan kelipatan 5")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d^2 = %d", n, hasil)
	case ganjil:
		n2 := n + 1
		hasil := n + n2
		fmt.Println("Kategori = Bilangan ganjil")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d + %d = %d", n, n2, hasil)
	case genap:
		n2 := n + 1
		hasil := n * n2
		fmt.Println("Kategori = Bilangan genap")
		fmt.Printf("Hasil penjumlahan bilangan berikutnya %d * %d = %d", n, n2, hasil)
	}
}

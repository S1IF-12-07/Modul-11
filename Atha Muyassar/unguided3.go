package main
import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	bagi := angka %10 == 0 && angka > 10
	kuadrat := angka %5 == 0 && angka > 5
	ganjil := angka %2 != 0
	genap := angka %2 == 0

	switch {
	case bagi :
		hasil := angka / 10
		fmt.Println("kelipatan 10")
		fmt.Printf("hasil penjumlahan bilangan berikutnya %d / 10 = %d", angka, hasil)
	case kuadrat :
		hasil := angka * angka
		fmt.Println("kelipatan 5")
		fmt.Printf("hasil penjumlahan bilangan berikutnya %d^2 = %d", angka, hasil)
	case ganjil :
		angka2 := angka + 1
		hasil := angka + angka2
		fmt.Println("bill ganjil")
		fmt.Printf("hasil penjumlahan bilangan berikutnya %d + %d = %d", angka, angka2, hasil)
	case genap :
		angka2 := angka + 1
		hasil := angka * angka2
		fmt.Println("bill genap")
		fmt.Printf("hasil penjumlahan bilangan berikutnya %d * %d = %d", angka, angka2, hasil)
	}
} 
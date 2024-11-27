package main
import "fmt"
func main() {
	var kendaraan string
	var waktu int
	fmt.Scan(&kendaraan, &waktu)
	switch kendaraan {
	case "motor":
		motor := 2000 * waktu
		fmt.Println("Rp ", motor)
		break
	case "mobil":
		mobil := 5000 * waktu
		fmt.Println("Rp", mobil)
		break
	case "truk":
		truk := 8000 * waktu
		fmt.Println("Rp", truk)
		break
	default:
		fmt.Println("Input tidak valid")
	}
}
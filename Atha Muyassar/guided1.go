package main
import "fmt"

func main() {
	var usia int
	fmt.Scan(&usia)

	switch {
	case usia < 0:
		fmt.Println("belum lahir")
		break
	case usia < 13:
		fmt.Println("anak-anak")
		break
	case usia < 20:
		fmt.Println("remaja")
		break
	case usia < 60:
		fmt.Println("dewasa")
	case usia >= 60:
		fmt.Println("lansia")
	default:
		fmt.Println("Umur tidak terdaftar dalam kategori")
	}
}
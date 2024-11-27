package main
import "fmt"

func main()  {
	var kendaraan string
	var waktu int

	fmt.Scan(&kendaraan, &waktu)
	if waktu < 1 {
		waktu = 1
	}

	switch kendaraan {
	case "motor":
		motor:= 2000 * waktu
		fmt.Println("tarif parkir : Rp", motor)
	case "mobil":
		mobil:= 5000 * waktu
		fmt.Println("tarif parkir : Rp", mobil)
	case "truk":
		truk:= 8000 * waktu
		fmt.Println("tarif parkir : Rp", truk)
	}
}
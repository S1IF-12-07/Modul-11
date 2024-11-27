package main

import "fmt"

func main() {
    var hari int
    fmt.Print("Masukkan angka hari (1-7): ")
    fmt.Scanln(&hari)

    switch hari {
    case 1:
        fmt.Println("Hari Senin")
    case 2:
        fmt.Println("Hari Selasa")
    case 3:
        fmt.Println("Hari Rabu")
    case 4:
        fmt.Println("Hari Kamis")
    case 5:
        fmt.Println("Hari Jumat")
    case 6:
        fmt.Println("Hari Sabtu")
    case 7:
        fmt.Println("Hari Minggu")
    default:
        fmt.Println("Angka hari tidak valid. Harus antara 1 dan 7.")
	}
}
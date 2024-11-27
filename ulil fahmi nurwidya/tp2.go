package main

import "fmt"
func getHari(kode int) string {
    switch kode {
    case 1:
        return "Senin"
    case 2:
        return "Selasa"
    case 3:
        return "Rabu"
    case 4:
        return "Kamis"
    case 5:
        return "Jumat"
    case 6:
        return "Sabtu"
    case 7:
        return "Minggu"
    default:
        return "Kode Tidak Ditemukan"
    }
}

func main() {
    var kode int

    fmt.Print("Masukkan kode hari: ")
    fmt.Scan(&kode)

    hari := getHari(kode)
    fmt.Println("Hari:", hari)
}

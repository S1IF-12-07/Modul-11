package main

import "fmt"

func kategoriUmur(umur int) string {
    switch {
    case umur == 0:
        return "Tidak Terdefinisi"
    case umur < 13:
        return "Anak-Anak"
    case umur < 20:
        return "Remaja"
    case umur < 60:
        return "Dewasa"
    case umur >= 60:
        return "Lansia"
    default:
        return "Data tidak valid"
    }
}

func main() {
    var umur int

    fmt.Print("Masukkan umur: ")
    fmt.Scan(&umur)

    fmt.Println("Termasuk dalam kategori:", kategoriUmur(umur))
}
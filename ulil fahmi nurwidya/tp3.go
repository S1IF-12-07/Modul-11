package main

import "fmt"

func deskripsiFilm(kode string) string {
    switch kode {
    case "G":
        return "untuk semua umur"
    case "PG":
        return "untuk anak-anak di atas 7 tahun"
    case "PG-13":
        return "untuk remaja di atas 13 tahun"
    case "R":
        return "untuk dewasa"
    default:
        return "kode tidak ditemukan"
    }
}

func main() {
    var kode string
    fmt.Print("Masukkan kode film: ")
    fmt.Scan(&kode)
    fmt.Println("Kategori:", deskripsiFilm(kode))
}

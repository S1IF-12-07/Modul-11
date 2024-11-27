package main

import "fmt"

func cekKelipatan10(angka int) (bool, int) {
    if angka%10 == 0 && angka > 10 {
        return true, angka / 10
    }
    return false, 0
}

func cekKelipatan5(angka int) (bool, int) {
    if angka%5 == 0 && angka > 5 {
        return true, angka * angka
    }
    return false, 0
}

func cekGanjil(angka int) (bool, int, int) {
    if angka%2 != 0 {
        angka2 := angka + 1
        return true, angka, angka + angka2
    }
    return false, 0, 0
}
func cekGenap(angka int) (bool, int, int) {
    if angka%2 == 0 {
        angka2 := angka + 1
        return true, angka, angka * angka2
    }
    return false, 0, 0
}

func main() {
    var angka int
    fmt.Print("Masukkan angka: ")
    fmt.Scan(&angka)

    if ok, hasil := cekKelipatan10(angka); ok {
        fmt.Println("Kategori = Bilangan kelipatan 10")
        fmt.Printf("Hasil pembagian berikutnya %d / 10 = %d\n", angka, hasil)
    } else if ok, hasil := cekKelipatan5(angka); ok {
        fmt.Println("Kategori = Bilangan kelipatan 5")
        fmt.Printf("Hasil kuadrat berikutnya %d^2 = %d\n", angka, hasil)
    } else if ok, angka1, hasil := cekGanjil(angka); ok {
        fmt.Println("Kategori = Bilangan ganjil")
        fmt.Printf("Hasil penjumlahan bilangan berikutnya %d + %d = %d\n", angka1, angka1+1, hasil)
    } else if ok, angka1, hasil := cekGenap(angka); ok {
        fmt.Println("Kategori = Bilangan genap")
        fmt.Printf("Hasil perkalian bilangan berikutnya %d * %d = %d\n", angka1, angka1+1, hasil)
    }
}

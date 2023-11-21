package main

import ("fmt")

func main() {
	var A, B, TP int
	var Lingkaran bool

	print("Masukkan Jari-Jari Lingkaran A : ")
	fmt.Scan(&A)

	print("Masukkan Jari-Jari Lingkaran B : ")
	fmt.Scan(&B)

	print("Masukkan Jarak Titik Pusat Liingkaran A dan B : ")
	fmt.Scan(&TP)

	Lingkaran = (A + B) <= TP // Rumus untuk mengetahui hasil tidak menghasilkan irisan

	fmt.Println(Lingkaran)
}
package main

import ("fmt")

func main() {
	var input int
	var kabisat bool

	print("Masukkan Tahun : ")
	fmt.Scan(&input)

	kabisat = input % 400 == 0 || input % 4 == 0 && input % 100 != 0

	fmt.Printf("%v", kabisat) 
}
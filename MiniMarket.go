package main

import ("fmt")

func main() {
	var total int
	var kartu, diskon, cashback bool

	print("Masukkan total belanja : ")
	fmt.Scan(&total)

	print("Apakah bersedia dibuatkan kartu ? (true/false) :  ")
	fmt.Scan(&kartu)

	diskon = total >= 100000
	cashback = total >= 200000 && kartu == true

	fmt.Println("================================")
	fmt.Println("Kartu ?", kartu)
	fmt.Println("Diskon ?", diskon)
	fmt.Println("Cashback ?", cashback)
}
package main

import ("fmt")

func main() {
	var a, b, c int
	var segitiga bool

	print("Masukkan nilai a : ")
	fmt.Scan(&a)

	print("Masukkan nilai b : ")
	fmt.Scan(&b)

	print("Masukkan nilai c : ")
	fmt.Scan(&c)

	segitiga = a + b > c && b + c > a && c + a > b

	fmt.Println(segitiga)
}
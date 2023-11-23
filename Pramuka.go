package main

import ("fmt")

func main() {
	var N int
	var perlengkapan bool
	var item1, item2, item3, item4, item5 bool

	perlengkapan = true

	print("Masukkan Jumlah Anggota TIM : ")
	fmt.Scan(&N)

	for i:=1 ; i<=N && perlengkapan ; i++ {
		fmt.Println("Anggota ", i)
		print("Masukkan kelengkapan item 1 (true/false) : ")
		fmt.Scan(&item1)
		print("Masukkan kelengkapan item 2 (true/false) : ")
		fmt.Scan(&item2)
		print("Masukkan kelengkapan item 3 (true/false) : ")
		fmt.Scan(&item3)
		print("Masukkan kelengkapan item 4 (true/false) : ")
		fmt.Scan(&item4)
		print("Masukkan kelengkapan item 5 (true/false) : ")
		fmt.Scan(&item5)

		perlengkapan = perlengkapan && item1 && item2 && item3 && item4 && item5
	}
	fmt.Println(perlengkapan)
}

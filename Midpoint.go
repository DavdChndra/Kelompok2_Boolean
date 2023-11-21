package main  

import ("fmt") 

func main() { 
    var a, b, c int
    var Midpoint bool

    print("Masukkan Bilangan 1 : ")
    fmt.Scan(&a)
    print("Masukkan Bilangan 2 : ")
    fmt.Scan(&b)
    print("Masukkan Bilangan 3 : ")
    fmt.Scan(&c)

    Midpoint = (a <= b && b <= c) || (c <= b && b <= a) || (b >= a && a >= c)

    fmt.Println(Midpoint)

}


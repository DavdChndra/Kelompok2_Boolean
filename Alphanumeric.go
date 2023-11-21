package main  

import ("fmt") 

func main() { 
    var input byte
    var alphanumeric bool 

    fmt.Print("Masukkan sebuah alphanumeric: ")
    fmt.Scanf("%c", &input) 

    alphanumeric = input >= 'A' && input <= 'Z'|| input >= 'a' && input <= 'z' || input >= '1' && input <= '9'

    fmt.Printf("%v", alphanumeric) 
}


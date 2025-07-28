package main

import "fmt"

func main ()  {
	var value = (((2+6)%3)*4-2) / 3
	var value2 = 8 *4
	fmt.Print("Hasilnya adalah:",value, value2)

	a := 5
    b := 3
    c := 2
    
    result := (a*b + b%c) * (c - a/b) / (a%c + 1)
    fmt.Println("Hasil operasi campuran:", result)
    
}
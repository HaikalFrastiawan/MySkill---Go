package main

import "fmt"

func main()  {
	var nilaiVar int

	fmt.Print("Masukan Nilai Anda:")
	_, err := fmt.Scan(&nilaiVar)
	if err != nil{
		fmt.Println("input tidak valid, masukan angka")
		return
	}



	switch{
	case nilaiVar == 18:	
		fmt.Println("perfect")
	case nilaiVar >= 10: 
		fmt.Println("awesome")

	default: 
	fmt.Println("not bad")
	}
}
package main

import "fmt"

func main()  {
	var point int

	fmt.Print("Masukan Nilai Anda:")
	_, err := fmt.Scan(&point)
	if err != nil{
		fmt.Println("input tidak valid, masukan angka")
		return
	}



	if point == 10 {
		fmt.Println("lulus dengan sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus")
	}

}
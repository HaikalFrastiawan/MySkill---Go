package main

import "fmt"

func main()  {
	

	for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	}

	//perulangan range

	angka := [] int {1,2,3,4,5}

	for indeks, nilai := range angka{
		fmt.Printf("index %d: Nilai %d \n", indeks,nilai)
	}
}
package main

import "fmt"

func main()  {
	    // Konversi int ke float64
    a := 42
    b := float64(a)
    fmt.Printf("%T: %v\n", b, b) // float64: 42

    // Konversi float ke int (nilai pecahan terpotong)
    c := 3.14
    d := int(c)
    fmt.Printf("%T: %v\n", d, d) // int: 3
}
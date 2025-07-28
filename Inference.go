package main

import "fmt"

func main() {
    // Type inference dengan :=
    var nama string = "Alice"      // jika tanpa inference
    age := 30            // int
    height := 5.9        // float64
    isStudent := true    // bool

    fmt.Printf("%T: %v\n", nama, nama)       // string: Alice
    fmt.Printf("%T: %v\n", age, age)         // int: 30
    fmt.Printf("%T: %v\n", height, height)   // float64: 5.9
    fmt.Printf("%T: %v\n", isStudent, isStudent) // bool: true
}
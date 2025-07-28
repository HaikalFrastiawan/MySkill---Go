 package main

 import "fmt"
 
 func main()  {
	var firstName string ="Haikal "
	const lastName string = "Frastiawan"
	fmt.Print("hallo ", firstName, lastName)

	//tipe data
	var bilanganBulat uint8 = 20
	var decimalNumber = 3.60
	var tipeBool = true


	fmt.Print("\n ","umur saya: ", bilanganBulat)
	fmt.Print("\n ","Ipk saya:", decimalNumber)
	fmt.Print("\n ","Saya ganteng: ", tipeBool)

	//variabel pointer
	var numberA int = 4
	var numberB *int = &numberA
	
	fmt.Print("\nnumber A (value) : ", numberA)
	fmt.Print("\nnumber A (address) : ", &numberA)
	fmt.Print("\nnumber B (value) : ", numberB)
	fmt.Print("\nnumber B (address) : ", &numberB)


 }
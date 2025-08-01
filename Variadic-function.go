package main

import "fmt"

func main()  {
	var avg = calculates(2,4,3,5,4,3,3,4,5,3,4)

	var msg = fmt.Sprintf("rata-rata: %.2f", avg)
	fmt.Println(msg)
}


func calculates (numbers... int)float64{
	var total int = 0
	for _, number := range numbers{
		total += number }

	var avg = float64(total) /float64(len(numbers))

	return avg
}
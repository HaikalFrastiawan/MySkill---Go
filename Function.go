package main

import (
	"fmt"
	"strings"
)


func main()  {
	var names = []string {"my", "skill"  }
	printMassage("halo", names)
}

func printMassage(massage string, arr []string)  {
	var nameString = strings.Join(arr, " ")
	fmt.Println(massage,nameString)
}
package main
import (
	"errors"
	"fmt"
	"strings"
)

func Validate (input string) (bool, error){
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return	true, nil
}

func main (){
	var name string
	fmt.Println("type your name:") 
	fmt.Scanln(&name)
	if valid, err:= Validate(name); valid{
		fmt.Println("halo", name)
	} else{
		panic(err.Error())
		fmt.Println("end")
	}
}

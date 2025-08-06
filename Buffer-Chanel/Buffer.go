package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2) // Mengatur jumlah core CPU yang digunakan menjadi 2

	messages := make(chan int, 3) // Membuat buffered channel kapasitas 3

	// Goroutine untuk menerima data dari channel
	go func() {
		for {
			i := <-messages // menerima data dari channel
			fmt.Println("receive data", i)
		}
	}()

	// Mengirim data ke channel
	for i := 0; i < 5; i++ {
		fmt.Println("Send data", i)
		messages <- i // mengirim data ke channel
	}
}

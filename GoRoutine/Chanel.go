package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Gunakan 2 core CPU untuk concurrency demo
	runtime.GOMAXPROCS(2)

	// Bikin channel string
	messages := make(chan string)

	// WaitGroup untuk tunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Fungsi yang akan dijalankan goroutine
	sayHelloTo := func(who string) {
		defer wg.Done() // Sinyal goroutine selesai
		message := fmt.Sprintf("hello %s", who)
		messages <- message
	}

	// Tambah 3 task ke WaitGroup
	wg.Add(3)

	// Jalankan goroutine
	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	// Goroutine tambahan untuk close channel setelah semua selesai
	go func() {
		wg.Wait()
		close(messages)
	}()

	// Terima semua pesan dari channel
	for msg := range messages {
		fmt.Println(msg)
	}
}

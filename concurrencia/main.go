package main

import (
	"fmt"
	"time"
)

func test(i int, text string) {
	fmt.Println(i, text)
}

func main() {
	for i := 1; i <= 10; i++ {
		go test(i, "hola")
		go test(i, "adios")
	}
	time.Sleep(1 * time.Second)
}

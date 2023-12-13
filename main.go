package main

import (
	"fmt"
	"goPlayground/poio"
)

func main() {
	written, err := poio.FileCopy("books.txt.bak", "books.txt")
	if err == nil {
		fmt.Printf("Copy done!\n%d bytes written.\n", written)
	}
}

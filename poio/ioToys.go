package poio

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// A small toy that can count words, lines, and characters.
// Input message is ended by capital 'S'.
// Feel free to switch lines.
func WordLetterCount() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">")
	input, err := reader.ReadString('S')
	if err != nil {
		fmt.Printf("Error occurs, reading exited.\n")
		return
	}
	wordCount := 0
	lineCount := 0
	charCount := 0
	inWord := false
	for index := range input {
		switch input[index] {
		case '\r', '\n':
			lineCount++
			inWord = false
		case ' ':
			charCount++
			inWord = false
		default:
			if !inWord && input[index] != 'S' {
				wordCount++
				inWord = true
			}
			if input[index] != 'S' {
				charCount++
			}
		}
	}
	if len(input) != 0 && input[len(input)-1] != '\n' && input[len(input)-1] != '\r' {
		lineCount++
	}
	fmt.Printf("There are %d characters, %d words and %d lines in the input message.\n", charCount, wordCount, lineCount)
}

type book struct {
	title  string
	price  float64
	amount int
}

// book generator
func NewBook(title string, price float64, amount int) book {
	return book{title, price, amount}
}

// string formatter to describe a book
func (book book) String() string {
	var rawInfo []byte
	rawInfo = fmt.Appendf(rawInfo, "%s, ¥%v, %d in storage.", book.title, book.price, book.amount)
	return string(rawInfo)
}

type Books []book

// To list all the books' detail in a booklist
func (books *Books) List() {
	for index, book := range *books {
		fmt.Printf("%d: %v\n", index, book)
	}
}

// To read a series of books from a file and returns a booklist
func ReadProductsFrom(filename string) Books {
	books := make(Books, 0)
	file, openErr := os.Open(filename)
	if openErr != nil {
		log.Fatalf("Error %s opening file \"%s\"", openErr, filename)
	}
	defer file.Close() // ensure the file is always closed after reading
	reader := bufio.NewReader(file)
	for {
		line, readErr := reader.ReadString('\n')
		line = string(line[:len(line)-1])
		info := strings.Split(line, ";")
		newBook := new(book)
		newBook.title = info[0]
		var eleErr error
		newBook.price, eleErr = strconv.ParseFloat(info[1], 64)
		if eleErr != nil {
			fmt.Printf("Error generating list: %s", eleErr)
		}
		newBook.amount, eleErr = strconv.Atoi(info[2])
		if eleErr != nil {
			fmt.Printf("Error generating list: %s", eleErr)
		}
		books = append(books, *newBook)
		if readErr == io.EOF {
			break
		}
	}
	fmt.Printf("We have read the following books from the file \"%s\"\n", filename)
	books.List()
	return books
}

// FileCopy make a copy of an existing file,
// returning bytes written and the first error occured, if any.
func FileCopy(dst, src string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

var lineNumber = flag.Bool("i", false, "number each line")

// core cat reader
func catCore(reader *bufio.Reader) {
	index := 1
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		// fmt.Fprintf(os.Stdout, "%s", line)
		if *lineNumber {
			fmt.Fprintf(os.Stdout, "%5d %s", index, line)
			index++
		} else {
			fmt.Fprintf(os.Stdout, "%s", line)
		}
	}
}

// Cat read a file and print line by line.
// Add file name to tell which file you want cat to read.
// With no file name cat will echo what the user typed
func Cat() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Printf("You have not appoint any file\nCat will echo what you type\nExit by typing ^D\n")
		*lineNumber = false
		catCore(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		defer file.Close()
		catCore(bufio.NewReader(file))
	}
}

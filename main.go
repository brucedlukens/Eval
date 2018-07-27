package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
	"log"
	"time"
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// Mostly stolen from stackoverflow with small changes.
func PrintMemUsage() {
	// convert bytes to megabytes
	bToMb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// This just does a simple pretty print
func PrettyPrintComparisonItem(item string) {
	fmt.Println()
	fmt.Println()
	fmt.Printf(item + ": ")
}

// This pretty prints the matched lines
func PrettyPrintMatchedLines(slice []int) {
	for index, item := range slice {
		if index > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf(strconv.Itoa(item))
	}
}

// This does the reading and comparing of files
func ReadFiles(file1 string, file2 string) (err error){
	// Opens the first file.
	f1, f1Err := os.Open(file1)
	if f1Err != nil {
		return f1Err
	}
	// Close f1 after the function is complete
	defer f1.Close()

	// Going to use bufio's scanner to iterate line by line
	f1Scanner := bufio.NewScanner(f1)

	for f1Scanner.Scan() {
		// Grab next item from file 1
		f1Item := strings.ToLower(f1Scanner.Text())
		PrettyPrintComparisonItem(f1Item)

		// Opens the second file.
		f2, f2Err := os.Open(file2)
		if f2Err != nil {
			return f2Err
		}
		// bufio scanner iterable
		f2Scanner := bufio.NewScanner(f2)

		// file 2 line number reset to 1 every new file 1 item
		f2Line := 1
		// Creating slice to store line numbers, seems easier to read this way instead of random print statements
		var slice []int

		for f2Scanner.Scan() {
			// Check if string is exactly equal to next line item
			if strings.EqualFold(strings.ToLower(f2Scanner.Text()), f1Item) {
				// Add line number to slice
				slice = append(slice, f2Line)
			}
			f2Line++
		}
		// Iterate through slice to print.
		PrettyPrintMatchedLines(slice)

		// Need to close and reopen file for every new file 1 item since scanner is basically a stream
		f2.Close()

		if err := f2Scanner.Err(); err != nil {
			return err
		}
	}

	if err := f1Scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	x := time.Now()
	//PrintMemUsage()
	fmt.Println("Starting file comparison")
    err := ReadFiles("names.txt", "list.txt")
    if err != nil {
		log.Print(err)
    }
	//fmt.Println()
	//fmt.Println()
	//PrintMemUsage()
	fmt.Println()
	fmt.Println()
	fmt.Println("Total time for comparisons: ", time.Since(x))
}
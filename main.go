package main

import (
    "fmt"
    // "io/ioutil"
    "time"
    "strings"
    "bufio"
    // "bytes"
    "os"
    "runtime"
    "strconv"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number 
// of garage collection cycles completed.
// Mostly stole from stackoverflow with small changes.
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

// This does the reading and comparing of files
func ReadFiles() (text string, err error){
	totalMatches := 0

	// Opens the names file.
	nameFile, nameErr := os.Open("names.txt")
	if nameErr != nil {
		return "", nameErr
	}
	// Close the file after the function is complete
	defer nameFile.Close()

	// Going to use bufio's scanner to iterate line by line
	nameScanner := bufio.NewScanner(nameFile)

	for nameScanner.Scan() {
		// Grab next name
		name := strings.ToLower(nameScanner.Text())
		
		// Opens the list file.
		list, listErr := os.Open("list.txt")
		if listErr != nil {
			return "", listErr
		}
		// bufio scanner iterable
		listScanner := bufio.NewScanner(list)

		// list line number reset to 1 every new name
		listLine := 1
		// Creating slice to store line numbers
		var tempSlice []int

		for listScanner.Scan() {
			// Grab next item
			item := strings.ToLower(listScanner.Text())

			// Check if string is exactly equal
			if strings.EqualFold(item, name) {
				// fmt.Println("match: ", line, item)
				// Add line to slice
				tempSlice = append(tempSlice, listLine)
				totalMatches++
			}
			listLine++
		}
		fmt.Println(tempSlice)

		// Need to close and reopen file for every new name since scanner is basically a stream
		list.Close()

		if err := listScanner.Err(); err != nil {
			return "", err
		}
	}

	if err := nameScanner.Err(); err != nil {
		return "", err
	}
	return strconv.Itoa(totalMatches), nil
}

func main() {
	x := time.Now()
	PrintMemUsage()

    text, err := ReadFiles()
    if err != nil {
    	fmt.Println(err)
    }
    if text != "" {
    	fmt.Println(text)
    }
	PrintMemUsage()
	fmt.Println(time.Since(x))
}
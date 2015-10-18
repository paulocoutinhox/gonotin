/*
AUTHOR: PAULO COUTINHO
*/

package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

// Main function
func main() {
	// general data
	var dataA = make(map[string]string)
	var dataB = make(map[string]string)

	// read file A
	fileA, errA := os.Open(os.Args[1])

	if errA != nil {
		log.Fatal(errA)
	}

	defer fileA.Close()

	scannerA := bufio.NewScanner(fileA)

	for scannerA.Scan() {
		dataA[scannerA.Text()] = scannerA.Text()
	}

	// read file B
	fileB, errB := os.Open(os.Args[2])

	if errB != nil {
		log.Fatal(errB)
	}

	defer fileB.Close()

	scannerB := bufio.NewScanner(fileB)

	for scannerB.Scan() {
		dataB[scannerB.Text()] = scannerB.Text()
	}

	// process and show data from A that not exists in B
	for _, valueA := range dataA {
		_, exists := dataB[valueA]

		if !exists {
			fmt.Println(valueA)
		}
	}
}

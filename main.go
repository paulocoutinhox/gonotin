package main
import (
	"os"
	"log"
	"bufio"
	"fmt"
	"sort"
)

func main() {
	// general data
	var dataA []string
	var dataB []string

	// read file A
	fileA, errA := os.Open(os.Args[1])

	if errA != nil {
		log.Fatal(errA)
	}

	defer fileA.Close()

	scannerA := bufio.NewScanner(fileA)

	for scannerA.Scan() {
		dataA = append(dataA, scannerA.Text())
	}

	// read file B
	fileB, errB := os.Open(os.Args[2])

	if errB != nil {
		log.Fatal(errB)
	}

	defer fileB.Close()

	scannerB := bufio.NewScanner(fileB)

	for scannerB.Scan() {
		dataB = append(dataB, scannerB.Text())
	}

	// debug
	fmt.Println("> Data in A: ", len(dataA), " | Data in B: ", len(dataB))

	// process and show data from A that not exists in B
	for _, valueA := range dataA {
		pos := sort.SearchStrings(dataB, valueA)

		if pos < 0 {
			fmt.Println(valueA)
		}
	}
}

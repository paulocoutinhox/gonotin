package main
import (
	"os"
	"log"
	"bufio"
	"fmt"
	"sort"
)

func main() {
	// select mode: '1' using array and sort, '2' to use map strategy (mode 2 is too fast) and 3 that is a poor impl.
	var mode = 2

	if mode == 1 {

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
		total := len(dataB)
		sort.Strings(dataB)

		for _, valueA := range dataA {
			pos := sort.SearchStrings(dataB, valueA)

			if !(pos < total && dataB[pos] == valueA) {
				fmt.Println(valueA)
			}
		}

	} else if mode == 2 {

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

		// debug
		fmt.Println("> Data in A: ", len(dataA), " | Data in B: ", len(dataB))

		// process and show data from A that not exists in B
		for _, valueA := range dataA {
			_, exists := dataB[valueA]

			if !exists {
				fmt.Println(valueA)
			}
		}

	} else if mode == 3 {

		// general data
		var dataA = make(map[string]int)
		var dataB = make(map[string]int)

		// read file A
		fileA, errA := os.Open(os.Args[1])

		if errA != nil {
			log.Fatal(errA)
		}

		defer fileA.Close()

		readerA := bufio.NewReader(fileA)

		for {
			line, _, err := readerA.ReadLine()
			if err != nil {
				break
			}
			dataA[string(line)] = 0
		}

		// read file B
		fileB, errB := os.Open(os.Args[2])

		if errB != nil {
			log.Fatal(errB)
		}

		defer fileB.Close()

		readerB := bufio.NewReader(fileB)

		for {
			line, _, err := readerB.ReadLine()
			if err != nil {
				break
			}
			dataA[string(line)] = 0
		}

		// debug
		fmt.Println("> Data in A: ", len(dataA), " | Data in B: ", len(dataB))

		// process and show data from A that not exists in B
		for keyA, _ := range dataA {
			_, exists := dataB[keyA]

			if !exists {
				fmt.Println(keyA)
			}
		}

	}
}

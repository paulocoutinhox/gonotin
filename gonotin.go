package main
import (
	"os"
	"log"
	"bufio"
	"fmt"
	"sort"
	"io"
)

func readMap(filename string) map[string]bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	m := make(map[string]bool)
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m[string(b)] = true
	}
	return m
}

func readMap2(filename string) map[string]struct{} {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	m := make(map[string]struct{})
	empty := struct{}{}
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		m[string(b)] = empty
	}
	return m
}

func main() {
	// select mode:
	// '1' - using array and sort
	// '2' - to use map strategy (mode 2 is too fast)
	// '3' - that is a poorimpl.
	// '4' - Uli Kunitz version from go-nuts
	// '5' - Justin version from go-nuts

	var mode = 5

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

	} else if mode == 4 {

		// read files
		mapB := readMap(os.Args[2])
		f, err := os.Open(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		// debug
		fmt.Println("> Data in B: ", len(mapB))

		// process and show data from A that not exists in B
		r := bufio.NewReader(f)
		for {
			b, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			l := string(b)
			if _, ok := mapB[l]; !ok {
				fmt.Println(l)
			}
		}

	} else if mode == 5 {

		// read files
		mapA := readMap(os.Args[1])
		mapB := readMap(os.Args[2])

		w := bufio.NewWriter(os.Stdout)

		// debug
		fmt.Println("> Data in A: ", len(mapA), " | Data in B: ", len(mapB))

		// process and show data from A that not exists in B
		for item := range mapA {
			if _, ok := mapB[item]; !ok {
				w.WriteString(item)
				w.WriteByte('\n')
			}
		}

		w.Flush()

	}
}

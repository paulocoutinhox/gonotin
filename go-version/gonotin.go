package main
import (
	"os"
	"log"
	"bufio"
	"fmt"
	"sort"
	"io"
	"strconv"
)

// Uli functions
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

// Justin functions
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

// Matt functions
func setError(errp *error, errf func() error) {
	if *errp == nil {
		*errp = errf()
	}
}

func filter(w io.Writer, r io.Reader, seen map[string]struct{}) (err error) {
	bw := bufio.NewWriter(w)
	defer setError(&err, bw.Flush)
	s := bufio.NewScanner(r)
	defer setError(&err, s.Err)
	for s.Scan() {
		if _, ok := seen[string(s.Bytes())]; !ok {
			bw.Write(s.Bytes())
			bw.WriteByte('\n')
		}
	}
	return nil
}

func load(r io.Reader, seen map[string]struct{}) (err error) {
	empty := struct{}{}
	s := bufio.NewScanner(r)
	defer setError(&err, s.Err)
	for s.Scan() {
		seen[string(s.Bytes())] = empty
	}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mustOpen(filename string) *os.File {
	if f, err := os.Open(filename); err != nil {
		log.Fatal(err)
		panic("notreached")
	} else {
		return f
	}
}

// Andrew functions
func reader(r io.Reader) func() (uint64, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return func() (uint64, error) {
		if !s.Scan() {
			if err := s.Err(); err != nil {
				return 0, err
			}
			return 0, io.EOF
		}
		return btou(s.Bytes())
	}
}

func btou(b []byte) (uint64, error) {
	var n uint64
	for _, v := range b {
		if v < '0' || v > '9' {
			return 0, strconv.ErrSyntax
		}
		if n >= ^uint64(0)/10+1 {
			return ^uint64(0), strconv.ErrRange
		}
		n *= 10
		n1 := n + uint64(v-'0')
		if n1 < n {
			return ^uint64(0), strconv.ErrRange
		}
		n = n1
	}
	return n, nil
}

func check2(err error) {
	if err != nil {
		die(err)
	}
}

func die(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

// Mode functions

func mode1() {
	// it is generating wrong results
	fmt.Println("it is generating wrong results")
	return

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

	// process and show data from A that not exists in B
	sort.Strings(dataB)

	for _, valueA := range dataA {
		if !(dataB[sort.SearchStrings(dataB, valueA)] == valueA) {
			fmt.Println(valueA)
		}
	}
}

func mode2() {
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

func mode3() {
	// it is generating wrong results
	fmt.Println("it is generating wrong results")
	return

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

	// process and show data from A that not exists in B
	for keyA, _ := range dataA {
		_, exists := dataB[keyA]

		if !exists {
			fmt.Println(keyA)
		}
	}
}

func mode4() {
	// it is generating wrong results
	fmt.Println("it is generating wrong results")
	return

	// read files
	mapB := readMap(os.Args[2])
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

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
}

func mode5() {
	// read files
	mapA := readMap2(os.Args[1])
	mapB := readMap2(os.Args[2])

	w := bufio.NewWriter(os.Stdout)

	// process and show data from A that not exists in B
	for item := range mapA {
		if _, ok := mapB[item]; !ok {
			w.WriteString(item)
			w.WriteByte('\n')
		}
	}

	w.Flush()
}

func mode6() {
	// it is generating wrong results
	fmt.Println("it is generating wrong results")
	return

	// read files and process
	seen := make(map[string]struct{})
	func() {
		r := mustOpen(os.Args[1])
		defer r.Close()
		check(load(r, seen))
	}()
	func() {
		r := mustOpen(os.Args[2])
		defer r.Close()
		check(filter(os.Stdout, r, seen))
	}()
}

func mode7() {
	f0, err := os.Open(os.Args[1])
	check2(err)
	defer f0.Close()

	f1, err := os.Open(os.Args[2])
	check2(err)
	defer f1.Close()

	m := make(map[uint64]struct{})
	r0 := reader(f1)
	for {
		n, err := r0()
		if err != nil {
			if err == io.EOF {
				break
			}
			die(err)
		}
		m[n] = struct{}{}
	}

	var b [20]byte
	r1 := reader(f0)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, err := r1()
		if err != nil {
			if err == io.EOF {
				return
			}
			die(err)
		}
		if _, ok := m[n]; ok {
			continue
		}
		m[n] = struct{}{}
		w.Write(strconv.AppendUint(b[:0], n, 10))
		w.WriteByte('\n')
	}
}

// Main function
func main() {
	// select mode:
	// '1' - using array and sort
	// '2' - to use map strategy
	// '3' - that is a poorimpl.
	// '4' - Uli Kunitz version from go-nuts
	// '5' - Justin version from go-nuts
	// '6' - Matt version from go-nuts
	// '7' - Andrew version from go-nuts

	var mode, _ = strconv.Atoi(os.Args[3])

	if mode == 1 {
		mode1()
	} else if mode == 2 {
		mode2()
	} else if mode == 3 {
		mode3()
	} else if mode == 4 {
		mode4()
	} else if mode == 5 {
		mode5()
	} else if mode == 6 {
		mode6()
	} else if mode == 7 {
		mode7()
	}
}

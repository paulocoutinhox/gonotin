package main
import (
	"os"
	"log"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"strings"
)

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func main() {
	// general data
	rand.Seed(int64(time.Now().Nanosecond()))
	filename := os.Args[1]

	totalNumbersArg, _ := strconv.ParseInt(os.Args[2], 10, 0)
	totalNumbers := int(totalNumbersArg)

	maxNumberArg, _ := strconv.ParseInt(os.Args[3], 10, 0)
	maxNumber := int(maxNumberArg)

	// generate data for file
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	// random number
	for x := 0; x < totalNumbers; x++ {
		var number = random(0, maxNumber)
		file.Write([]byte(strconv.Itoa(number)))
		file.WriteString("\n")
	}

	defer file.Close()

	fmt.Println("> File generated at: ", strings.Trim(filename, " "))
}

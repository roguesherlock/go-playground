package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

var filename = flag.String("input", "input.txt", "filename for the input")

type Range struct {
	Start int
	End   int
}

func (r Range) contains(other Range) bool {
	return r.Start <= other.Start && r.End >= other.End
}

func overlap(r1, r2 Range) bool {
	return (r1.End >= r2.Start && r1.End <= r2.End) || (r2.End >= r1.Start && r2.End <= r1.End)
}

func Split(line string) (r1, r2 Range) {
	ranges := strings.Split(line, ",")
	range1 := strings.Split(ranges[0], "-")
	range2 := strings.Split(ranges[1], "-")

	r1.Start, _ = strconv.Atoi(range1[0])
	r1.End, _ = strconv.Atoi(range1[1])
	r2.Start, _ = strconv.Atoi(range2[0])
	r2.End, _ = strconv.Atoi(range2[1])
	return r1, r2
}

func main() {

	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening file, %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// part 1
	var sum int
	var sum2 int
	for scanner.Scan() {
		line := scanner.Text()
		range1, range2 := Split(line)
		if range1.contains(range2) || range2.contains(range1) {
			sum += 1
		}
		if overlap(range1, range2) {
			sum2 += 1
		}
	}

	log.Println("Sum for part 1:", sum)
	log.Println("Sum for part 2:", sum2)

}

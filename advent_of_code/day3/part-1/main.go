package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var filename = flag.String("input", "input.txt", "filename for the input")

func main() {

	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening file, %v", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	priorityMap := make(map[rune]int)

	const ascii = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, s := range strings.ToLower(ascii) {
		priorityMap[s] = i + 1
	}
	for i, s := range ascii {
		priorityMap[s] = i + 27
	}

	// part 1
	var items []rune
	for scanner.Scan() {
		line := scanner.Text()
		c1 := line[0:(len(line) / 2)]
		c2 := line[(len(line) / 2):]
		for _, r := range c1 {
			if strings.Contains(c2, string(r)) {
				items = append(items, r)
				//log.Printf("Found: %c, c1: %s, c2: %s\n", rune(r), c1, c2)
				break
			}
		}
	}
	var sum int
	for _, r := range items {
		//log.Printf("rune: %c, value: %d, sum till now: %d\n", rune(r), priorityMap[rune(r)], sum)
		sum += priorityMap[r]
	}

	log.Println("Sum for part 1:", sum)

}

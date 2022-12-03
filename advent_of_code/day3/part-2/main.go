package main

import (
	"bufio"
	"flag"
	"io"
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

	reader := bufio.NewReader(input)

	priorityMap := make(map[rune]int)

	const ascii = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, s := range strings.ToLower(ascii) {
		priorityMap[s] = i + 1
	}
	for i, s := range ascii {
		priorityMap[s] = i + 27
	}

	// part 2
	var items []rune
	for {
		line1, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line2, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line3, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		//log.Printf("line1, line2, line3: %s, %s, %s", line1, line2, line3)
		for _, r := range line1 {
			if strings.Contains(line2, string(r)) && strings.Contains(line3, string(r)) {
				items = append(items, r)
				//log.Printf("Found: %c, line1: %s, line2: %s, line3: %s\n", rune(r), line1, line2, line3)
				break
			}
		}

	}
	sum := 0
	//log.Printf("items: %v\n", items)
	for _, r := range items {
		//log.Printf("rune: %c, value: %d, sum till now: %d\n", rune(r), priorityMap[rune(r)], sum)
		sum += priorityMap[r]
	}
	log.Println("Sum for part 2:", sum)

}

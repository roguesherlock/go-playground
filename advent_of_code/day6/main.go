package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var filename = flag.String("input", "input.txt", "filename for the input")

func isUnique(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	flag.Parse()

	// Read the file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Close the file when we're done
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var marker int
	var message int
	for scanner.Scan() {
		line := scanner.Text()

		// part 1
		for i := 0; i < len(line)-4; i++ {
			if isUnique(line[i : i+4]) {
				log.Printf("found unique %q at %d", line[i:i+4], i)
				marker = i + 4
				break
			}
		}

		for i := 0; i < len(line)-14; i++ {
			if isUnique(line[i : i+14]) {
				log.Printf("found unique message %q at %d", line[i:i+14], i)
				message = i + 14
				break
			}
		}
	}

	log.Println("Marker:", marker)
	log.Println("Message:", message)

}

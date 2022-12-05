package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var filename = flag.String("input", "input.txt", "filename for the input")

type Stack struct {
	crates []string
}

func (s *Stack) Push(crate string) {
	if len(crate) == 0 {
		return
	}
	s.crates = append(s.crates, crate)
}

func (s *Stack) Pop() (string, error) {
	if len(s.crates) == 0 {
		return "", fmt.Errorf("Stack is empty")
	}
	crate := s.crates[len(s.crates)-1]
	s.crates = s.crates[:len(s.crates)-1]
	return (crate), nil
}

// MoveTo moves number of items to a given stack. It keeps the order of the items.
func (s *Stack) MoveTo(stack *Stack, numberOfCrates int) {
	crates := make([]string, numberOfCrates)
	for i := 0; i < numberOfCrates; i++ {
		crate, err := s.Pop()
		if err != nil {
			break
		}
		crates = append(crates, crate)
	}
	for i := len(crates) - 1; i >= 0; i-- {
		stack.Push(crates[i])
	}
}

func (s *Stack) Reverse() {
	runes := s.crates
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

func createStacks(line []byte, drawing []string) []Stack {
	numberOfStacks := 9

	stacks := make([]Stack, numberOfStacks)
	numberRegex := regexp.MustCompile(`[0-9]+`)
	asciiRegex := regexp.MustCompile(`[a-zA-Z]+`)

	for i := 0; i < len(drawing); i++ {
		for j := 0; j < len(line); j++ {
			if numberRegex.MatchString(string(line[j])) && len(drawing[i]) >= j && asciiRegex.MatchString(string(drawing[i][j])) {
				stackIndex, _ := strconv.Atoi(string(line[j]))
				stacks[stackIndex-1].Push(string(drawing[i][j]))
			}
		}
	}

	for _, stack := range stacks {
		stack.Reverse()
	}

	return stacks
}

func updateStacks(stacks []Stack, line []byte) {
	operationRegex := regexp.MustCompile(`^move ([0-9]+) from ([0-9]+) to ([0-9]+)$`)

	operations := operationRegex.FindStringSubmatch(string(line))
	if len(operations) != 4 {
		return
	}

	numberOfCrates, _ := strconv.Atoi(operations[1])
	from, _ := strconv.Atoi(operations[2])
	to, _ := strconv.Atoi(operations[3])
	from -= 1
	to -= 1
	if (from < 0 || from >= len(stacks)) || (to < 0 || to >= len(stacks)) {
		return
	}
	if numberOfCrates == 1 {
		crate, _ := stacks[from].Pop()
		stacks[to].Push(crate)
	} else {
		stacks[from].MoveTo(&stacks[to], numberOfCrates)
	}

}

func main() {

	flag.Parse()

	input, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Error opening file, %v", err)
	}
	defer input.Close()

	reader := bufio.NewReader(input)

	// part 2
	var drawing []string
	var stacks []Stack
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if strings.Contains(string(line), "move") {
			updateStacks(stacks, line)
		} else if len(line) >= 1 && strings.Contains(string(line), "1") {
			stacks = createStacks(line, drawing)
		} else if len(line) >= 1 {
			drawing = append(drawing, string(line))
		}
	}

	str := ""
	for i, _ := range stacks {
		v, _ := stacks[i].Pop()
		str += fmt.Sprintf("%v", v)

	}

	log.Println("Part 2:", str)

}

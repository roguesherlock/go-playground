package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
)

var filename = flag.String("input", "input.txt", "filename for the input")

func main() {
	flag.Parse()

	// Read the file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]int

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i], _ = strconv.Atoi(string(c))
		}
		grid = append(grid, row)
	}

	visibleTrees := len(grid)*2 + len(grid[0])*2 - 4

	isVisible := func(grid [][]int, i, j int) bool {
		// check left
		visible := true
		for k := 0; k < j; k++ {
			if grid[i][k] >= grid[i][j] {
				visible = false
			}
		}
		if visible {
			return visible
		}

		// check right
		visible = true
		for k := j + 1; k < len(grid[i]); k++ {
			if grid[i][k] >= grid[i][j] {
				visible = false
			}
		}
		if visible {
			return visible
		}

		// check up
		visible = true
		for k := 0; k < i; k++ {
			if grid[k][j] >= grid[i][j] {
				visible = false
			}
		}
		if visible {
			return visible
		}

		// check down
		visible = true
		for k := i + 1; k < len(grid); k++ {
			if grid[k][j] >= grid[i][j] {
				visible = false
			}
		}
		return visible
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if isVisible(grid, i, j) {
				visibleTrees++
			}
		}
	}
	log.Println("visible trees:", visibleTrees)
}

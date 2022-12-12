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

	calculateScenicScore := func(grid [][]int, i, j int) int {
		// look left
		leftScore := 0
		for k := j - 1; k >= 0; k-- {
			leftScore++
			if grid[i][k] >= grid[i][j] {
				break
			}
		}

		// look right
		rightScore := 0
		for k := j + 1; k < len(grid[i]); k++ {
			rightScore++
			if grid[i][k] >= grid[i][j] {
				break
			}
		}

		// look up
		upScore := 0
		for k := i - 1; k >= 0; k-- {
			upScore++
			if grid[k][j] >= grid[i][j] {
				break
			}
		}

		// look down
		downScore := 0
		for k := i + 1; k < len(grid); k++ {
			downScore++
			if grid[k][j] >= grid[i][j] {
				break
			}
		}
		return leftScore * rightScore * upScore * downScore
	}

	maxScenicScore := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if isVisible(grid, i, j) {
				visibleTrees++
			}
			score := calculateScenicScore(grid, i, j)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	log.Println("visible trees:", visibleTrees)
	log.Println("highest scenic score:", maxScenicScore)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	filePath := os.Args[1]

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var reader = bufio.NewReader(f)
	var sums []int = make([]int, 0)
	var sum int

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		var num int
		i, _ := fmt.Sscanf(string(line), "%d", &num)
		if i == 1 {
			sum += num
		} else {
			sums = append(sums, sum)
			sum = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	fmt.Println("Part 1:", sums[0])

	sum = 0
	for i := 0; i < 3; i++ {
		sum += sums[i]
	}

	fmt.Println("Part 2:", sum)

}

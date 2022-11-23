package quiz2

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var filename = flag.String("filename", "problems.csv", "a csv file in the format of 'question,answer'")
var limit = flag.Int("limit", 30, "the time limit for the quiz in seconds")
var debug = flag.Bool("debug", false, "enable debug mode")

func main() {

	flag.Parse()

	if *debug {
		fmt.Println("You chose the file:", *filename)

		fmt.Println("You chose the time limit:", *limit)
	}

	var file, err = os.OpenFile(*filename, os.O_RDONLY, 0644)
	if err != nil {
		panic("Error opening file")
	}
	defer file.Close()

	var csvReader = csv.NewReader(file)

	score := 0
	totalQuestions := 0
	correct := 0

	fmt.Println("Welcome to the quiz!", *filename, *limit)

	for {
		var timeout = make(chan bool, 1)
		var result = make(chan string, 1)
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		totalQuestions++

		fmt.Println(record[0])
		go func() {
			time.Sleep(time.Duration(*limit) * time.Second)
			timeout <- true
		}()

		go func() {
			var answer string
			fmt.Scanln(&answer)
			result <- answer
		}()

		select {
		case <-timeout:
			fmt.Println("You ran out of time!")
			fmt.Printf("You scored %v out of %v for a percentage of %.2f%%\n", score, totalQuestions, float64(correct)/float64(totalQuestions)*100)
			return
		case answer := <-result:
			if answer == record[1] {
				score++
				correct++
			}
		}
	}

	fmt.Printf("You scored %v out of %v for a percentage of %.2f%%\n", score, totalQuestions, float64(correct)/float64(totalQuestions)*100)
	return

}

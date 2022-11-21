package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Comic struct {
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
	Alt        string `json:"alt"`
}

var filename = flag.String("filename", "xkcd.json", "a json file in the format to save results to")

func main() {

	flag.Parse()
	fmt.Println(flag.Args())

	file, err := os.Open(*filename)
	if err != nil {
		panic(fmt.Errorf("Couldn't opoen file %v:\n %v ", *filename, err))
	}

	comics := make([]Comic, 0)
	if err := json.NewDecoder(file).Decode(&comics); err != nil {
		panic(fmt.Errorf("Couldn't decode file %v:\n %v ", *filename, err))
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No search terms provided")
		return
	}

	for _, comic := range comics {
		for _, arg := range args {
			alreadyMatched := false
			if strings.Contains(comic.Title, arg) && !alreadyMatched {
				alreadyMatched = true
				fmt.Printf("%d: %s\n", comic.Num, comic.Title)
			}
		}

	}

}

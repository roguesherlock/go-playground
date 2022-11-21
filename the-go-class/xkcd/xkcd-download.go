package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

const URL = `https://xkcd.com/${n}/info.0.json`

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

func getComic(number int) (Comic, error) {
	comic := Comic{}
	resp, err := http.Get(strings.Replace(URL, "${n}", fmt.Sprint(number), 1))
	if err != nil {
		return comic, err
	}
	if resp.StatusCode == 404 {
		return comic, fmt.Errorf("Comic %d not found", number)
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	resp.Body.Close()
	return comic, nil
}

func main() {

	flag.Parse()
	var wg sync.WaitGroup
	var bufferChannel = make(chan int, 10)

	// initial buffer
	for i := 0; i < 10; i++ {
		bufferChannel <- i
	}

	comics := make([]Comic, 0)

	for i := 1; i < 2702; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// release
			defer func() { bufferChannel <- i }()

			// wait
			<-bufferChannel
			comic, err := getComic(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			comics = append(comics, comic)
		}(i)
	}

	wg.Wait()

	fmt.Println("Read", len(comics), "comics")

	file, err := os.Create(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(comics)

	// err := os.WriteFile(*filename, []byte(fmt.Sprintf("%v", comics)), 0644)
	// if err != nil {
	// 	fmt.Printf("Error while writing to file: %v", err)
	// }
}

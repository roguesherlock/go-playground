package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Starting server at :8080")

}

type dollars float32

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v", db)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	value := req.URL.Query().Get("value")
	if _, ok := db[item]; ok {
		fmt.Fprintf(w, "Couldn't create item. It already exists.")
	} else {
		if v, err := strconv.ParseFloat(value, 32); err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Fprintf(w, "Value isn't a number, please try again.\n Item: %v\n Value: %v", item, value)
		} else {
			fmt.Printf("Error: %v\n", err)
			db[item] = dollars(v)
			fmt.Fprintf(w, "Created item with value.\n Item: %v\n Value: %v", item, value)
		}
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	value := req.URL.Query().Get("value")
	if _, ok := db[item]; ok {
		if v, err := strconv.ParseFloat(value, 32); err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Fprintf(w, "Value isn't a number, please try again.\n Item: %v\n Value: %v", item, value)
		} else {
			db[item] = dollars(v)
			fmt.Fprintf(w, "Updated item.\n Item: %v\n Value: %v", item, value)
		}
	} else {
		fmt.Fprintf(w, "Couldn't find item: %v", item)
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprint(w, "Yo it's removed")
	} else {
		fmt.Fprintf(w, "Couldn't find item: %v", item)
	}
}

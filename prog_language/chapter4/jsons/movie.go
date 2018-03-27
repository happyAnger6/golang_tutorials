package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		Movie{"hello", 1997, false, []string{"a", "b"}},
		Movie{"the", 1998, true, []string{"c", "d"}},
		Movie{"world", 1999, false, []string{"e", "f"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", movies)
	fmt.Printf("%s\n", data)

	data1, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data1)
}

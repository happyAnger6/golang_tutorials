package main

import (
	"net/http"
	"fmt"
)

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db{
		fmt.Fprint(w, "%s: %s\n", k, v)
	}
}

func main() {
	db := database{"a": 10}
	mux := http.ServeMux{}
	mux.Handle("/", http.HandlerFunc(db.list))
}

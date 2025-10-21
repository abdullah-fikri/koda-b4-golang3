package main

import (
	"fmt"
	"golang3/search"
)


func main() {
	names := []string{"fiki", "yoga", "anggi", "itsna", "federus", "ari", "sidik"}
	var keyword string
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&keyword)

	search.SearchPerson(names, keyword)
}

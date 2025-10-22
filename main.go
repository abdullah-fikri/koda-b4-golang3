package main

import (
	"bufio"
	"fmt"
	"golang3/search"
	"os"
	"strings"
)

func main() {

	names := []string{"fiki", "yoga", "anggi", "itsna", "federus", "ari", "sidik"}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")
	keyword,_:= reader.ReadString('\n') 
	keyword = strings.TrimSpace(keyword)  

	search.SearchPerson(names, &keyword)
}

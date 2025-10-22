package main

import (
	"bufio"
	"fmt"
	"golang3/search"
	"os"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("input salah:", r)
			os.Exit(1)
		}
	}()

	names := []string{"fiki", "yoga", "anggi", "itsna", "federus", "ari", "sidik"}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama: ")
	keyword, _ := reader.ReadString('\n') 
	keyword = strings.TrimSpace(keyword)  

	search.SearchPerson(names, &keyword)
}

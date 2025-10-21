package main

import "fmt"

func main(){
	name := []string{"fiki", "yoga", "anggi", "itsna", "federus", "ari", "sidik"}
	var keywoord string
	found := false
	fmt.Println("masukkan nama ")
	fmt.Scan(&keywoord)
	for i := range name{
		if keywoord == name[i]{
			found = true
			fmt.Println("nama ditemukan: ", keywoord)
		}
	}
	if !found {
		fmt.Println("nama tidak ditemukan: ", keywoord)
	}
}

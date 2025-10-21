package search

import (
	"fmt"
)

func SearchPerson(names []string, keyword string) {
	found := false
	for i := range names {
		if keyword == names[i] {
			found = true
			fmt.Println("nama ditemukan:", keyword)
		}
	}
	if !found {
		fmt.Println("nama tidak ditemukan:", keyword)
	}
}


package search

import (
	"fmt"
)

func SearchPerson(names []string, keyword string) {
	found := false
	for i := range names {
		if keyword == names[i] {
			var yes []string
			yes = append(yes, keyword)
			found = true
			fmt.Println("nama ditemukan:", yes)
		}
	}
	var notFound []string
	if !found {
		fmt.Println("nama tidak ditemukan:", notFound)
	}
}


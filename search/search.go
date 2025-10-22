package search

import (
	"fmt"
	"strings"
)

func SearchPerson(names []string, keyword string) {
	if strings.TrimSpace(keyword) == "" {
		panic("input kosong!")
	}

	found := false
	var yes []string
	for i := range names {
		if keyword == names[i] {
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

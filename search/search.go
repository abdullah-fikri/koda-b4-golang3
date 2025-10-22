package search

import (
	"fmt"
	"os"
	"strings"
)

func SearchPerson(names []string, keyword *string)[]string {
	if strings.TrimSpace(*keyword) == "" {
		panic("input kosong!")
	}

	found := false
	var yes []string
	for i := range names {
		if *keyword == names[i] {
			yes = append(yes, *keyword)
			found = true
			fmt.Println("nama ditemukan:", yes)
		}
	}

	
	if !found {
		fmt.Println("nama tidak ditemukan:", []string{})
		defer os.Exit(0)
	}
	return []string{}
}

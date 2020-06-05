package main

import "fmt"

type Developer struct {
	Name string
	Age  int
}

func ExistName(names []string, name string) bool {
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}

func FilterUnique(developers []Developer) []string {
	// TODO Implement
	var names []string
	for i := 0; i < len(developers); i++ {
		if !ExistName(names, developers[i].Name) {
			names = append(names, developers[i].Name)
		}
	}
	return names
}

func main() {
	fmt.Println("Filter Unique Challenge")
	developers := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Alan"},
		Developer{Name: "Jennifer"},
		Developer{Name: "Graham"},
		Developer{Name: "Paul"},
		Developer{Name: "Alan"},
	}
	unique_name_developers := FilterUnique(developers)
	fmt.Println(unique_name_developers)
}

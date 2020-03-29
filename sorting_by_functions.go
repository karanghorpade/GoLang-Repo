package main

import (
	"fmt"
	"sort"
)

type user struct {
	name        string
	designation string
}

type byName []user

func (n byName) Len() int {
	return len(n)
}
func (n byName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n byName) Less(i, j int) bool {
	return n[i].name < n[j].name
}

type byDesignation []user

func (d byDesignation) Len() int {
	return len(d)
}
func (d byDesignation) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
func (d byDesignation) Less(i, j int) bool {
	return d[i].designation < d[j].designation
}

func main() {
	u1 := user{

		name:        "James",
		designation: "Software Engineer",
	}

	u2 := user{

		name:        "Cristiano",
		designation: "Network Engineer",
	}

	users := []user{u1, u2}

	sort.Sort(byName(users))

	for i, u := range users {
		fmt.Println(i, u.name)
	}

	sort.Sort(byDesignation(users))

	for i, u := range users {
		fmt.Println(i, u.designation)
	}

}
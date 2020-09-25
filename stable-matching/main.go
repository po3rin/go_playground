package main

import "fmt"

type person struct {
	name              string
	orderOfPreference []string
	rejectedList      map[string]struct{}
	matched           string
}

func main() {
	mList := []person{
		{
			name:              "mA",
			orderOfPreference: []string{"fA", "fB", "fC", "fD"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "mB",
			orderOfPreference: []string{"fC", "fB", "fA", "fD"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "mC",
			orderOfPreference: []string{"fA", "fB", "fD", "fC"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "mD",
			orderOfPreference: []string{"fA", "fB", "fD", "fC"},
			rejectedList:      make(map[string]struct{}, 4),
		},
	}
	fList := []person{
		{
			name:              "fA",
			orderOfPreference: []string{"mA", "mB", "mC", "mD"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "fB",
			orderOfPreference: []string{"mB", "mA", "mD", "mC"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "fC",
			orderOfPreference: []string{"mB", "mC", "mA", "mD"},
			rejectedList:      make(map[string]struct{}, 4),
		},
		{
			name:              "fD",
			orderOfPreference: []string{"mA", "mD", "mC", "mB"},
			rejectedList:      make(map[string]struct{}, 4),
		},
	}

	matching := stableMatching(mList, fList)
	fmt.Println(matching)
}

func matchedList(mList []person) map[string]string {
	matching := make(map[string]string, len(mList))
	for _, m := range mList {
		matching[m.name] = m.matched
	}
	return matching
}

func noMatched(mList []person) (person, bool) {
	for _, m := range mList {
		if m.matched == "" {
			return m, true
		}
	}
	return person{}, false
}

// TODO: imp
func stableMatching(mList []person, fList []person) map[string]string {
	var matchNum int
	for matchNum == len(mList) {

	}

	return matchedList(mList)
}

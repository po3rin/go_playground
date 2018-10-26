package main

import (
	"fmt"
	"go-playground/eure_tech/shuffle"
	"os"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type Dice []int

func (d Dice) Seed() int64   { return int64(os.Getpid()) }
func (d Dice) Len() int      { return len(d) }
func (d Dice) Swap(i, j int) { d[i], d[j] = d[j], d[i] }

func main() {
	// fase1 --------------------

	people := []Person{
		{"Bob", 31}, {"John", 42}, {"Michael", 17}, {"Jenny", 26},
	}
	fmt.Println(people)
	// ascending sort.Sort(ByAge(people)) fmt.Println(people)
	// descending
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	// fase2 --------------------

	people = []Person{
		{"Bob", 31}, {"John", 42}, {"Michael", 17}, {"Jenny", 26},
	}
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// fase3 --------------------
	var dice = Dice([]int{1, 2, 3, 4, 5, 6})
	fmt.Printf("%v\n", dice)
	shuffle.Shuffle(dice)
	fmt.Printf("%v\n", dice)
}

package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func md5hash(str string) string {
	// string型を[]byte型の変更して使う
	md5 := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", md5)
}

func heapPermutation(mystr []string, n int, permutations *[][]string) {
	if n == 1 {
		*permutations = append(*permutations, append([]string{}, mystr...))
		return
	}
	for i := 0; i < n; i++ { //calculate the permutation using heap
		heapPermutation(mystr, n-1, permutations)
		if n%2 == 0 {
			mystr[i], mystr[n-1] = mystr[n-1], mystr[i]
		} else {
			mystr[0], mystr[n-1] = mystr[n-1], mystr[0]
		}
	}
}

func splitN(msg string, n int) []string {
	runes := []rune(msg)
	result := []string{}
	for i := 0; i < len(runes); i += n {
		if i+n < len(runes) {
			result = append(result, string(runes[i:(i+n)]))
		} else {
			result = append(result, string(runes[i:]))
		}
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	candidate := "5p8s1s7s4m7s1z3z2s2z9m4m5z6z8m"
	candidateSplit := splitN(candidate, 2)

	fmt.Println("------candidate------")
	fmt.Println(candidateSplit)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(candidateSplit), func(i, j int) {
		candidateSplit[i], candidateSplit[j] = candidateSplit[j], candidateSplit[i]
	})

	answer := strings.Join(candidateSplit, "")
	answerMd5 := md5hash(answer)
	fmt.Println("------want------")
	fmt.Println(answer)
	fmt.Println("------want md5------")
	fmt.Println(answerMd5)

	permutations := &[][]string{}
	heapPermutation(candidateSplit, 3, permutations)

	for _, p := range *permutations {
		s := strings.Join(p, "")
		if answerMd5 == md5hash(s) {
			fmt.Println("------got------")
			fmt.Println(s)
		}
	}
}

package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var hai = []string{
	"1m", "2m", "3m", "4m", "5m", "6m", "7m", "8m", "9m",
	"1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p",
	"1s", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s",
	"1z", "2z", "3z", "4z", "5z", "6z", "7z",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func yama() []string {
	yama := hai
	yama = append(yama, hai...)
	yama = append(yama, hai...)
	yama = append(yama, hai...)
	rand.Shuffle(len(yama), func(i, j int) {
		yama[i], yama[j] = yama[j], yama[i]
	})
	return yama
}

func md5hash(str string) string {
	// string型を[]byte型の変更して使う
	md5 := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", md5)
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

func checkJansoul() {
	s := "1p4p6z8p1s9m6z4z7s7p2z9s1s1z4p5z9s1m7p5s8p1m2z6s4z3z1z1m2s4p6z9s5s2z9m3p7s9s2p5z6p9m1p1s4s2p4z4s6z3p4s8s2s1z1p9p7s3p5p5s7z3s9p3z8s1p4z6p"
	fmt.Println(md5hash(s))
}

func main() {
	// checkJansoul()

	candidate := "5p1z1s7s4m7s1z3z2s2z4m"
	candidateSplit := splitN(candidate, 2)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(candidateSplit), func(i, j int) {
		candidateSplit[i], candidateSplit[j] = candidateSplit[j], candidateSplit[i]
	})

	answer := strings.Join(candidateSplit, "")
	answerMd5 := md5hash(answer)

	fmt.Println("------answer------")
	fmt.Println(answer)
	fmt.Println(len(candidateSplit))
	fmt.Println(answerMd5)

	rand.Shuffle(len(candidateSplit), func(i, j int) {
		candidateSplit[i], candidateSplit[j] = candidateSplit[j], candidateSplit[i]
	})

	permutations := backtrackPermutate(candidateSplit)
	fmt.Println("------permutations info------")
	fmt.Printf("len: %v\n", len(permutations))

	rand.Shuffle(len(candidateSplit), func(i, j int) {
		candidateSplit[i], candidateSplit[j] = candidateSplit[j], candidateSplit[i]
	})

	for _, p := range permutations {
		s := strings.Join(p, "")
		hash := md5hash(s)
		if answerMd5 == hash {
			fmt.Println("------got------")
			fmt.Println(s)
			fmt.Println(hash)
			break
		}
	}
}

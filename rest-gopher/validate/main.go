// method によって処理を差し込む例
package main

import "fmt"

func main() {
	userScore := 100
	c := NewTotalCalcScore()
	fmt.Println(c(userScore))                // 100
	fmt.Println(c.WithValidation(userScore)) // 100
	userScore = 0
	fmt.Println(c.WithValidation(userScore)) // panic: score = 0
}

type CalcTotalScore func(userScore int) (totalScore int)

func NewTotalCalcScore() CalcTotalScore {
	f := func(userScore int) (totalScore int) {
		// ここでバリデーション計算ロジック処理はいる想定
		return userScore
	}
	return CalcTotalScore(f)
}
func (c CalcTotalScore) WithValidation(userScore int) (totalScore int) {
	if userScore == 0 {
		panic("score = 0")
	}
	return c(userScore)
}

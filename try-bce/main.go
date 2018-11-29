package main // import "try-bce"

// sample1

func f1(s []int) {
	_ = s[0] // line 5: bounds check
	_ = s[1] // line 6: bounds check
	_ = s[2] // line 7: bounds check
}

func f2(s []int) {
	_ = s[2] // line 11: bounds check
	_ = s[1] // line 12: bounds check eliminated!
	_ = s[0] // line 13: bounds check eliminated!
}

func f3(s []int, index int) {
	_ = s[index] // line 17: bounds check
	_ = s[index] // line 18: bounds check eliminated!
}

func f4(a [5]int) {
	_ = a[4] // line 22: bounds check eliminated!
}

// sample2

func f5(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f6(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

func f7(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		_ = s[i]
		_ = s[i:len(s)]
	}
}

func f8(s []int, index int) {
	if index >= 0 && index < len(s) {
		_ = s[index]
		_ = s[index:len(s)]
	}
}

func f9(s []int) {
	if len(s) > 2 {
		_, _, _ = s[0], s[1], s[2]
	}
}

func main() {}

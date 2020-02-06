package replacer_test

import (
	"strings"
	"testing"
)

func replace() string {
	s := "abcdefghijklmnopqrstuvwxyz"

	s = strings.Replace(s, "abc", "あいう", -1)
	s = strings.Replace(s, "def", "えおか", -1)
	s = strings.Replace(s, "ghi", "きくけ", -1)
	s = strings.Replace(s, "jkl", "こさし", -1)
	s = strings.Replace(s, "mno", "すせそ", -1)
	s = strings.Replace(s, "pqr", "たちつ", -1)
	s = strings.Replace(s, "stu", "てとな", -1)
	s = strings.Replace(s, "vwx", "にぬね", -1)

	return s
}

func replacer() string {
	s := "abcdefghijklmnopqrstuvwxyz"

	return strings.NewReplacer(
		"abc", "あいう",
		"def", "えおか",
		"ghi", "きくけ",
		"jkl", "こさし",
		"mno", "すせそ",
		"pqr", "たちつ",
		"stu", "てとな",
		"vwx", "にぬね",
	).Replace(s)
}

// func BenchmarkReplace(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := replace()
// 		if s != "あいうえおかきくけこさしすせそたちつてとなにぬねyz" {
// 			fmt.Println("unexpected1")
// 		}
// 	}
// }

// func BenchmarkReplacer(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := replacer()
// 		if s != "あいうえおかきくけこさしすせそたちつてとなにぬねyz" {
// 			fmt.Println("unexpected2")
// 		}
// 	}
// }

func BenchmarkReplace(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyzaabbccddeeffgghhiijjkkllmmnnooppqqiissttuuvvwwxxyyzz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = strings.Replace(s, "abc", "あいう", -1)
		s = strings.Replace(s, "def", "えおか", -1)
		s = strings.Replace(s, "ghi", "きくけ", -1)
		s = strings.Replace(s, "jkl", "こさし", -1)
		s = strings.Replace(s, "mno", "すせそ", -1)
		s = strings.Replace(s, "pqr", "たちつ", -1)
		s = strings.Replace(s, "stu", "てとな", -1)
		s = strings.Replace(s, "vwx", "にぬね", -1)

		s = strings.Replace(s, "aabb", "あいう", -1)
		s = strings.Replace(s, "ccdd", "えおか", -1)
		s = strings.Replace(s, "eeff", "きくけ", -1)
		s = strings.Replace(s, "gghh", "こさし", -1)
		s = strings.Replace(s, "iijj", "すせそ", -1)
		s = strings.Replace(s, "kkll", "たちつ", -1)
		s = strings.Replace(s, "mmnn", "てとな", -1)
		s = strings.Replace(s, "oopp", "にぬね", -1)

		_ = s
	}
}

func BenchmarkReplacer(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyzaabbccddeeffgghhiijjkkllmmnnooppqqiissttuuvvwwxxyyzz"
	r := strings.NewReplacer(
		"abc", "あいう",
		"def", "えおか",
		"ghi", "きくけ",
		"jkl", "こさし",
		"mno", "すせそ",
		"pqr", "たちつ",
		"stu", "てとな",
		"vwx", "にぬね",
		"aabb", "あいう",
		"ccdd", "えおか",
		"eeff", "きくけ",
		"gghh", "こさし",
		"iijj", "すせそ",
		"kkll", "たちつ",
		"mmnn", "てとな",
		"oopp", "にぬね",
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = r.Replace(s)
		_ = s
	}
}

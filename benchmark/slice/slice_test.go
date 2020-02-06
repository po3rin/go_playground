package slice_test

import "testing"

type tag struct {
	id    uint64
	name  string
	query string
	appid int
}

type tagSpare struct {
	id    uint64
	name  string
	query string
	appid int
}

func prepare() []tag {
	tags := make([]tag, 100)
	for i := 0; i < 100; i++ {
		tags[i] = tag{
			id:    111,
			name:  "benchmark",
			query: "benchmark",
			appid: 222,
		}
	}
	return tags
}

func BenchmarkAppend(b *testing.B) {
	tags := prepare()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := make([]tagSpare, 0)
		for _, t := range tags {
			result = append(result, tagSpare{
				id:    t.id,
				name:  t.name,
				query: t.query,
				appid: t.appid,
			})
		}
	}
}

func BenchmarkAssign(b *testing.B) {
	tags := prepare()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := make([]tagSpare, len(tags))
		for i, t := range tags {
			result[i] = tagSpare{
				id:    t.id,
				name:  t.name,
				query: t.query,
				appid: t.appid,
			}
		}
	}
}

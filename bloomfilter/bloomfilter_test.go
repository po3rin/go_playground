package bloomfilter_test

import (
	"testing"

	"github.com/po3rin/bloomfilter"
)

func TestExists(t *testing.T) {
	tests := []struct {
		name       string
		s          string
		substrList []string
		size       int
		exist      bool
	}{
		{
			name: "non exit test",
			s:    "rabinkarp",
			substrList: []string{
				"bloom",
				"filter",
			},
			size:  1000,
			exist: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			filter, err := bloomfilter.NewBloomFilter(tt.size, len(tt.substrList))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			for _, s := range tt.substrList {
				filter.Add(s)
			}

			if got := filter.MightContain(tt.s); got != tt.exist {
				t.Errorf("unexpected result. want: %v, got: %v", tt.exist, got)
			}
		})
	}
}

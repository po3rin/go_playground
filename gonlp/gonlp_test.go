package gonlp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProcess(t *testing.T) {
	tests := []struct {
		input string
		want  Corpus
	}{
		{
			"hello po3rin. goodbye po3rin. hello golang.",
			Corpus{0, 1, 2, 3, 1, 2, 0, 4, 2},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _, _ := NewProcess(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("failed for %v ...", tt.input)
			}
		})
	}
}

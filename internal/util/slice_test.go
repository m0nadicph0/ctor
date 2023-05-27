package util

import (
	"reflect"
	"testing"
)

func TestSplitArgs(t *testing.T) {

	tests := []struct {
		name  string
		args  []string
		first []string
		rest  []string
	}{
		{
			name:  "Case 1",
			args:  []string{"migrate", "--", "up"},
			first: []string{"migrate"},
			rest:  []string{"up"},
		},
		{
			name:  "Case 2",
			args:  []string{"migrate", "--"},
			first: []string{"migrate"},
			rest:  []string{},
		},
		{
			name:  "Case 3",
			args:  []string{"build", "clean"},
			first: []string{"build", "clean"},
			rest:  []string{},
		},
		{
			name:  "Case 4",
			args:  []string{"build", "clean", "--", "binaries", "--", "objects"},
			first: []string{"build", "clean"},
			rest:  []string{"binaries", "--", "objects"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SplitArgs(tt.args, "--")
			if !reflect.DeepEqual(got, tt.first) {
				t.Errorf("SplitArgs() got = %v, want %v", got, tt.first)
			}
			if !reflect.DeepEqual(got1, tt.rest) {
				t.Errorf("SplitArgs() got1 = %v, want %v", got1, tt.rest)
			}
		})
	}
}

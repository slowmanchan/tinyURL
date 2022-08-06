package encoder

import "testing"

func Test_Decode(t *testing.T) {
	type test struct {
		name string
		url  string
		want int
	}

	tests := []test{
		{"test 1", "cb", 125},
		{"test 2", "e9a", 19158},
		{"test 3", "emjc", 1000000},
		{"test 4", "bfP3Qq", 1000000000},
		{"test 5", "rLHWfKe", 1000000000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Decode(tt.url)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

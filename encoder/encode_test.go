package encoder

import "testing"

func Test_Encode(t *testing.T) {
	type test struct {
		name string
		num  int
		want string
	}

	tests := []test{
		{"test 1", 125, "cb"},
		{"test 2", 19158, "e9a"},
		{"test 3", 1000000, "emjc"},
		{"test 4", 1000000000, "bfP3Qq"},
		{"test 5", 1000000000000, "rLHWfKe"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Encode(tt.num)
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

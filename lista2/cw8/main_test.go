package main

import (
	"slices"
	"testing"
)

func TestGenerateBitSlice(t *testing.T) {
	tests := []struct {
		num    int
		length int
		want   []byte
	}{
		{0, 1, []byte{0}},
		{1, 1, []byte{1}},
		{2, 2, []byte{0, 1}},
		{3, 2, []byte{1, 1}},
		{4, 3, []byte{0, 0, 1}},
		{5, 3, []byte{1, 0, 1}},
		{6, 3, []byte{0, 1, 1}},
		{7, 3, []byte{1, 1, 1}},
		{8, 4, []byte{0, 0, 0, 1}},
		{9, 4, []byte{1, 0, 0, 1}},
		{10, 4, []byte{0, 1, 0, 1}},
		{11, 4, []byte{1, 1, 0, 1}},
		{12, 4, []byte{0, 0, 1, 1}},
		{13, 4, []byte{1, 0, 1, 1}},
		{14, 4, []byte{0, 1, 1, 1}},
		{15, 4, []byte{1, 1, 1, 1}},
		{16, 5, []byte{0, 0, 0, 0, 1}},
		{17, 5, []byte{1, 0, 0, 0, 1}},
		{18, 5, []byte{0, 1, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := generateBitSlice(tt.num, tt.length); !slices.Equal[[]byte](got, tt.want) {
				t.Errorf("generateBitSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateLCSPairs(t *testing.T) {
	actual := generateLCSPairs(2)
	want := map[[2][2]byte]struct{}{
		{{0, 0}, {0, 0}}: {},
		{{0, 0}, {0, 1}}: {},
		{{0, 0}, {1, 0}}: {},
		{{0, 0}, {1, 1}}: {},
		{{0, 1}, {0, 0}}: {},
		{{0, 1}, {0, 1}}: {},
		{{0, 1}, {1, 0}}: {},
		{{0, 1}, {1, 1}}: {},
		{{1, 0}, {0, 0}}: {},
		{{1, 0}, {0, 1}}: {},
		{{1, 0}, {1, 0}}: {},
		{{1, 0}, {1, 1}}: {},
		{{1, 1}, {0, 0}}: {},
		{{1, 1}, {0, 1}}: {},
		{{1, 1}, {1, 0}}: {},
		{{1, 1}, {1, 1}}: {},
	}

	for pair := range actual {
		temp := [2][2]byte{{pair[0][0], pair[0][1]}, {pair[1][0], pair[1][1]}}
		if _, ok := want[temp]; !ok {
			t.Errorf("generateLCSPairs() = %v, want %v", actual, want)
		}
		delete(want, temp)
	}

	if len(want) != 0 {
		t.Errorf("Some sequences pair were not generated: %v", want)
	}
}

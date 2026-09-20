package matching

import (
	"testing"
)

func TestNaive(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name     string
		input    args
		expected int
	}{
		{name: "漢字一文字", input: args{"烈", "烈"}, expected: 0},
		{name: "ひらがな単一", input: args{"あいうえおえういあ", "え"}, expected: 3},
		{name: "数字", input: args{"012345678910", "1"}, expected: 1},
		{name: "空白", input: args{"Text matching test", " "}, expected: 4},
		{name: "見つからない", input: args{"abcdef", "z"}, expected: -1},
		{name: "空針", input: args{"abcdef", ""}, expected: 0},
		{name: "空本体", input: args{"", "a"}, expected: -1},
		{name: "複数ルーン針", input: args{"あいうえお", "うえ"}, expected: 2},
		{name: "重なり", input: args{"aaa", "aa"}, expected: 0},
		{name: "Runeインデックス", input: args{"あいう", "い"}, expected: 1},
		{name: "末尾一致", input: args{"abcdef", "ef"}, expected: 4},
		{name: "部分一致後に一致", input: args{"aaaaab", "aaab"}, expected: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Naive(tt.input.a, tt.input.b)
			if got != tt.expected {
				t.Errorf("Naive(%s, %s) = %v; want %v", tt.input.a, tt.input.b, got, tt.expected)
			}
		})
	}
}
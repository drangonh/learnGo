package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		num int
	}{
		{"123456", 6},
		{"我是一只小小鸟", 5},
	}

	for _, v := range tests {
		if actual := lengthOfNonRepeatingSubStr(v.s); actual != v.num {
			t.Errorf("lengthOfNonRepeatingSubStr(%s);actual:%d,expect,%d", v.s, actual, v.num)
		}
	}
}

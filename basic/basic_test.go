package main

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{6, 8, 10},
		{5, 12, 13},
		{30000, 40000, 50000},
	}

	for _, v := range tests {
		if actual := calcTriangle(v.a, v.b); actual != v.c {
			t.Errorf("calcTriangle(%d,%d),actual:%d,expect:%d", v.a, v.b, actual, v.c)
		}
	}
}

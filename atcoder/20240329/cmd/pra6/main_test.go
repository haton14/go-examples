package main

import "testing"

func Test_valid(t *testing.T) {
	test := []struct {
		target int
		a      int
		b      int
		want   bool
	}{
		// {1, 1, 1, true},
		// {10, 1, 1, true},
		{100, 1, 1, true},
		// {11, 1, 2, true},
		// {19, 1, 10, true},
	}
	for _, tt := range test {
		got := valid(tt.target, tt.a, tt.b)
		if got != tt.want {
			t.Errorf("valid(%v, %v, %v) = %v; want %v", tt.target, tt.a, tt.b, got, tt.want)
		}
	}
}

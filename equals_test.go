package cyclicint

import (
	"math"
	"testing"
)

func TestEqualsUint16(t *testing.T) {
	tests := []struct {
		name string
		a, b uint16
		eps  uint16
		want bool
	}{
		{"difference less than eps and a > b", 100, 104, 5, true},
		{"difference less than eps and a < b", 104, 100, 5, true},
		{"difference equal to eps", 100, 105, 5, true},
		{"difference greater than eps", 100, 106, 5, false},
		{"zero eps", 100, 100, 0, true},
		{"integer overload and a > b", math.MaxUint16, 0, 1, true},
		{"integer overload and a < b", 0, math.MaxUint16, 1, true},
		{"max eps", 100, 100, math.MaxUint16, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equals(tt.a, tt.b, tt.eps)
			if got != tt.want {
				t.Errorf("Equals[uint16](%d, %d, %d) got = %v, want %v", tt.a, tt.b, tt.eps, got, tt.want)
			}
		})
	}
}

func TestEqualsInt(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		eps  int
		want bool
	}{
		{"difference less than eps and a > b", 100, 104, 5, true},
		{"difference less than eps and a < b", 104, 100, 5, true},
		{"difference equal to eps", 100, 105, 5, true},
		{"difference greater than eps", 100, 106, 5, false},
		{"zero eps", 100, 100, 0, true},
		{"integer overload and a > b", math.MaxInt, math.MinInt, 1, true},
		{"integer overload and a < b", math.MinInt, math.MaxInt, 1, true},
		{"max eps", 100, 100, math.MaxInt, true},
		{"min eps", 100, 100, math.MinInt, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equals(tt.a, tt.b, tt.eps)
			if got != tt.want {
				t.Errorf("Equals[int](%d, %d, %d) got = %v, want %v", tt.a, tt.b, tt.eps, got, tt.want)
			}
		})
	}
}


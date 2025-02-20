package main

import (
	"fmt"
	"testing"
)

// TestAdd tests the Add function.
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 2, 3},          // 1 + 2 = 3
		{-1, -2, -3},       // -1 + -2 = -3
		{0, 0, 0},          // 0 + 0 = 0
		{1000, 2000, 3000}, // 1000 + 2000 = 3000
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.result {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
			}
		})
	}
}

// TestMinus tests the minus function.
func TestMinus(t *testing.T) {
	tests := []struct {
		a, b   int
		result int
	}{
		{1, 2, -1},          // 1 - 2 = -1
		{-1, -2, 1},         // -1 - -2 = 1
		{0, 0, 0},           // 0 - 0 = 0
		{1000, 2000, -1000}, // 1000 - 2000 = -1000
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d", tt.a, tt.b), func(t *testing.T) {
			got := minus(tt.a, tt.b)
			if got != tt.result {
				t.Errorf("minus(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.result)
			}
		})
	}
}

package hops

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMinHops(t *testing.T) {
	testCases := []struct {
		name     string
		input    []TestCase
		expected []string
	}{
		{
			name: "Single test case with no obstacles",
			input: []TestCase{
				{
					GridWidth:  5,
					GridHeight: 5,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 4, Y: 4},
				},
			},
			expected: []string{"Optimal solution takes 3 hops."},
		},
		{
			name: "Test case with obstacles",
			input: []TestCase{
				{
					GridWidth:  5,
					GridHeight: 5,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 4, Y: 4},
					ObstacleBounds: []ObstacleBounds{
						{Left: 1, Right: 2, Upper: 1, Lower: 2},
					},
				},
			},
			expected: []string{"Optimal solution takes 4 hops."},
		},
		{
			name: "Multiple test cases",
			input: []TestCase{
				{
					GridWidth:  3,
					GridHeight: 3,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 2, Y: 2},
				},
				{
					GridWidth:  4,
					GridHeight: 4,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 3, Y: 3},
					ObstacleBounds: []ObstacleBounds{
						{Left: 1, Right: 2, Upper: 1, Lower: 2},
					},
				},
			},
			expected: []string{
				"Optimal solution takes 2 hops.",
				"Optimal solution takes 4 hops.",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MinHops(tc.input...)
			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("MinHops() result mismatch (-wanted +got):\n%s", diff)
			}
		})
	}
}

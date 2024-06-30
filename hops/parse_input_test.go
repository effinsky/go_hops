package hops

import (
	"strings"
	"testing"
)

func TestMustParseInputValid(t *testing.T) {
	tests := []struct {
		desc     string
		input    string
		expected []TestCase
	}{
		{
			desc: "Single test case with no obstacles",
			input: `
			1
			5 5
			0 0 4 4
			0
			`,
			expected: []TestCase{
				{
					GridWidth:      5,
					GridHeight:     5,
					StartPos:       Point{X: 0, Y: 0},
					FinishPos:      Point{X: 4, Y: 4},
					ObstacleBounds: []ObstacleBounds{},
				},
			},
		},
		{
			desc: "Single test case with obstacles",
			input: `
			1
			5 5
			0 0 4 4
			1
			2 2 2 2
			`,
			expected: []TestCase{
				{
					GridWidth:  5,
					GridHeight: 5,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 4, Y: 4},
					ObstacleBounds: []ObstacleBounds{
						{Left: 2, Right: 2, Upper: 2, Lower: 2},
					},
				},
			},
		},
		{
			desc: "Multiple test cases",
			input: `
			2 
			5 5
			0 0 4 4
			1
			2 2 2 2
			3 3
			0 0 2 2
			1
			1 1 1 1
			`,
			expected: []TestCase{
				{
					GridWidth:  5,
					GridHeight: 5,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 4, Y: 4},
					ObstacleBounds: []ObstacleBounds{
						{Left: 2, Right: 2, Upper: 2, Lower: 2},
					},
				},
				{
					GridWidth:  3,
					GridHeight: 3,
					StartPos:   Point{X: 0, Y: 0},
					FinishPos:  Point{X: 2, Y: 2},
					ObstacleBounds: []ObstacleBounds{
						{Left: 1, Right: 1, Upper: 1, Lower: 1},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			input := strings.TrimSpace(tt.input)
			result := MustParseInput(input)
			if len(result) != len(tt.expected) {
				t.Fatalf(
					"got %d test cases, want %d test cases",
					len(result), len(tt.expected),
				)
			}
			for i, testCase := range result {
				if testCase.GridWidth != tt.expected[i].GridWidth {
					t.Errorf(
						"got grid width %d, want %d",
						testCase.GridWidth, tt.expected[i].GridWidth,
					)
				}
				if testCase.GridHeight != tt.expected[i].GridHeight {
					t.Errorf(
						"got grid height %d, want %d",
						testCase.GridHeight, tt.expected[i].GridHeight,
					)
				}
				if testCase.StartPos != tt.expected[i].StartPos {
					t.Errorf(
						"got start position %+v, want %+v",
						testCase.StartPos, tt.expected[i].StartPos,
					)
				}
				if testCase.FinishPos != tt.expected[i].FinishPos {
					t.Errorf(
						"got finish position %+v, want %+v",
						testCase.FinishPos, tt.expected[i].FinishPos,
					)
				}
				if len(testCase.ObstacleBounds) !=
					len(tt.expected[i].ObstacleBounds) {
					t.Fatalf(
						"got %d obstacles, want %d obstacles",
						len(testCase.ObstacleBounds),
						len(tt.expected[i].ObstacleBounds),
					)
				}
				for j, obs := range testCase.ObstacleBounds {
					if obs != tt.expected[i].ObstacleBounds[j] {
						t.Errorf(
							"got obstacle %+v, want %+v",
							obs, tt.expected[i].ObstacleBounds[j],
						)
					}
				}
			}
		})
	}
}

func TestMustParseInputInvalid(t *testing.T) { // Expecting panics on bad input
	errorCases := []struct {
		desc  string
		input string
	}{
		{
			desc: "Invalid number of test cases",
			input: `
			abc
			5 5
			0 0 4 4
			0
			`,
		},
		{
			desc: "Invalid grid dimensions",
			input: `
			1
			5 abc
			0 0 4 4
			0
			`,
		},
		{
			desc: "Invalid start/finish points",
			input: `
			1
			5 5
			0 abc 4 4
			0
			`,
		},
		{
			desc: "Invalid number of obstacles",
			input: `
			1
			5 5
			0 0 4 4
			abc
			`,
		},
		{
			desc: "Invalid obstacle coordinates",
			input: `
			1
			5 5
			0 0 4 4
			1
			2 abc 2 2
			`,
		},
	}

	for _, tt := range errorCases {
		t.Run(tt.desc, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("expected panic but did not occur")
				}
			}()
			MustParseInput(strings.TrimSpace(tt.input))
		})
	}
}

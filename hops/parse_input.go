package hops

import (
	"fmt"
	"strconv"
	"strings"
)

// MustParseInput parses the input string and returns a slice of TestCase.
// It will panic on bad inputs. This is a script, it's over fast, so panicking
// here is fine by me.
func MustParseInput(input string) []TestCase {
	lines := strings.Split(input, "\n")
	// So much better for multiline strings getting out of whack.
	// Trim whitespaces per line and filter out empty lines.
	lines = filterMap(lines, func(l string) (string, bool) {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			return "", false
		}
		return l, true
	})

	numTestCases, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(fmt.Errorf("parsing number of test cases: %w", err))
	}

	idx := 1 // Start from line 2.
	var testCases []TestCase

	for tcIdx := range numTestCases {
		// Get grid dimensions for the case.
		dimensions := strings.Fields(lines[idx])
		if len(dimensions) != 2 {
			panic(fmt.Errorf("parsing grid dimensions at test case %d", tcIdx+1))
		}
		gridWidth, err := strconv.Atoi(dimensions[0])
		if err != nil {
			panic(fmt.Errorf(
				"parsing grid width at test case %d: %w", tcIdx+1, err,
			))
		}
		gridHeight, err := strconv.Atoi(dimensions[1])
		if err != nil {
			panic(fmt.Errorf(
				"parsing grid height at test case %d: %w", tcIdx+1, err,
			))
		}
		idx++

		// Get start and finish positions.
		points := strings.Fields(lines[idx])
		if len(points) != 4 {
			panic(fmt.Errorf(
				"parsing start/finish points at test case %d", tcIdx+1,
			))
		}
		startX, err := strconv.Atoi(points[0])
		if err != nil {
			panic(fmt.Errorf("parsing start X at test case %d: %w", tcIdx+1, err))
		}
		startY, err := strconv.Atoi(points[1])
		if err != nil {
			panic(fmt.Errorf("parsing start Y at test case %d: %w", tcIdx+1, err))
		}
		finishX, err := strconv.Atoi(points[2])
		if err != nil {
			panic(fmt.Errorf("parsing finish X at test case %d: %w", tcIdx+1, err))
		}
		finishY, err := strconv.Atoi(points[3])
		if err != nil {
			panic(fmt.Errorf(
				"parsing finish Y at test case %d: %w", tcIdx+1, err,
			))
		}
		start := Point{X: startX, Y: startY}
		finish := Point{X: finishX, Y: finishY}
		idx++

		// Get number of obstacles for grid.
		numObstacles, err := strconv.Atoi(lines[idx])
		if err != nil {
			panic(fmt.Errorf(
				"parsing number of obstacles at test case %d: %w", tcIdx+1, err,
			))
		}
		idx++

		// Get obstacle mapping from remaining lines
		var obstacleBounds []ObstacleBounds
		for obsIdx := range numObstacles {
			obstacleCoords := strings.Fields(lines[idx])
			if len(obstacleCoords) != 4 {
				panic(fmt.Errorf(
					"parsing obstacle at test case %d, obstacle %d",
					tcIdx+1, obsIdx+1,
				))
			}
			x1, err := strconv.Atoi(obstacleCoords[0])
			if err != nil {
				panic(fmt.Errorf(
					" parsing obstacle x1 at test case %d, obstacle %d: %w",
					tcIdx+1, obsIdx+1, err,
				))
			}
			x2, err := strconv.Atoi(obstacleCoords[1])
			if err != nil {
				panic(fmt.Errorf(
					" parsing obstacle x2 at test case %d, obstacle %d: %w",
					tcIdx+1, obsIdx+1, err,
				))
			}
			y1, err := strconv.Atoi(obstacleCoords[2])
			if err != nil {
				panic(fmt.Errorf(
					" parsing obstacle y1 at test case %d, obstacle %d: %w",
					tcIdx+1, obsIdx+1, err,
				))
			}
			y2, err := strconv.Atoi(obstacleCoords[3])
			if err != nil {
				panic(fmt.Errorf(
					"parsing obstacle y2 at test case %d, obstacle %d: %w",
					tcIdx+1, obsIdx+1, err,
				))
			}
			obstacleBounds = append(obstacleBounds, ObstacleBounds{
				Left:  x1,
				Right: x2,
				Upper: y1,
				Lower: y2,
			})
			idx++
		}
		testCases = append(testCases, TestCase{
			GridWidth:      gridWidth,
			GridHeight:     gridHeight,
			StartPos:       start,
			FinishPos:      finish,
			ObstacleBounds: obstacleBounds,
		})
	}
	return testCases
}

func filterMap[T, R any](in []T, f func(it T) (R, bool)) []R {
	out := make([]R, 0, len(in))
	for i := range in {
		v, ok := f(in[i])
		if ok {
			out = append(out, v)
		}
	}
	return out
}

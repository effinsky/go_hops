package hops

import (
	"container/list"
	"fmt"
)

const (
	optimalSolution = "Optimal solution takes %d hops."
	noSolution      = "No solution."

	maxSpeed = 3
)

type (
	Grid struct {
		Cells  Cells
		Width  int
		Height int
	}

	Cells [][]int

	HopperState struct {
		Pos  Point
		VX   int
		VY   int
		hops int
	}

	Point struct {
		X int
		Y int
	}

	TestCase struct {
		GridWidth      int
		GridHeight     int
		StartPos       Point
		FinishPos      Point
		ObstacleBounds []ObstacleBounds
	}

	ObstacleBounds struct {
		Left  int
		Right int
		Upper int
		Lower int
	}
)

// MinHops parses the input string, processes each test case, and returns the
// results.
func MinHops(testcases ...TestCase) []string {
	var results []string
	for _, tc := range testcases {
		grid := makeGrid(tc.GridWidth, tc.GridHeight, tc.ObstacleBounds)
		result := breadthFirstSearch(tc.StartPos, tc.FinishPos, grid)
		results = append(results, result)
	}
	return results
}

func makeGrid(width int, height int, obs []ObstacleBounds) Grid {
	cells := make(Cells, height)
	for i := range cells {
		cells[i] = make([]int, width) // Explicitly initialize the grid to 0s.
	}
	for _, ob := range obs {
		for x := ob.Left; x <= ob.Right; x++ {
			for y := ob.Upper; y <= ob.Lower; y++ {
				cells[y][x] = 1
			}
		}
	}
	return Grid{
		Cells:  cells,
		Width:  width,
		Height: height,
	}
}

// breadthFirstSearch does a breadth-first search to find the shortest path
// from the start to the finish point on the given grid.
//
// The function initializes the search with a starting state. This initial
// state is added to a queue. A map is used to keep track of visited states
// to prevent redundant/repeated processing of same states.
//
// State processing happens through a loop until there are no more  hopper
// states in the queue.
//  1. We remove the frontmost state from the queue and check if its
//     position matches the finish. If so, we return the optimal hops msg.
//  2. If finish is not reached, we generate all new possible states
//     from the current state.
//  3. For each new state generated, if it has not been visited, it is marked
//     as visited and added to the queue for further processing.
//  4. We repeat the steps until we either get a match vs the finish pos or run
//     out of states to process. In the latter case, we return the failure msg.
func breadthFirstSearch(start Point, finish Point, grid Grid) string {
	directions := [3]int{-1, 0, 1}       // Do not mutate --
	initState := HopperState{Pos: start} // + vx, vy, hops
	queue := list.New()
	queue.PushBack(initState)

	visited := make(map[HopperState]bool)
	visited[initState] = true

	for queue.Len() > 0 {
		cs := queue.Remove(queue.Front()).(HopperState)
		if cs.Pos == finish {
			return fmt.Sprintf(optimalSolution, cs.hops)
		}

		for _, ns := range generateNewStates(cs, directions, maxSpeed, grid) {
			if !visited[ns] {
				visited[ns] = true
				queue.PushBack(ns)
			}
		}
	}
	return noSolution
}

// generateNewStates generates a list of possible new states from the current
// state based on the given directions, grid occupancy, and maximum velocity
// constraints.
func generateNewStates(
	currState HopperState,
	directions [3]int,
	maxV int,
	g Grid,
) []HopperState {
	var newStates []HopperState
	for _, dvx := range directions {
		for _, dvy := range directions {
			newVx, newVy := currState.VX+dvx, currState.VY+dvy
			if isValidVelocity(newVx, newVy, maxV) {
				newPosX, newPosY := currState.Pos.X+newVx, currState.Pos.Y+newVy
				if isValidPosition(newPosX, newPosY, g) {
					newState := HopperState{
						Pos:  Point{X: newPosX, Y: newPosY},
						VX:   newVx,
						VY:   newVy,
						hops: currState.hops + 1,
					}
					newStates = append(newStates, newState)
				}
			}
		}
	}
	return newStates
}

func isValidPosition(newPosX int, newPosY int, g Grid) bool {
	return isPositionWithinBounds(newPosX, newPosY, g.Width, g.Height) &&
		isCellNotOccupied(newPosX, newPosY, g)
}

func isCellNotOccupied(newPosX int, newPosY int, g Grid) bool {
	return g.Cells[newPosY][newPosX] == 0
}

func isPositionWithinBounds(posX, posY, gridWidth, gridHeight int) bool {
	return (posX >= 0 && posX < gridWidth) && (posY >= 0 && posY < gridHeight)
}

func isValidVelocity(vx, vy, maxSpeed int) bool {
	return (vx >= -maxSpeed && vx <= maxSpeed) &&
		(vy >= -maxSpeed && vy <= maxSpeed)
}
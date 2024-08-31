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
	Grid [][]int

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

// MinHops processes each test case using bfs search and returns the results.
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
	g := make(Grid, height)
	for i := range g {
		g[i] = make([]int, width) // Explicitly initialize the grid to 0s.
	}
	for _, ob := range obs {
		for x := ob.Left; x <= ob.Right; x++ {
			for y := ob.Upper; y <= ob.Lower; y++ {
				g[y][x] = 1
			}
		}
	}
	return g
}

// breadthFirstSearch does a bfs to find the shortest path from the start to
// the finish point on the given grid.
//
// The function initializes the search with a starting state that is then added
// to a queue. A map is used to keep track of visited states to prevent
// redundant/repeated processing.
//
// State processing happens in a loop until there are no more hopper states
// in the queue.
//  1. We remove the frontmost state from the queue and check if its
//     position matches the finish position. If so, we return the optimal hops
//     message.
//  2. If finish position is not reached, we generate all new possible states
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

	// Since we have no set to work with in stdlib, using map state to empty
	// struct map.
	visited := make(map[HopperState]struct{})
	visited[initState] = struct{}{}

	for queue.Len() > 0 {
		// Skipping conversion error check.
		cs := queue.Remove(queue.Front()).(HopperState)
		if cs.Pos == finish {
			return fmt.Sprintf(optimalSolution, cs.hops)
		}

		for _, ns := range generateNewStates(cs, directions, maxSpeed, grid) {
			// If not visited, now visited and add to enqueue.
			if _, ok := visited[ns]; !ok {
				visited[ns] = struct{}{}
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
	// If within given directions we move at a valid velocity (slowing down,
	// maintaining speed, speeding up), and then land on a valid position
	// (within bounds and not on an obstacle), then we have a new valid state.
	newStates := make([]HopperState, 0)
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
	return isPositionWithinBounds(
		newPosX,
		newPosY,
		len(g[0]), // grid width
		len(g),    // grid height
	) &&
		!isCellOccupied(newPosX, newPosY, g)
}

func isCellOccupied(newPosX int, newPosY int, g Grid) bool {
	return g[newPosY][newPosX] == 1
}

func isPositionWithinBounds(posX, posY, gridWidth, gridHeight int) bool {
	return (posX >= 0 && posX < gridWidth) && (posY >= 0 && posY < gridHeight)
}

func isValidVelocity(vx, vy, maxSpeed int) bool {
	return (vx >= -maxSpeed && vx <= maxSpeed) &&
		(vy >= -maxSpeed && vy <= maxSpeed)
}

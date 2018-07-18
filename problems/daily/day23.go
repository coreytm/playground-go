package daily

import (
	"fmt"
	"math"
)

// This problem was asked by Google.

// You are given an M by N matrix consisting of booleans that represents a board.
// Each True boolean represents a wall.
// Each False boolean represents a tile you can walk on.
// Given this matrix, a start coordinate, and an end coordinate, return the minimum number of steps required to reach the end coordinate from the start.
// If there is no possible path, then return null.
// You can move up, left, down, and right.
// You cannot move through walls.
// You cannot wrap around the edges of the board.
// For example, given the following board:

// [[f, f, f, f],
// [t, t, f, t],
// [f, f, f, f],
// [f, f, f, f]]

// and start = (3, 0) (bottom left) and end = (0, 0) (top left),
// the minimum number of steps required to reach the end is 7,
// since we would need to go through (1, 2) because there is a wall everywhere else on the second row.

// shortestPath returns the shortest path from [m-1, n-1] to [0, 0].
// An error is returned if there is no path.
//
// This function implements a BFS iteratively using a stack and memoization to remember cost of visiting previous nodes.
// Runtime: O(n) because it may have to traverse all nodes if shortest path is the last one.
// Space:   O(n) because of memoization array and stack.
func shortestPath(in [][]bool) (int, error) {
	m := len(in)
	if m == 0 {
		return 0, fmt.Errorf("input is empty")
	}
	n := len(in[0])
	if n == 0 {
		return 0, fmt.Errorf("input is empty")
	}

	hops := make([][]int, m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			hops[i][j] = math.MaxInt32
		}
	}

	hops[m-1][n-1] = 0

	cur := []int{m - 1, n - 1}
	stack := [][]int{cur}

	for len(stack) > 0 {
		// pop the stack
		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		neighbours := getNeighbours(in, hops, cur[0], cur[1])
		stack = append(stack, neighbours...)
	}

	if hops[0][0] == math.MaxInt32 {
		return 0, fmt.Errorf("there is no path from [%d, %d] to [0, 0]", m-1, n-1)
	}

	return hops[0][0], nil
}

// getNeighbours will return the immediate neighbours of [i, j]
// (i.e. top [i-1, j], bottom [[i+1, j], left [i, j-1], right [i, j+1])
// if they are not false (i.e. no wall in the way)
// if they are not out of bound of [m, n]
// if the number of hops to get from [i, j] to neighbour is less than existing number of hops to get there
func getNeighbours(in [][]bool, hops [][]int, i, j int) [][]int {
	neighbours := [][]int{}

	// assume in is not empty
	m := len(in)
	n := len(in[0])

	// top
	if i > 0 && !in[i-1][j] && hops[i][j]+1 < hops[i-1][j] {
		neighbours = append(neighbours, []int{i - 1, j})
		hops[i-1][j] = hops[i][j] + 1
	}

	// bottom
	if i < m-1 && !in[i+1][j] && hops[i][j]+1 < hops[i+1][j] {
		neighbours = append(neighbours, []int{i + 1, j})
		hops[i+1][j] = hops[i][j] + 1
	}

	// left
	if j > 0 && !in[i][j-1] && hops[i][j]+1 < hops[i][j-1] {
		neighbours = append(neighbours, []int{i, j - 1})
		hops[i][j-1] = hops[i][j] + 1
	}

	// right
	if j < n-1 && !in[i][j+1] && hops[i][j]+1 < hops[i][j+1] {
		neighbours = append(neighbours, []int{i, j + 1})
		hops[i][j+1] = hops[i][j] + 1
	}

	return neighbours
}

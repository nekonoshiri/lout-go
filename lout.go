// Package lout provides types and functions to implement
// a low-level runtime for "Lights Out"-like games (lout).
package lout

// Grid is a 2D grid of lights.
//
// Let g be a Grid.
// If g[y][x] is true, the light at the coordinates (x, y) is on.
// If g[y][x] is false, the light at the coordinates (x, y) is off.
//
// Use [NewGrid] to create a new Grid.
type Grid [][]bool

// NewGrid creates a new Grid with the given width and height.
// All lights in the created Grid are initially off (false).
//
// Width and height must be positive (greater than zero).
// Panics if width <= 0 or height <= 0.
func NewGrid(width, height int) Grid {
	if width <= 0 || height <= 0 {
		panic("width and height must be positive")
	}
	grid := make(Grid, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

// Width returns the width of the Grid.
func (g Grid) Width() int {
	return len(g[0])
}

// Height returns the height of the Grid.
func (g Grid) Height() int {
	return len(g)
}

// Set sets the on/off state of the light at (x, y).
//
// Returns false if the coordinates are out of bounds.
func (g Grid) Set(x, y int, on bool) bool {
	if x < 0 || x >= g.Width() || y < 0 || y >= g.Height() {
		return false
	}
	g[y][x] = on
	return true
}

// Press flips the on/off state of the light at (x, y) and its four adjacent lights.
// The adjacent lights are the horizontal neighbors (x-1, y), (x+1, y) and
// the vertical neighbors (x, y-1), (x, y+1).
//
// Lights at coordinates outside the grid boundaries are ignored.
//
// If (x, y) itself is outside the grid, the method still processes any adjacent lights
// that are within the grid boundaries. For example, if the grid is 4x3 and Press(-1, 0)
// is called, only the light at (0, 0) will be flipped because it's the only adjacent
// light that is within the grid.
func (g Grid) Press(x, y int) {
	w := g.Width()
	h := g.Height()

	toggle := func(x, y int) {
		if x >= 0 && x < w && y >= 0 && y < h {
			g[y][x] = !g[y][x]
		}
	}

	toggle(x, y)
	toggle(x-1, y)
	toggle(x+1, y)
	toggle(x, y-1)
	toggle(x, y+1)
}

// IsOn returns true if the light at the given coordinates is on, false otherwise.
//
// Returns false if the coordinates are out of bounds.
func (g Grid) IsOn(x, y int) bool {
	if x < 0 || x >= g.Width() || y < 0 || y >= g.Height() {
		return false
	}
	return g[y][x]
}

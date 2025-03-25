package lout

import (
	"fmt"
	"testing"
)

func TestNewGrid(t *testing.T) {
	for _, tc := range []struct {
		width  int
		height int
	}{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 2}, {2, 3},
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", tc.width, tc.height), func(t *testing.T) {
			g := NewGrid(tc.width, tc.height)
			if g.Width() != tc.width {
				t.Errorf("g.Width() = %d, want %d", g.Width(), tc.width)
			}
			if g.Height() != tc.height {
				t.Errorf("g.Height() = %d, want %d", g.Height(), tc.height)
			}

			for y := range g.Height() {
				for x := range g.Width() {
					if g.IsOn(x, y) {
						t.Errorf("g.IsOn(%d, %d) = true, want false", x, y)
					}
				}
			}
		})
	}
}

func TestGrid_Width(t *testing.T) {
	for _, tc := range []struct {
		width  int
		height int
	}{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 2}, {2, 3},
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", tc.width, tc.height), func(t *testing.T) {
			g := NewGrid(tc.width, tc.height)
			if g.Width() != tc.width {
				t.Errorf("g.Width() = %d, want %d", g.Width(), tc.width)
			}
		})
	}
}

func TestGrid_Height(t *testing.T) {
	for _, tc := range []struct {
		width  int
		height int
	}{
		{1, 1}, {1, 2}, {1, 3},
		{2, 1}, {2, 2}, {2, 3},
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", tc.width, tc.height), func(t *testing.T) {
			g := NewGrid(tc.width, tc.height)
			if g.Height() != tc.height {
				t.Errorf("g.Height() = %d, want %d", g.Height(), tc.height)
			}
		})
	}
}

func TestNewGrid_Panics(t *testing.T) {
	for _, tc := range []struct {
		width  int
		height int
	}{
		{0, 1}, {1, 0}, {0, 0},
		{-1, 1}, {1, -1}, {-1, -1},
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", tc.width, tc.height), func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Errorf("NewGrid(%d, %d) did not panic", tc.width, tc.height)
				}
			}()
			_ = NewGrid(tc.width, tc.height)
		})
	}
}

func TestGrid_Set(t *testing.T) {
	for _, g := range []Grid{
		NewGrid(1, 1),
		NewGrid(2, 3),
		NewGrid(3, 2),
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", g.Width(), g.Height()), func(t *testing.T) {
			for y := range g.Height() {
				for x := range g.Width() {
					if ok := g.Set(x, y, true); !ok {
						t.Errorf("g.Set(%d, %d, true) = false, want true", x, y)
					}
					if !g.IsOn(x, y) {
						t.Errorf("g.IsOn(%d, %d) = false after g.Set(%[1]d, %[2]d, true), want true", x, y)
					}

					if ok := g.Set(x, y, false); !ok {
						t.Errorf("g.Set(%d, %d, false) = false, want true", x, y)
					}
					if g.IsOn(x, y) {
						t.Errorf("g.IsOn(%d, %d) = true after g.Set(%[1]d, %[2]d, false), want false", x, y)
					}
				}
			}
		})
	}
}

func TestGrid_Set_OutOfBounds(t *testing.T) {
	for _, g := range []Grid{
		NewGrid(1, 1),
		NewGrid(2, 3),
		NewGrid(3, 2),
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", g.Width(), g.Height()), func(t *testing.T) {
			for _, c := range []struct {
				x int
				y int
			}{
				{-1, 0}, {0, -1}, {-1, -1},
				{g.Width(), 0}, {0, g.Height()}, {g.Width(), g.Height()},
			} {
				if ok := g.Set(c.x, c.y, true); ok {
					t.Errorf("g.Set(%d, %d, true) = true, want false", c.x, c.y)
				}
			}
		})
	}
}

func TestGrid_Press(t *testing.T) {
	// createGrid("000", "111") creates the following grid:
	//  ----> x
	// | 000
	// | 111
	// v
	// y
	createGrid := func(s ...string) Grid {
		g := NewGrid(len(s[0]), len(s))
		for y := range g.Height() {
			for x := range g.Width() {
				g.Set(x, y, s[y][x] == '1')
			}
		}
		return g
	}

	for _, tc := range []struct {
		g    Grid
		x    int
		y    int
		want Grid
	}{
		{createGrid("0000", "0000", "0000"), -1, -1, createGrid("0000", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 0, -1, createGrid("1000", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 1, -1, createGrid("0100", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 2, -1, createGrid("0010", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 3, -1, createGrid("0001", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 4, -1, createGrid("0000", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), -1, 0, createGrid("1000", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 0, 0, createGrid("1100", "1000", "0000")},
		{createGrid("0000", "0000", "0000"), 1, 0, createGrid("1110", "0100", "0000")},
		{createGrid("0000", "0000", "0000"), 2, 0, createGrid("0111", "0010", "0000")},
		{createGrid("0000", "0000", "0000"), 3, 0, createGrid("0011", "0001", "0000")},
		{createGrid("0000", "0000", "0000"), 4, 0, createGrid("0001", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), -1, 1, createGrid("0000", "1000", "0000")},
		{createGrid("0000", "0000", "0000"), 0, 1, createGrid("1000", "1100", "1000")},
		{createGrid("0000", "0000", "0000"), 1, 1, createGrid("0100", "1110", "0100")},
		{createGrid("0000", "0000", "0000"), 2, 1, createGrid("0010", "0111", "0010")},
		{createGrid("0000", "0000", "0000"), 3, 1, createGrid("0001", "0011", "0001")},
		{createGrid("0000", "0000", "0000"), 4, 1, createGrid("0000", "0001", "0000")},
		{createGrid("0000", "0000", "0000"), -1, 2, createGrid("0000", "0000", "1000")},
		{createGrid("0000", "0000", "0000"), 0, 2, createGrid("0000", "1000", "1100")},
		{createGrid("0000", "0000", "0000"), 1, 2, createGrid("0000", "0100", "1110")},
		{createGrid("0000", "0000", "0000"), 2, 2, createGrid("0000", "0010", "0111")},
		{createGrid("0000", "0000", "0000"), 3, 2, createGrid("0000", "0001", "0011")},
		{createGrid("0000", "0000", "0000"), 4, 2, createGrid("0000", "0000", "0001")},
		{createGrid("0000", "0000", "0000"), -1, 3, createGrid("0000", "0000", "0000")},
		{createGrid("0000", "0000", "0000"), 0, 3, createGrid("0000", "0000", "1000")},
		{createGrid("0000", "0000", "0000"), 1, 3, createGrid("0000", "0000", "0100")},
		{createGrid("0000", "0000", "0000"), 2, 3, createGrid("0000", "0000", "0010")},
		{createGrid("0000", "0000", "0000"), 3, 3, createGrid("0000", "0000", "0001")},
		{createGrid("0000", "0000", "0000"), 4, 3, createGrid("0000", "0000", "0000")},

		{createGrid("1111", "1111", "1111"), -1, -1, createGrid("1111", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 0, -1, createGrid("0111", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 1, -1, createGrid("1011", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 2, -1, createGrid("1101", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 3, -1, createGrid("1110", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 4, -1, createGrid("1111", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), -1, 0, createGrid("0111", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 0, 0, createGrid("0011", "0111", "1111")},
		{createGrid("1111", "1111", "1111"), 1, 0, createGrid("0001", "1011", "1111")},
		{createGrid("1111", "1111", "1111"), 2, 0, createGrid("1000", "1101", "1111")},
		{createGrid("1111", "1111", "1111"), 3, 0, createGrid("1100", "1110", "1111")},
		{createGrid("1111", "1111", "1111"), 4, 0, createGrid("1110", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), -1, 1, createGrid("1111", "0111", "1111")},
		{createGrid("1111", "1111", "1111"), 0, 1, createGrid("0111", "0011", "0111")},
		{createGrid("1111", "1111", "1111"), 1, 1, createGrid("1011", "0001", "1011")},
		{createGrid("1111", "1111", "1111"), 2, 1, createGrid("1101", "1000", "1101")},
		{createGrid("1111", "1111", "1111"), 3, 1, createGrid("1110", "1100", "1110")},
		{createGrid("1111", "1111", "1111"), 4, 1, createGrid("1111", "1110", "1111")},
		{createGrid("1111", "1111", "1111"), -1, 2, createGrid("1111", "1111", "0111")},
		{createGrid("1111", "1111", "1111"), 0, 2, createGrid("1111", "0111", "0011")},
		{createGrid("1111", "1111", "1111"), 1, 2, createGrid("1111", "1011", "0001")},
		{createGrid("1111", "1111", "1111"), 2, 2, createGrid("1111", "1101", "1000")},
		{createGrid("1111", "1111", "1111"), 3, 2, createGrid("1111", "1110", "1100")},
		{createGrid("1111", "1111", "1111"), 4, 2, createGrid("1111", "1111", "1110")},
		{createGrid("1111", "1111", "1111"), -1, 3, createGrid("1111", "1111", "1111")},
		{createGrid("1111", "1111", "1111"), 0, 3, createGrid("1111", "1111", "0111")},
		{createGrid("1111", "1111", "1111"), 1, 3, createGrid("1111", "1111", "1011")},
		{createGrid("1111", "1111", "1111"), 2, 3, createGrid("1111", "1111", "1101")},
		{createGrid("1111", "1111", "1111"), 3, 3, createGrid("1111", "1111", "1110")},
		{createGrid("1111", "1111", "1111"), 4, 3, createGrid("1111", "1111", "1111")},

		{createGrid("0000", "1111", "0000"), -1, -1, createGrid("0000", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 0, -1, createGrid("1000", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 1, -1, createGrid("0100", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 2, -1, createGrid("0010", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 3, -1, createGrid("0001", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 4, -1, createGrid("0000", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), -1, 0, createGrid("1000", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 0, 0, createGrid("1100", "0111", "0000")},
		{createGrid("0000", "1111", "0000"), 1, 0, createGrid("1110", "1011", "0000")},
		{createGrid("0000", "1111", "0000"), 2, 0, createGrid("0111", "1101", "0000")},
		{createGrid("0000", "1111", "0000"), 3, 0, createGrid("0011", "1110", "0000")},
		{createGrid("0000", "1111", "0000"), 4, 0, createGrid("0001", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), -1, 1, createGrid("0000", "0111", "0000")},
		{createGrid("0000", "1111", "0000"), 0, 1, createGrid("1000", "0011", "1000")},
		{createGrid("0000", "1111", "0000"), 1, 1, createGrid("0100", "0001", "0100")},
		{createGrid("0000", "1111", "0000"), 2, 1, createGrid("0010", "1000", "0010")},
		{createGrid("0000", "1111", "0000"), 3, 1, createGrid("0001", "1100", "0001")},
		{createGrid("0000", "1111", "0000"), 4, 1, createGrid("0000", "1110", "0000")},
		{createGrid("0000", "1111", "0000"), -1, 2, createGrid("0000", "1111", "1000")},
		{createGrid("0000", "1111", "0000"), 0, 2, createGrid("0000", "0111", "1100")},
		{createGrid("0000", "1111", "0000"), 1, 2, createGrid("0000", "1011", "1110")},
		{createGrid("0000", "1111", "0000"), 2, 2, createGrid("0000", "1101", "0111")},
		{createGrid("0000", "1111", "0000"), 3, 2, createGrid("0000", "1110", "0011")},
		{createGrid("0000", "1111", "0000"), 4, 2, createGrid("0000", "1111", "0001")},
		{createGrid("0000", "1111", "0000"), -1, 3, createGrid("0000", "1111", "0000")},
		{createGrid("0000", "1111", "0000"), 0, 3, createGrid("0000", "1111", "1000")},
		{createGrid("0000", "1111", "0000"), 1, 3, createGrid("0000", "1111", "0100")},
		{createGrid("0000", "1111", "0000"), 2, 3, createGrid("0000", "1111", "0010")},
		{createGrid("0000", "1111", "0000"), 3, 3, createGrid("0000", "1111", "0001")},
		{createGrid("0000", "1111", "0000"), 4, 3, createGrid("0000", "1111", "0000")},
	} {
		t.Run(fmt.Sprintf("g=%v,x=%d,y=%d", tc.g, tc.x, tc.y), func(t *testing.T) {
			tc.g.Press(tc.x, tc.y)

			for y := range tc.want.Height() {
				for x := range tc.want.Width() {
					if tc.g.IsOn(x, y) != tc.want.IsOn(x, y) {
						t.Fatalf("g.Press(%d, %d) = %v, want %v", tc.x, tc.y, tc.g, tc.want)
					}
				}
			}
		})
	}
}

func TestGrid_IsOn(t *testing.T) {
	for _, g := range []Grid{
		NewGrid(1, 1),
		NewGrid(2, 3),
		NewGrid(3, 2),
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", g.Width(), g.Height()), func(t *testing.T) {
			for y := range g.Height() {
				for x := range g.Width() {
					_ = g.Set(x, y, true)
					if !g.IsOn(x, y) {
						t.Errorf("g.IsOn(%d, %d) = false after g.Set(%[1]d, %[2]d, true), want true", x, y)
					}

					_ = g.Set(x, y, false)
					if g.IsOn(x, y) {
						t.Errorf("g.IsOn(%d, %d) = true after g.Set(%[1]d, %[2]d, false), want false", x, y)
					}
				}
			}
		})
	}
}

func TestGrid_IsOn_OutOfBounds(t *testing.T) {
	for _, g := range []Grid{
		NewGrid(1, 1),
		NewGrid(2, 3),
		NewGrid(3, 2),
	} {
		t.Run(fmt.Sprintf("g=NewGrid(%d,%d)", g.Width(), g.Height()), func(t *testing.T) {
			for _, c := range []struct {
				x int
				y int
			}{
				{-1, 0}, {0, -1}, {-1, -1},
				{g.Width(), 0}, {0, g.Height()}, {g.Width(), g.Height()},
			} {
				if g.IsOn(c.x, c.y) {
					t.Errorf("g.IsOn(%d, %d) = true, want false", c.x, c.y)
				}
			}
		})
	}
}

package arraygrid

import (
	"fmt"
	"io"
)

type Grid [][]string

func New(input []string) Grid {
	grid := make(Grid, len(input))

	for i := range input {
		s := input[i]
		row := make([]string, 0, len(s))
		for _, r := range s {
			row = append(row, string(r))
		}
		grid[i] = row
	}

	return grid
}

func (g Grid) Print(w io.Writer) {
	fmt.Fprint(w, "    ")
	for x := 0; x < len(g[0]); x++ {
		fmt.Fprintf(w, "%3d ", x)
	}

	fmt.Fprintln(w)
	for y := 0; y <= len(g)-1; y++ {
		fmt.Fprintf(w, "%2d ", y)
		for _, cell := range g[y] {
			fmt.Fprintf(w, "%4s", cell)
		}
		fmt.Fprintln(w)
	}

}

func (g Grid) Get(x, y int) string {
	return g[x][y]
}

// set x,y in grid.
// no-op if x or y is invalid
func (g *Grid) Set(x, y int, v string) {
	if g == nil {
		return
	}
	if y < 0 || y >= len(*g) {
		return
	}
	if x < 0 || x >= len((*g)[y]) {
		return
	}

	(*g)[x][y] = v
}

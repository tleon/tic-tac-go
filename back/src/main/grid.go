package main

// Cell Represent a Cell of the Grid
type Cell struct {
	IsChecked bool
	ByPlayer  int
}

// CheckBox Mark a cell as checked by a player
func (b *Cell) CheckBox(player int) {
	b.IsChecked = true
	b.ByPlayer = player
}

// Grid The complete Grid of Cell
type Grid struct {
	content [3][3]Cell
}

// HandlePlay Modify the grid struct to take in account the last cell checked
func (grid *Grid) HandlePlay(line int, cell int, player int) Grid {
	grid.content[line][cell].CheckBox(player)
	return *grid
}

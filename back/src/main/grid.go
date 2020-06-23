package main

type Cell struct {
	IsChecked bool
	ByPlayer int
}

func (b *Cell) CheckBox(player int) {
	b.IsChecked = true
	b.ByPlayer = player
}

type Grid struct {
	content [3][3]Cell
}

// coords can be A0-2 || B0-2  || C0-2
func HandlePlay(grid *Grid, line string, cell int, player int) Grid {
	switch line {
		case "A" :
			grid.content[0][cell].CheckBox(player)
		case "B" :
			grid.content[1][cell].CheckBox(player)
		case "C" :
			grid.content[2][cell].CheckBox(player)
	}

	return *grid
}
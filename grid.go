package gonogram

type cell int

const (
	unknown cell = iota
	black
	white
)

type clue int

type line []cell

func (l line) solve(clues []clue) {

}

func (l line) mostLeftSolution(clues []clue) (clueStart []int) {
	clueIx := 0
	clueStart = make([]int, len(clues))
	// We try to fit the block. If we find a contradiction we shift the start of the block
	// If we are at the end of the block we move to the next clue
	for i := 0; i < len(l); i++ {
		blockLength := i - clueStart[clueIx] + 1
		switch l[i] {
		case white:
			if blockLength == int(clues[clueIx])+1 {
				clueIx++
			}
			// We should always move the start of the block
			// Either we just moved to the next block
			// Or we have a white in the middel of the block (not possible)
			clueStart[clueIx] = i + 1
		case black:
			if blockLength == int(clues[clueIx])+1 {
				clueStart[clueIx]++
			}
		case unknown:
			if blockLength == int(clues[clueIx])+1 {
				clueIx++
				clueStart[clueIx] = i + 1
			}
		}
	}
	return clueStart
}

func (l line) mostRightSolution(clues []clue) (clueStart []int) {
	clueIx := len(clues) - 1
	clueStart = make([]int, len(clues))
	// Same as for the left but we start from the right and work back
	for i := len(l); i >= 0; i-- {
		blockLength := clueStart[clueIx] - i + 1
		switch l[i] {
		case white:
			if blockLength == int(clues[clueIx])+1 {
				clueIx--
			}

			clueStart[clueIx] = i - 1
		case black:
			if blockLength == int(clues[clueIx])+1 {
				clueStart[clueIx]--
			}
		case unknown:
			if blockLength == int(clues[clueIx])+1 {
				clueIx--
				clueStart[clueIx] = i - 1
			}
		}
	}
	return clueStart
}

type grid struct {
	colClues [][]clue
	rowClues [][]clue
	cells    []line
}

func new(colClues [][]clue, rowClues [][]clue) *grid {
	cells := make([]line, len(rowClues))
	for i := range cells {
		cells[i] = make(line, len(colClues))
	}
	return &grid{colClues, rowClues, cells}
}

func (g *grid) copyOfRow(row int) (result line) {
	copy(result, g.cells[row])
	return result
}

func (g *grid) copyOfCol(col int) (result line) {
	result = make(line, len(g.rowClues))
	for row := 0; row <= len(result); row++ {
		result[row] = g.cells[row][col]
	}
	return result
}

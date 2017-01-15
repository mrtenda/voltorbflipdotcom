package voltorbflip

import "fmt"

func (psol *vfPartialSolution) applyHeuristics() bool {
	invertedDataIsOutOfDate := false

	rowWasUpdated := false
	columnWasUpdated := false

	for {
		for y := range psol.tiles {
			row := &psol.tiles[y]
			rowData := &psol.rowData[y]
			if rowData.NumUnsolvedTiles == 0 {
				continue
			}
			for _, hf := range allVfHeuristicFuncs {
				wasUpdated := hf(*rowData, row)
				if wasUpdated {
					//fmt.Printf("Updated y=%d, hi=%d, t: %v\n", y, hi, psol.tiles[y])
					if !psol.updateRowData(y) {
						return false
					}
					invertedDataIsOutOfDate = true
					rowWasUpdated = true
					break
				}
			}
		}

		if rowWasUpdated {
			rowWasUpdated = false
			continue
		}

		if invertedDataIsOutOfDate {
			psol.updateInvertedTilesFromTiles()
			if !psol.updateAllColumnData() {
				return false
			}
			invertedDataIsOutOfDate = false
		}

		for x := range psol.invertedTiles {
			column := &psol.invertedTiles[x]
			columnData := &psol.columnData[x]
			if columnData.NumUnsolvedTiles == 0 {
				continue
			}
			for _, hf := range allVfHeuristicFuncs {
				wasUpdated := hf(*columnData, column)
				if wasUpdated {
					//fmt.Printf("Updated x=%d, hi=%d, t: %v\n", x, hi, psol.invertedTiles[x])
					if !psol.updateColumnData(x) {
						return false
					}
					columnWasUpdated = true
					break
				}
			}
		}

		if columnWasUpdated {
			columnWasUpdated = false
			psol.updateTilesFromInvertedTiles()
			if !psol.updateAllRowData() {
				return false
			}
			continue
		}

		return true
	}
}

type vfHeuristicFunc func(lineData vfPSolLineData, line *[5]VfPSolTile) bool

var allVfHeuristicFuncs = []vfHeuristicFunc{
	heuristic0,
	heuristic1,
	heuristic2,
	heuristic3,
	heuristic4,
	heuristic5,
	heuristic6,
	heuristic7,
	heuristic8,
}

// Heuristic #0 - If RemainingVoltorbs + RemainingPoints == NumUnsolvedTiles, eliminate all possibilities except V and 1 from unsolved tiles
func heuristic0(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.RemainingVoltorbs+lineData.RemainingPoints != lineData.NumUnsolvedTiles {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			wasUpdated = wasUpdated || tile[2] || tile[3]
			tile[2] = false
			tile[3] = false
		}
	}
	return wasUpdated
}

// Heuristic #1 - If NumUnsolvedTiles - RemainingVoltorbs == RemainingPoints - 1, remove 3 as a possible option
func heuristic1(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.NumUnsolvedTiles-lineData.RemainingVoltorbs != lineData.RemainingPoints-1 {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() && tile[3] {
			tile[3] = false
			wasUpdated = true
		}
	}
	return wasUpdated
}

// Heuristic #2 - If RemainingVoltorbs == 0, eliminate any possible voltorbs
func heuristic2(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.RemainingVoltorbs != 0 {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() && tile[0] {
			tile[0] = false
			wasUpdated = true
		}
	}
	return wasUpdated
}

// Heuristic #3 - If NumUnsolvedTiles - 1 == RemainingVoltorbs, mark all tiles as either voltorbs or tiles with a value of RemainingPoints
func heuristic3(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.NumUnsolvedTiles-1 != lineData.RemainingVoltorbs {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			for j := range tile[1:] {
				j++
				if lineData.RemainingPoints != j && tile[j] {
					tile[j] = false
					wasUpdated = true
				}
			}
		}
	}
	return wasUpdated
}

// Heuristic #4 - If (NumUnsolvedTiles - RemainingVoltobs) <= ((RemainingPoints + 1)/3), eliminate 1 as a possibility from all tiles
func heuristic4(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if (lineData.NumUnsolvedTiles - lineData.RemainingVoltorbs) > ((lineData.RemainingPoints + 1) / 3) {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() && tile[1] {
			wasUpdated = true
			tile[1] = false
		}
	}
	return wasUpdated
}

// Heuristic #5 - If RemainingVoltorbs == 0 and RemainingPoints = NumUnsolvedTiles, mark all unsolved tiles as definitely 1s
func heuristic5(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if !((lineData.RemainingVoltorbs == 0) && (lineData.RemainingPoints == lineData.NumUnsolvedTiles)) {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			wasUpdated = true
			tile[0] = false
			tile[1] = true
			tile[2] = false
			tile[3] = false
		}
	}
	return wasUpdated
}

// Heuristic #6 - If RemainingVoltorbs == NumUnsolvedTiles, mark all unsolved tiles as definitely voltorbs
func heuristic6(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.RemainingVoltorbs != lineData.NumUnsolvedTiles {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			wasUpdated = true
			tile[0] = true
			tile[1] = false
			tile[2] = false
			tile[3] = false
		}
	}
	return wasUpdated
}

// Heuristic #7 - If NumUnsolvedTiles == 1, fill in the single unsolved tile using RemainingPoints
func heuristic7(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.NumUnsolvedTiles != 1 {
		return false
	}

	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			if !tile[lineData.RemainingPoints] {
				panic(fmt.Sprintf("Only one unsolved tile but value is not possible: %v , %v", lineData, tiles))
			}
			tile[0] = lineData.RemainingPoints == 0
			tile[1] = lineData.RemainingPoints == 1
			tile[2] = lineData.RemainingPoints == 2
			tile[3] = lineData.RemainingPoints == 3
			return true
		}
	}
	panic(fmt.Sprintf("Could not find any unsolved tiles: %v , %v", lineData, tiles))
}

// Heuristic #8 - If RemainingPoints == NumUnsolvedTiles*3, mark all unknowns as definitely 3s
func heuristic8(lineData vfPSolLineData, tiles *[5]VfPSolTile) bool {
	if lineData.RemainingPoints != lineData.NumUnsolvedTiles*3 {
		return false
	}

	wasUpdated := false
	for i := range tiles {
		tile := &tiles[i]
		if !tile.IsSolved() {
			wasUpdated = true
			tile[0] = false
			tile[1] = false
			tile[2] = false
			tile[3] = true
		}
	}
	return wasUpdated
}

package voltorbflip

import "fmt"

type VfPSolTile [4]bool

func NewSolvedVfPSolTile(points int) VfPSolTile {
	return VfPSolTile{0 == points, 1 == points, 2 == points, 3 == points}
}

func NewUnsolvedVfPSolTile() VfPSolTile {
	return VfPSolTile{true, true, true, true}
}

func (tile VfPSolTile) IsPossible() bool {
	return tile[0] || tile[1] || tile[2] || tile[3]
}

func (tile VfPSolTile) IsSolved() bool {
	numPossibilities := 0
	for _, x := range tile {
		if x {
			numPossibilities++
		}
	}
	return (numPossibilities == 1)
}

func (tile VfPSolTile) IsVoltorb() bool {
	if !tile.IsSolved() {
		panic(fmt.Sprintf("Can't get point value of unsolved tile %v", tile))
	}

	return tile[0]
}

func (tile VfPSolTile) IsPossiblyNecessaryToWin() bool {
	return tile[2] || tile[3]
}

func (tile VfPSolTile) Points() int {
	if !tile.IsSolved() {
		panic(fmt.Sprintf("Can't get point value of unsolved tile %v", tile))
	}

	for i, value := range tile {
		if value {
			return i
		}
	}

	panic(fmt.Sprintf("Can't get point value of unsolved tile %v", tile))
}

func (tile *VfPSolTile) SetPoints(value int) {
	tile[0] = value == 0
	tile[1] = value == 1
	tile[2] = value == 2
	tile[3] = value == 3
}

func (tile VfPSolTile) String() string {
	output := "T["
	if tile[0] {
		output += "V"
	}
	if tile[1] {
		output += "1"
	}
	if tile[2] {
		output += "2"
	}
	if tile[3] {
		output += "3"
	}
	output += "]"
	return output
}

type VfPSolBoard [5][5]VfPSolTile

func NewBlankVfPSolBoard() VfPSolBoard {
	result := VfPSolBoard{}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for i := 0; i < 4; i++ {
				result[x][y][i] = true
			}
		}
	}
	return result
}

type vfPSolLineData struct {
	RemainingPoints   int
	RemainingVoltorbs int
	NumUnsolvedTiles  int
}

func (lineData vfPSolLineData) IsPossible() bool {
	return ((lineData.RemainingPoints >= 0) &&
		(lineData.RemainingVoltorbs >= 0) &&
		(lineData.NumUnsolvedTiles >= 0) &&
		(lineData.RemainingVoltorbs <= lineData.NumUnsolvedTiles) &&
		((lineData.NumUnsolvedTiles-lineData.RemainingVoltorbs)*3 >= lineData.RemainingPoints))
}

func (lineData vfPSolLineData) String() string {
	return fmt.Sprintf("(RP=%d,RV=%d,#UT=%d)",
		lineData.RemainingPoints, lineData.RemainingVoltorbs, lineData.NumUnsolvedTiles)
}

type vfPartialSolution struct {
	board         *VfBoardTotals
	tiles         VfPSolBoard
	invertedTiles VfPSolBoard
	rowData       [5]vfPSolLineData
	columnData    [5]vfPSolLineData
}

func (psol vfPartialSolution) GetTile(position VfBoardPosition) VfPSolTile {
	return psol.tiles[position.Y][position.X]
}

func (psol vfPartialSolution) String() string {
	return fmt.Sprintf("(board=%v,tiles=%v,invertedTiles=%v,rowData=%v,columnData%v)",
		psol.board,
		psol.tiles,
		psol.invertedTiles,
		psol.rowData,
		psol.columnData,
	)
}

func newVfPartialSolution(board *VfBoardTotals, psolBoard VfPSolBoard) (bool, vfPartialSolution) {
	result := vfPartialSolution{board: board}

	result.tiles = psolBoard

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for i := 0; i < 4; i++ {
				result.invertedTiles[x][y][i] = result.tiles[y][x][i]
			}
		}
	}

	result.updateAllRowData()
	result.updateAllColumnData()

	isPossible := result.applyHeuristics()

	return isPossible, result
}

func (psol vfPartialSolution) IsSolved() bool {
	for _, rowData := range psol.rowData {
		if rowData.NumUnsolvedTiles != 0 {
			return false
		}
	}

	for _, columnData := range psol.columnData {
		if columnData.NumUnsolvedTiles != 0 {
			return false
		}
	}

	return true
}

func (psol vfPartialSolution) IsWon() bool {
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for i := 0; i < 4; i++ {
				if !psol.tiles[y][x].IsSolved() && psol.tiles[y][x].IsPossiblyNecessaryToWin() {
					return false
				}
			}
		}
	}
	return true
}

func (psol vfPartialSolution) IsPossible() bool {
	for _, rowData := range psol.rowData {
		if (rowData.RemainingVoltorbs > rowData.NumUnsolvedTiles) ||
			(rowData.RemainingVoltorbs < 0) ||
			(rowData.RemainingPoints > (rowData.NumUnsolvedTiles-rowData.RemainingVoltorbs)*3) ||
			(rowData.RemainingPoints < 0) {
			return false
		}
	}

	return true
}

func (psol vfPartialSolution) SolvedBoard() VfSolvedBoard {
	if !psol.IsSolved() {
		panic(fmt.Sprintf("Can't convert unsolved VfPartialSolution to VfSolvedBoard: %v", psol))
	}

	solvedBoard := [5][5]int{}
	for y, row := range psol.tiles {
		for x, tile := range row {
			solvedBoard[y][x] = tile.Points()
		}
	}
	return VfSolvedBoard(solvedBoard)
}

func (psol vfPartialSolution) UnsolvedTiles() map[VfBoardPosition]VfPSolTile {
	result := make(map[VfBoardPosition]VfPSolTile, 0)
	for y, row := range psol.tiles {
		for x, tile := range row {
			if !tile.IsSolved() {
				result[VfBoardPosition{x, y}] = tile
			}
		}
	}
	return result
}

func (psol *vfPartialSolution) Update(pos VfBoardPosition, tile VfPSolTile) bool {
	psol.tiles[pos.Y][pos.X] = tile
	if !psol.updateRowData(pos.Y) {
		return false
	}

	psol.invertedTiles[pos.X][pos.Y] = tile
	if !psol.updateColumnData(pos.X) {
		return false
	}

	isPossible := psol.applyHeuristics()
	return isPossible
}

func (psol vfPartialSolution) UpdatedPartialSolution(pos VfBoardPosition, tile VfPSolTile) (bool, vfPartialSolution) {
	isPossible := psol.Update(pos, tile)
	return isPossible, psol
}

func (psol *vfPartialSolution) updateRowData(y int) bool {
	remainingPoints := &psol.rowData[y].RemainingPoints
	remainingVoltorbs := &psol.rowData[y].RemainingVoltorbs
	numUnsolvedTiles := &psol.rowData[y].NumUnsolvedTiles

	*remainingPoints = psol.board.RowTotals[y].Points
	*remainingVoltorbs = psol.board.RowTotals[y].Voltorbs
	*numUnsolvedTiles = 0

	row := &psol.tiles[y]
	for _, tile := range row {
		if !tile.IsPossible() {
			return false
		} else if tile.IsSolved() {
			if tile.IsVoltorb() {
				*remainingVoltorbs--
			} else {
				*remainingPoints -= tile.Points()
			}
		} else {
			*numUnsolvedTiles++
		}
	}

	return psol.rowData[y].IsPossible()
}

func (psol *vfPartialSolution) updateAllRowData() bool {
	result := psol.updateRowData(0)
	result = result && psol.updateRowData(1)
	result = result && psol.updateRowData(2)
	result = result && psol.updateRowData(3)
	result = result && psol.updateRowData(4)
	return result
}

func (psol *vfPartialSolution) updateColumnData(x int) bool {
	remainingPoints := &psol.columnData[x].RemainingPoints
	remainingVoltorbs := &psol.columnData[x].RemainingVoltorbs
	numUnsolvedTiles := &psol.columnData[x].NumUnsolvedTiles

	*remainingPoints = psol.board.ColumnTotals[x].Points
	*remainingVoltorbs = psol.board.ColumnTotals[x].Voltorbs
	*numUnsolvedTiles = 0

	row := &psol.invertedTiles[x]
	for _, tile := range row {
		if !tile.IsPossible() {
			return false
		} else if tile.IsSolved() {
			if tile.IsVoltorb() {
				*remainingVoltorbs--
			} else {
				*remainingPoints -= tile.Points()
			}
		} else {
			*numUnsolvedTiles++
		}
	}

	return psol.columnData[x].IsPossible()
}

func (psol *vfPartialSolution) updateAllColumnData() bool {
	result := psol.updateColumnData(0)
	result = result && psol.updateColumnData(1)
	result = result && psol.updateColumnData(2)
	result = result && psol.updateColumnData(3)
	result = result && psol.updateColumnData(4)
	return result
}

func (psol *vfPartialSolution) updateInvertedTilesFromTiles() {
	updateTilesFromOtherTiles(&psol.tiles, &psol.invertedTiles)
}

func (psol *vfPartialSolution) updateTilesFromInvertedTiles() {
	updateTilesFromOtherTiles(&psol.invertedTiles, &psol.tiles)
}

func updateTilesFromOtherTiles(source *VfPSolBoard, destination *VfPSolBoard) {
	for i := range source {
		for j := range source[i] {
			destination[i][j] = source[j][i]
		}
	}
}

package voltorbflip

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVfPSolTile_IsSolved(t *testing.T) {
	assert.False(t, VfPSolTile{true, true, true, true}.IsSolved())

	assert.False(t, VfPSolTile{false, true, true, true}.IsSolved())
	assert.False(t, VfPSolTile{true, false, true, true}.IsSolved())
	assert.False(t, VfPSolTile{true, true, false, true}.IsSolved())
	assert.False(t, VfPSolTile{true, true, true, false}.IsSolved())

	assert.False(t, VfPSolTile{false, false, true, true}.IsSolved())
	assert.False(t, VfPSolTile{true, false, false, true}.IsSolved())
	assert.False(t, VfPSolTile{true, true, false, false}.IsSolved())

	assert.True(t, VfPSolTile{true, false, false, false}.IsSolved())
	assert.True(t, VfPSolTile{false, true, false, false}.IsSolved())
	assert.True(t, VfPSolTile{false, false, true, false}.IsSolved())
	assert.True(t, VfPSolTile{false, false, false, true}.IsSolved())
}

func TestVfPSolTile_IsVoltorb(t *testing.T) {
	assert.True(t, VfPSolTile{true, false, false, false}.IsVoltorb())
	assert.False(t, VfPSolTile{false, true, false, false}.IsVoltorb())
	assert.False(t, VfPSolTile{false, false, true, false}.IsVoltorb())
	assert.False(t, VfPSolTile{false, false, false, true}.IsVoltorb())
}

func TestVfPSolTile_Points(t *testing.T) {
	assert.Equal(t, 0, VfPSolTile{true, false, false, false}.Points())
	assert.Equal(t, 1, VfPSolTile{false, true, false, false}.Points())
	assert.Equal(t, 2, VfPSolTile{false, false, true, false}.Points())
	assert.Equal(t, 3, VfPSolTile{false, false, false, true}.Points())
}

func NewSolvedPsol() (vfPartialSolution, VfSolvedBoard) {
	board := VfBoardTotals{
		RowTotals:    [5]VfLineTotal{{7, 1}, {5, 2}, {4, 1}, {4, 1}, {5, 1}},
		ColumnTotals: [5]VfLineTotal{{5, 2}, {7, 2}, {5, 0}, {5, 0}, {3, 2}},
	}
	tiles := [5][5]VfPSolTile{
		[5]VfPSolTile{NewSolvedVfPSolTile(3), NewSolvedVfPSolTile(2), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(0)},
		[5]VfPSolTile{NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(3), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(0)},
		[5]VfPSolTile{NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1)},
		[5]VfPSolTile{NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1)},
		[5]VfPSolTile{NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(2), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1)},
	}
	psol := vfPartialSolution{
		board: &board,
		tiles: tiles,
	}
	psol.updateInvertedTilesFromTiles()

	expectedSolvedBoard := VfSolvedBoard{
		{3, 2, 1, 1, 0},
		{0, 3, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 1, 1},
		{0, 2, 1, 1, 1},
	}

	return psol, expectedSolvedBoard
}

func NewImpossiblePsol() vfPartialSolution {
	board := VfBoardTotals{
		RowTotals:    [5]VfLineTotal{{7, 9}, {5, 2}, {4, 1}, {4, 1}, {5, 1}},
		ColumnTotals: [5]VfLineTotal{{5, 2}, {7, 2}, {5, 0}, {5, 0}, {3, 2}},
	}
	_, psol := newVfPartialSolution(&board, NewBlankVfPSolBoard())

	return psol
}

func NewUnsolvedPsol() vfPartialSolution {
	board := VfBoardTotals{
		RowTotals:    [5]VfLineTotal{{5, 2}, {6, 1}, {6, 1}, {4, 1}, {4, 1}},
		ColumnTotals: [5]VfLineTotal{{3, 2}, {6, 1}, {6, 1}, {4, 2}, {6, 0}},
	}
	_, psol := newVfPartialSolution(&board, NewBlankVfPSolBoard())

	return psol
}

func NewPartiallySolvedPsol() vfPartialSolution {
	board := VfBoardTotals{
		RowTotals:    [5]VfLineTotal{{4, 2}, {7, 1}, {5, 0}, {3, 2}, {6, 1}},
		ColumnTotals: [5]VfLineTotal{{5, 2}, {5, 0}, {7, 1}, {5, 1}, {3, 2}},
	}
	tiles := [5][5]VfPSolTile{
		[5]VfPSolTile{NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(0), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()},
		[5]VfPSolTile{NewSolvedVfPSolTile(3), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(2), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()},
		[5]VfPSolTile{NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()},
		[5]VfPSolTile{NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()},
		[5]VfPSolTile{NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()},
	}
	psol := vfPartialSolution{
		board: &board,
		tiles: tiles,
	}
	psol.updateInvertedTilesFromTiles()

	return psol
}

func TestVfPartialSolution(t *testing.T) {
	Convey("Given a solved psol", t, func() {
		psol, expectedSolvedBoard := NewSolvedPsol()

		So(psol.IsPossible(), ShouldBeTrue)
		So(psol.IsSolved(), ShouldBeTrue)
		So(psol.SolvedBoard(), ShouldEqual, expectedSolvedBoard)
		So(psol.UnsolvedTiles(), ShouldBeEmpty)

		for _, rowData := range psol.rowData {
			So(rowData == vfPSolLineData{0, 0, 0}, ShouldBeTrue)
		}
	})

	Convey("Given an impossible psol", t, func() {
		psol := NewImpossiblePsol()

		So(psol.IsPossible(), ShouldBeFalse)
		So(psol.IsSolved(), ShouldBeFalse)
		So(func() { psol.SolvedBoard() }, ShouldPanic)
	})

	Convey("Given an unsolved psol", t, func() {
		psol := NewUnsolvedPsol()

		So(psol.IsPossible(), ShouldBeTrue)
		So(psol.IsSolved(), ShouldBeFalse)
		So(func() { psol.SolvedBoard() }, ShouldPanic)

		var unsolvedPositions []VfBoardPosition
		for k := range psol.UnsolvedTiles() {
			unsolvedPositions = append(unsolvedPositions, k)
		}

		So(unsolvedPositions, ShouldNotBeEmpty)
		So(unsolvedPositions, ShouldContain, VfBoardPosition{0, 0})
		So(unsolvedPositions, ShouldContain, VfBoardPosition{4, 0})
		So(unsolvedPositions, ShouldNotContain, VfBoardPosition{4, 4})
		So(unsolvedPositions, ShouldContain, VfBoardPosition{3, 1})

		_, psol2 := psol.UpdatedPartialSolution(VfBoardPosition{3, 1}, NewSolvedVfPSolTile(0))
		var unsolvedPositions2 []VfBoardPosition
		for k := range psol2.UnsolvedTiles() {
			unsolvedPositions2 = append(unsolvedPositions2, k)
		}
		So(unsolvedPositions2, ShouldNotContain, VfBoardPosition{3, 1})

		So(psol2.IsPossible(), ShouldBeTrue)
		So(psol2.IsSolved(), ShouldBeFalse)
		So(func() { psol2.SolvedBoard() }, ShouldPanic)

		So(psol.rowData[0] == vfPSolLineData{5, 2, 5}, ShouldBeTrue)
		So(psol.rowData[3] == vfPSolLineData{3, 1, 4}, ShouldBeTrue)
		So(psol.columnData[4] == vfPSolLineData{4, 0, 3}, ShouldBeTrue)
	})
}

func TestVfPartialSolution_UpdateRowData(t *testing.T) {
	Convey("Given a partially solved psol", t, func() {
		psol := NewPartiallySolvedPsol()

		psol.updateRowData(0)
		So(psol.rowData[0] == vfPSolLineData{2, 1, 2}, ShouldBeTrue)

		psol.tiles[0][3].SetPoints(2)
		psol.updateRowData(0)
		So(psol.rowData[0] == vfPSolLineData{0, 1, 1}, ShouldBeTrue)

		psol.tiles[0][4].SetPoints(0)
		psol.updateRowData(0)
		So(psol.rowData[0] == vfPSolLineData{0, 0, 0}, ShouldBeTrue)
	})
}

func TestVfPartialSolution_UpdateColumnData(t *testing.T) {
	Convey("Given a partially solved psol", t, func() {
		psol := NewPartiallySolvedPsol()

		psol.updateColumnData(0)
		So(psol.columnData[0] == vfPSolLineData{1, 2, 3}, ShouldBeTrue)

		psol.invertedTiles[0][2].SetPoints(1)
		psol.updateColumnData(0)
		So(psol.columnData[0] == vfPSolLineData{0, 2, 2}, ShouldBeTrue)

		psol.invertedTiles[0][3].SetPoints(0)
		psol.updateColumnData(0)
		So(psol.columnData[0] == vfPSolLineData{0, 1, 1}, ShouldBeTrue)
	})
}

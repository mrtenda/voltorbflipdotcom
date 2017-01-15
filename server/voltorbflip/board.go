package voltorbflip

import "fmt"

type VfLineTotal struct {
	Points   int
	Voltorbs int
}

func (lineTotal VfLineTotal) String() string {
	return fmt.Sprintf("(P:%d,V:%d)", lineTotal.Points, lineTotal.Voltorbs)
}

type VfBoardTotals struct {
	RowTotals    [5]VfLineTotal
	ColumnTotals [5]VfLineTotal
}

func (board VfBoardTotals) String() string {
	return fmt.Sprintf("(RowTotals:%v,ColumnTotals:%v)", board.RowTotals, board.ColumnTotals)
}

type VfSolvedBoard [5][5]int

type VfBoardPosition struct {
	X int
	Y int
}

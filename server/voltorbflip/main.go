package voltorbflip

func Solve(boardTotals *VfBoardTotals, partiallySolvedBoard VfPSolBoard) (bool, bool, VfPSolBoard, VfBoardPosition, float32) {
	isPossible, psol := newVfPartialSolution(boardTotals, partiallySolvedBoard)
	if !isPossible {
		return false, false, psol.tiles, VfBoardPosition{}, 0
	}
	if psol.IsWon() {
		return true, true, psol.tiles, VfBoardPosition{}, 0
	}
	isPossible, isWon, safestPosition, safety := psol.SafestUnsolvedPosition()
	return isPossible, isWon, psol.tiles, safestPosition, safety
}

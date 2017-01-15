package voltorbflip

import (
	"fmt"
)

func (psol vfPartialSolution) AllPossibleSolvedBoards() (map[VfSolvedBoard]int, int) {
	if psol.IsSolved() {
		panic(fmt.Sprintf("Can't get solved boards for already solved psol: %v", psol))
	}

	psolStack := []vfPartialSolution{psol}
	possibleSolvedBoardCounts := make(map[VfSolvedBoard]int)
	numberOfPossibilities := 0
	var currentPsol vfPartialSolution

	for len(psolStack) != 0 {
		currentPsol, psolStack = psolStack[len(psolStack)-1], psolStack[:len(psolStack)-1]

	OUTER:
		for y, row := range currentPsol.tiles {
			for x, tile := range row {
				if !tile.IsSolved() {
					for value, isValuePossible := range tile {
						if isValuePossible {
							isPossible, newPsol := currentPsol.UpdatedPartialSolution(
								VfBoardPosition{x, y}, NewSolvedVfPSolTile(value))
							if !isPossible || !newPsol.IsPossible() {
								continue
							} else if newPsol.IsSolved() {
								possibleSolvedBoardCounts[newPsol.SolvedBoard()]++
								numberOfPossibilities++
							} else {
								psolStack = append(psolStack, newPsol)
							}
						}
					}
					break OUTER
				}
			}
		}
	}

	return possibleSolvedBoardCounts, numberOfPossibilities
}

func (psol vfPartialSolution) SafetyOfEachGuess() (map[VfBoardPosition]int, int) {
	allPossibleSolvedBoards, numberOfPossibilities := psol.AllPossibleSolvedBoards()

	result := make(map[VfBoardPosition]int)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !psol.tiles[y][x].IsSolved() {
				safeCount := 0
				for board, count := range allPossibleSolvedBoards {
					if board[y][x] != 0 {
						safeCount += count
					}
				}
				result[VfBoardPosition{x, y}] += safeCount
			}
		}
	}

	return result, numberOfPossibilities
}

func (psol vfPartialSolution) SafestUnsolvedPosition() (bool, bool, VfBoardPosition, float32) {
	if psol.IsWon() {
		return true, true, VfBoardPosition{}, 0
	}

	unsolvedPositions := psol.UnsolvedTiles()
	for position, tile := range unsolvedPositions {
		if !tile[0] {
			return true, false, position, 1
		}
	}

	positionSafeties, numberOfPossibilities := psol.SafetyOfEachGuess()

	var safestPosition VfBoardPosition
	safestNumberOfPossibilities := -1
	for position, safePossibilities := range positionSafeties {
		if safePossibilities == numberOfPossibilities {
			tile := psol.GetTile(position)
			tile[0] = false
			if !tile.IsPossible() {
				return false, false, VfBoardPosition{}, 0
			}
			isPossible := psol.Update(position, tile)
			if !isPossible {
				return false, false, VfBoardPosition{}, 0
			}
			return psol.SafestUnsolvedPosition()
		} else if safePossibilities > safestNumberOfPossibilities {
			safestPosition = position
			safestNumberOfPossibilities = safePossibilities
		}
	}

	return true, false, safestPosition, (float32(safestNumberOfPossibilities) / float32(numberOfPossibilities))
}

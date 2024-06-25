package kata

const draw = 0
const empty = 0
const notFinished = -1

// IsSolved / https://www.codewars.com/kata/525caa5c1bf619d28c000335/train/go
func IsSolved(board [3][3]int) int {
	myBoard := NewBoard(board)
	winner := myBoard.Winner()

	if winner != draw {
		return winner
	}

	if myBoard.IsFinished() {
		return draw
	}

	return notFinished
}

type Board struct {
	rows  [3]rowColumnOrDiagonal
	cols  [3]rowColumnOrDiagonal
	diags [2]rowColumnOrDiagonal
}

func NewBoard(board [3][3]int) Board {
	return Board{
		rows:  extractRows(board),
		cols:  extractCols(board),
		diags: extractDiags(board),
	}
}

func (b Board) Winner() int {
	if winner := b.rowWinner(); winner != draw {
		return winner
	}

	if winner := b.colWinner(); winner != draw {
		return winner
	}

	if winner := b.diagWinner(); winner != draw {
		return winner
	}

	return draw
}

func (b Board) IsFinished() bool {
	for i := 0; i < 3; i++ {
		if !b.rows[i].isFinished() {
			return false
		}
	}

	return true
}

func (b Board) rowWinner() int {
	for i := 0; i < 3; i++ {
		winner := b.rows[i].winner()

		if winner != empty {
			return winner
		}
	}

	return draw
}

func (b Board) colWinner() int {
	for i := 0; i < 3; i++ {
		winner := b.cols[i].winner()

		if winner != empty {
			return winner
		}
	}

	return draw
}

func (b Board) diagWinner() int {
	for i := 0; i < 2; i++ {
		winner := b.diags[i].winner()

		if winner != empty {
			return winner
		}
	}

	return draw
}

func extractRows(board [3][3]int) [3]rowColumnOrDiagonal {
	rows := [3]rowColumnOrDiagonal{}

	for i := 0; i < 3; i++ {
		rows[i] = newRowColumnOrDiagonal(board[i])
	}

	return rows
}

func extractCols(board [3][3]int) [3]rowColumnOrDiagonal {
	cols := [3]rowColumnOrDiagonal{}

	for i := 0; i < 3; i++ {
		cols[i] = newRowColumnOrDiagonal([3]int{board[0][i], board[1][i], board[2][i]})
	}

	return cols
}

func extractDiags(board [3][3]int) [2]rowColumnOrDiagonal {
	diags := [2]rowColumnOrDiagonal{}

	for i := 0; i < 3; i++ {
		diags[0] = newRowColumnOrDiagonal([3]int{board[0][0], board[1][1], board[2][2]})
		diags[1] = newRowColumnOrDiagonal([3]int{board[0][2], board[1][1], board[2][0]})
	}

	return diags
}

type rowColumnOrDiagonal struct {
	fields [3]int
}

func newRowColumnOrDiagonal(fields [3]int) rowColumnOrDiagonal {
	return rowColumnOrDiagonal{
		fields: fields,
	}
}

func (r rowColumnOrDiagonal) winner() int {
	if r.fields[0] == r.fields[1] && r.fields[1] == r.fields[2] {
		return r.fields[0]
	}

	return draw
}

func (r rowColumnOrDiagonal) isFinished() bool {
	return r.fields[0] != 0 && r.fields[1] != 0 && r.fields[2] != 0
}

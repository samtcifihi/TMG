package goban

import ()

// sq represents a rectangular goban
type sq struct {
	/*
		-1 = White
		0  = Neutral
		1  = Black
	*/
	board [][]int8
}

// NewSq adds a new pyramid to the PyrPile
func NewSq(i uint, j uint) *sq {
	newSq := new(sq)

	m, n := i, j
	for ; m > 0; m-- {
		for ; n > 0; n-- {
			newSq.board[m-1][n-1] = 0
		}
	}

	return newSq
}

// GetPoint returns the value of the specified point
func (s *sq) GetPoint(i uint, j uint) int8 {
	return s.board[i][j]
}

// ColorPoint colors a specified point the specified color
func (s *sq) ColorPoint(i uint, j uint, color int8) {
	s.board[i][j] = color
}

// ClearColor removes all groups of the specified color with no liberties
func (s *sq) ClearColor(color int8) {
	if color == -1 {
		// Clear White Stones
	} else {
		// Clear Black Stones
	}
}

// Size returns the height and width of the goban.sq respectively
func (s *sq) Size() (uint, uint) {
	return uint(len(s.board)), uint(len(s.board[0]))
}

// Row returns the specified row from the sq
func (s *sq) Row(i uint) []int8 {
	return s.board[i]
}

// Score returns the score of the game (+ for B, - for W)
func (s *sq) Score(komi int) int {
	/*
		For every neutral point, color
		* -1 if can only reach W,
		* 0 if can reach both B and W,
		* 1 if can only reach B
	*/

	score := 0
	// Add Black points to score
	// Subtract White points from score
	// Subtract komi from score
	return score
}

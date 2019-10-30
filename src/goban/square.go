package goban

import (
	"fmt"
)

// sq represents a rectangular goban
type sq struct {
	/*
		-1 = White
		0  = Neutral
		1  = Black
	*/
	board [][]int8
}

// NewSq constructs an new sq
func NewSq(i uint, j uint) *sq {
	newSq := new(sq)

	i, j = i-1, j-1

	row := []int8{}
	for n := 0; uint(n) <= j; n++ {
		row = append(row, 0)
	}

	for m := 0; uint(m) <= i; m++ {
		// for n := 0; uint(n) <= j; n++ {
		// newSq.board[m] = append(newSq.board[m], 0)
		// }

		newSq.board = append(newSq.board, row)
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

// Height gives Height
func (s *sq) Height() uint {
	return uint(len(s.board))
}

// Width gives width
func (s *sq) Width() uint {
	return uint(len(s.board[0]))
}

// Row returns the specified row from the sq
func (s *sq) Row(i uint) []int8 {
	return s.board[i]
}

// Score returns the score of the game (+ for B, - for W)
func (s *sq) Score(komi float64) float64 {
	score := 0.0

	/*
		For every neutral point, score
		* -1 if can only reach W,
		* 0 if can reach both B and W,
		* 1 if can only reach B
	*/
	for i := range s.board {
		for j := range s.board[i] {
			if s.GetPoint(uint(i), uint(j)) == -1 {
				score = score - 1
			} else if s.GetPoint(uint(i), uint(j)) == 1 {
				score = score + 1
			} else {
				if s.CanReach(uint(i), uint(j), true, -1) {
					score = score - 1
				} else if s.CanReach(uint(i), uint(j), true, 1) {
					score = score + 1
				}
			}
		}
	}

	return score - komi
}

// CanReach checks which colors a point can reach
func (s *sq) CanReach(i uint, j uint, isExclusive bool, colors ...int8) bool {
	pointColor := s.GetPoint(i, j)
	var out bool

	if isExclusive == false {
		// return "out = true" upon reaching a color in "colors"
		// else return "out = false"
		out = false
		for _, k := range colors {
			out = out || adjColors(s, i, j, pointColor, k)
		}

	} else {
		// return "out = false" upon failing to reach a color in "colors"
		// return "out = false" upon reaching a color not in "colors"
		// else return "out = true"

		out = true
		var check bool

		for k := -1; k <= 1; k-- {
			check = false
			for _, l := range colors {
				if l == int8(k) {
					check = true
				}
			}
			if check == true {
				out = out && adjColors(s, i, j, pointColor, int8(k))
			} else {
				out = out && !(adjColors(s, i, j, pointColor, int8(k)))
			}
		}
	}
	return out
}

// adjColors is a recursive function to check if a group can reach a given color
func adjColors(s *sq, i uint, j uint, currentColor int8, targetColor int8) bool {
	foundColor := false
	m, n := s.Size()

	if s.GetPoint(i, j) == targetColor {
		foundColor = true
	} else if s.GetPoint(i, j) == currentColor {
		if i > 0 {
			foundColor = adjColors(s, i-1, j, currentColor, targetColor)
		}
		if j > 0 {
			foundColor = foundColor || adjColors(s, i, j-1, currentColor, targetColor)
		}
		if i < m {
			foundColor = foundColor || adjColors(s, i+1, j, currentColor, targetColor)
		}
		if j < n {
			foundColor = foundColor || adjColors(s, i, j+1, currentColor, targetColor)
		}
	}

	return foundColor
}

// Print outputs the current game state
func (s *sq) Print() {
	for i := range s.board {
		fmt.Println(s.Row(uint(i)))
	}
}

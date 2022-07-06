package sudokutype

import (
	"math"
)

const (
	SizeCommon = 9
	SizeBig    = 16
	SizeHuge   = 25
)

type (
	// SudokuCommon represents the common 9x9 sudoku.
	SudokuCommon [SizeCommon][SizeCommon]int8
	// SudokuBig represents a big 16x16 sudoku.
	SudokuBig [SizeBig][SizeBig]int8
	// SudokuHuge represents a huge 25x25 sudoku.
	SudokuHuge [SizeHuge][SizeHuge]int8
)

// Sudoku is the type summ for all the 3 sudoku varieties.
type Sudoku interface {
	*SudokuCommon | *SudokuBig | *SudokuHuge
	Get(y, x int) int8
	Set(y, x int, v int8)
}

func (s *SudokuCommon) Get(y, x int) int8    { return s[y][x] }
func (s *SudokuCommon) Set(y, x int, v int8) { s[y][x] = v }
func (s *SudokuBig) Get(y, x int) int8       { return s[y][x] }
func (s *SudokuBig) Set(y, x int, v int8)    { s[y][x] = v }
func (s *SudokuHuge) Get(y, x int) int8      { return s[y][x] }
func (s *SudokuHuge) Set(y, x int, v int8)   { s[y][x] = v }

// BaseFill is making a base valid sudoku of the given type.
func BaseFill[S Sudoku](s S) {
	sq := int(math.Sqrt(float64(len(s))))
	for sy := 0; sy < sq; sy++ {
		for y := 0; y < sq; y++ {
			for x := 0; x < len(s); x++ {
				s.Set(sy*sq+y, x, int8((x+sy+y*sq)%len(s)+1))
			}
		}
	}
}

// Copy duplicate the one sudoku to another.
func Copy[S Sudoku](dst, src S) {
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src); x++ {
			dst.Set(y, x, src.Get(y, x))
		}
	}
}

// Equal compares 2 same type sudoku.
func Equal[S Sudoku](dst, src S) bool {
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src); x++ {
			if dst.Get(y, x) != src.Get(y, x) {
				return false
			}
		}
	}

	return true
}

// EqualPartial compares 2 same type sudoku.
// Some cells in src could be 0, so they will be ignored.
func EqualPartial[S Sudoku](dst, src S) bool {
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src); x++ {
			v := src.Get(y, x)
			if v == 0 {
				continue
			}

			if dst.Get(y, x) != v {
				return false
			}
		}
	}

	return true
}

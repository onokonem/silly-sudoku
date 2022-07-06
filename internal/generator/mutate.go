package generator

import (
	"math"
	"math/rand"

	"silly-sudoku/internal/sudokutype"
)

//nolint:gomnd // false positives
func Randomize[S sudokutype.Sudoku](s S, rnd *rand.Rand) {
	for i := 0; i < 100; i++ {
		switch rnd.Intn(4) {
		case 0:
			mirror(s)
		case 1:
			rotate(s)
		case 2:
			l1, l2 := randLines(s, rnd)
			swapRows(s, l1, l2)
		case 3:
			l1, l2 := randLines(s, rnd)
			swapColumns(s, l1, l2)
		}
	}
}

func randLines[S sudokutype.Sudoku](s S, rnd *rand.Rand) (int, int) {
	var (
		sq     = int(math.Sqrt(float64(len(s))))
		region = rnd.Intn(sq)
	)

	return region*sq + rnd.Intn(sq), region*sq + rnd.Intn(sq)
}

func mirror[S sudokutype.Sudoku](s S) {
	for y := 0; y < len(s); y++ {
		for x := y; x < len(s); x++ {
			v := s.Get(x, y)
			s.Set(x, y, s.Get(y, x))
			s.Set(y, x, v)
		}
	}
}

func rotate[S sudokutype.Sudoku](s S) {
	l := len(s)

	for y := 0; y < l/2; y++ {
		for x := y; x < l-1-y; x++ {
			v1 := s.Get(l-1-x, y)
			v2 := s.Get(y, x)
			v3 := s.Get(x, l-1-y)
			v4 := s.Get(l-1-y, l-1-x)

			s.Set(l-1-x, y, v4)
			s.Set(y, x, v1)
			s.Set(x, l-1-y, v2)
			s.Set(l-1-y, l-1-x, v3)
		}
	}
}

func swapRows[S sudokutype.Sudoku](s S, row1, row2 int) {
	for x := 0; x < len(s); x++ {
		v := s.Get(row1, x)
		s.Set(row1, x, s.Get(row2, x))
		s.Set(row2, x, v)
	}
}

func swapColumns[S sudokutype.Sudoku](s S, col1, col2 int) {
	for y := 0; y < len(s); y++ {
		v := s.Get(y, col1)
		s.Set(y, col1, s.Get(y, col2))
		s.Set(y, col2, v)
	}
}

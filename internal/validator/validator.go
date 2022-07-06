package validator

import (
	"errors"
	"fmt"
	"math"

	"silly-sudoku/internal/sudokutype"
)

var ErrInvalidValue = errors.New("invalid value")

// Validate checks the sudoku for the these rules:
// numbers must be unique for each row, each column and each region.
func Validate[S sudokutype.Sudoku](s S) (bool, error) {
	if err := checkValues(s, 1); err != nil {
		return false, fmt.Errorf("validating: %w", err)
	}

	return checkRows(s) && checkColumns(s) && checkSquares(s), nil
}

// ValidatePartial checks the sudoku for the these rules:
// numbers must be unique for each row, each column and each region.
// Unfilled (0) cells are ignored.
func ValidatePartial[S sudokutype.Sudoku](s S) (bool, error) {
	if err := checkValues(s, 0); err != nil {
		return false, fmt.Errorf("validating: %w", err)
	}

	return checkRows(s) && checkColumns(s) && checkSquares(s), nil
}

func checkValues[S sudokutype.Sudoku](s S, min int8) error {
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s); x++ {
			if v := s.Get(y, x); v < min || int(v) > len(s) {
				return fmt.Errorf("%d (%d:%d): %w", v, x, y, ErrInvalidValue)
			}
		}
	}

	return nil
}

func checkRows[S sudokutype.Sudoku](s S) bool {
	for y := 0; y < len(s); y++ {
		uniq := make([]bool, len(s))

		for x := 0; x < len(s); x++ {
			v := s.Get(y, x)
			if v == 0 {
				continue
			}

			if uniq[v-1] {
				return false
			}

			uniq[v-1] = true
		}
	}

	return true
}

func checkColumns[S sudokutype.Sudoku](s S) bool {
	for x := 0; x < len(s); x++ {
		uniq := make([]bool, len(s))

		for y := 0; y < len(s); y++ {
			v := s.Get(y, x)
			if v == 0 {
				continue
			}

			if uniq[v-1] {
				return false
			}

			uniq[v-1] = true
		}
	}

	return true
}

func checkSquares[S sudokutype.Sudoku](s S) bool {
	sq := int(math.Sqrt(float64(len(s))))

	for sx := 0; sx < sq; sx++ {
		for sy := 0; sy < sq; sy++ {
			uniq := make([]bool, len(s))

			for x := sx * sq; x < sx*sq+sq; x++ {
				for y := sy * sq; y < sy*sq+sq; y++ {
					v := s.Get(y, x)
					if v == 0 {
						continue
					}

					if uniq[v-1] {
						return false
					}

					uniq[v-1] = true
				}
			}
		}
	}

	return true
}

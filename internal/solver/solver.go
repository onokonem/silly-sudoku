package solver

import (
	"errors"
	"fmt"

	"silly-sudoku/internal/sudokutype"
	"silly-sudoku/internal/validator"
)

// ErrInvalidField is an error returned in case the initial field is invalid.
var ErrInvalidField = errors.New("invalid field")

// Solve is a simple backtrack sudoku solver.
// Returns true if sudoku is solvable,
// solution is written to the original sudoku in place.
// Returns false if sudoku is unsolvable.
// Return error in case provided field is invalid.
func Solve[S sudokutype.Sudoku](s S) (bool, error) {
	// initial check
	ok, err := validator.ValidatePartial(s)
	if err != nil {
		return false, fmt.Errorf("invalid sudoku: %w", err)
	}

	if !ok {
		return false, fmt.Errorf("invalid sudoku: %w", ErrInvalidField)
	}

	return solve(s), nil
}

func solve[S sudokutype.Sudoku](s S) bool {
	undone := false

	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s); x++ {
			if s.Get(y, x) == 0 {
				undone = true //nolint:ineffassign // false positive

				for v := len(s); v > 0; v-- {
					s.Set(y, x, int8(v))

					if !validatePartialMust(s) {
						s.Set(y, x, 0)

						continue
					}

					if solve(s) {
						return true
					}
				}

				return false
			}
		}
	}

	return !undone
}

func validatePartialMust[S sudokutype.Sudoku](s S) bool {
	ok, err := validator.ValidatePartial(s)
	if err != nil {
		panic(err)
	}

	return ok
}

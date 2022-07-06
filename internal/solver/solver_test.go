package solver_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/solver"
	"silly-sudoku/internal/sudokutype"
	"silly-sudoku/internal/validator"
)

//nolint:gochecknoglobals // it is ok to have the global variables in tests
var (
	good = &sudokutype.SudokuCommon{
		{1, 2, 0, 4, 5, 6, 7, 8, 9},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{7, 8, 0, 1, 2, 3, 4, 5, 6},
		{2, 3, 0, 5, 6, 7, 8, 9, 1},
		{5, 6, 0, 8, 9, 1, 2, 3, 4},
		{8, 9, 0, 2, 3, 4, 5, 6, 7},
		{3, 4, 0, 6, 7, 8, 9, 1, 2},
		{6, 7, 0, 9, 1, 2, 3, 4, 5},
		{9, 1, 0, 3, 4, 5, 6, 7, 8},
	}
	bad = &sudokutype.SudokuCommon{
		{1, 2, 0, 4, 5, 6, 7, 8, 9},
		{0, 0, 3, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	ugly1 = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	ugly2 = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 10},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
)

func TestSolverGood(t *testing.T) {
	t.Parallel()

	ok, err := solver.Solve(good)
	require.NoError(t, err, "solving good")
	require.True(t, ok, "solving good")
}

func TestSolverBad(t *testing.T) {
	t.Parallel()

	ok, err := solver.Solve(bad)
	require.NoError(t, err, "solving bad")
	require.False(t, ok, "solving bad")
}

func TestSolverUgly(t *testing.T) {
	t.Parallel()

	ok, err := solver.Solve(ugly1)
	require.ErrorIs(t, err, solver.ErrInvalidField, "solving ugly 1")
	require.False(t, ok, "solving ugly 1")

	ok, err = solver.Solve(ugly2)
	require.ErrorIs(t, err, validator.ErrInvalidValue, "solving ugly 2")
	require.False(t, ok, "solving ugly 2")
}

func TestSolverEmpty(t *testing.T) {
	t.Parallel()

	// SudokuHuge test will be way too long

	ok, err := solver.Solve(&sudokutype.SudokuCommon{})
	require.NoError(t, err, "solving empty common")
	require.True(t, ok, "solving empty common")

	ok, err = solver.Solve(&sudokutype.SudokuBig{})
	require.NoError(t, err, "solving empty big")
	require.True(t, ok, "solving empty big")
}

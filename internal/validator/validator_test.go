package validator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/sudokutype"
	"silly-sudoku/internal/validator"
)

//nolint:gochecknoglobals // it is ok to have the global variables in tests
var (
	good = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	goodPartial = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 0, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	badRow = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 1, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	badCol = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 1},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 4, 2, 3, 1, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	badSquare = &sudokutype.SudokuCommon{
		{1, 6, 3, 4, 5, 2, 7, 8, 9},
		{4, 9, 6, 7, 8, 5, 1, 2, 3},
		{7, 3, 9, 1, 2, 8, 4, 5, 6},
		{2, 7, 4, 5, 6, 3, 8, 9, 1},
		{5, 1, 7, 8, 9, 6, 2, 3, 4},
		{8, 4, 1, 2, 3, 9, 5, 6, 7},
		{3, 8, 5, 6, 7, 4, 9, 1, 2},
		{6, 2, 8, 9, 1, 7, 3, 4, 5},
		{9, 5, 2, 3, 4, 1, 6, 7, 8},
	}
	ugly = &sudokutype.SudokuCommon{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7, 8, 9, 10},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
)

func TestValidateGood(t *testing.T) {
	t.Parallel()

	ok, err := validator.Validate(good)
	require.NoError(t, err, "validating good")
	require.True(t, ok, "validating good")

	ok, err = validator.ValidatePartial(good)
	require.NoError(t, err, "validating good")
	require.True(t, ok, "validating good")

	ok, err = validator.Validate(goodPartial)
	require.ErrorIs(t, err, validator.ErrInvalidValue, "validating partially good")
	require.False(t, ok, "validating partially good")

	ok, err = validator.ValidatePartial(goodPartial)
	require.NoError(t, err, "validating partially good")
	require.True(t, ok, "validating partially good")
}

func TestValidateBad(t *testing.T) {
	t.Parallel()

	ok, err := validator.Validate(badRow)
	require.NoError(t, err, "validating bad row")
	require.False(t, ok, "validating bad row")

	ok, err = validator.ValidatePartial(badRow)
	require.NoError(t, err, "validating bad row")
	require.False(t, ok, "validating bad row")

	ok, err = validator.Validate(badCol)
	require.NoError(t, err, "validating bad column")
	require.False(t, ok, "validating bad column")

	ok, err = validator.ValidatePartial(badCol)
	require.NoError(t, err, "validating bad column")
	require.False(t, ok, "validating bad column")

	ok, err = validator.Validate(badSquare)
	require.NoError(t, err, "validating bad square")
	require.False(t, ok, "validating bad square")

	ok, err = validator.ValidatePartial(badSquare)
	require.NoError(t, err, "validating bad square")
	require.False(t, ok, "validating bad square")
}

func TestValidateUgly(t *testing.T) {
	t.Parallel()

	ok, err := validator.Validate(ugly)
	require.ErrorIs(t, err, validator.ErrInvalidValue, "validating ugly")
	require.False(t, ok, "validating ugly")

	ok, err = validator.ValidatePartial(ugly)
	require.ErrorIs(t, err, validator.ErrInvalidValue, "validating ugly")
	require.False(t, ok, "validating ugly")
}

package sudokutype_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/sudokutype"
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
)

func TestBase(t *testing.T) {
	t.Parallel()

	s := &sudokutype.SudokuCommon{}
	sudokutype.BaseFill(s)

	require.True(t, reflect.DeepEqual(good, s), "check equality after base filling")
}

func TestCopy(t *testing.T) {
	t.Parallel()

	src := &sudokutype.SudokuCommon{}
	dst := &sudokutype.SudokuCommon{}

	src.Set(8, 8, 9)

	require.False(t, reflect.DeepEqual(dst, src), "check equality before copy")

	sudokutype.Copy(dst, src)

	require.True(t, reflect.DeepEqual(dst, src), "check equality after copy")
}

func TestEqual(t *testing.T) {
	t.Parallel()

	src := &sudokutype.SudokuCommon{}
	dst := &sudokutype.SudokuCommon{}

	src.Set(8, 8, 9)

	require.False(t, sudokutype.Equal(dst, src), "check equality before copy")

	sudokutype.Copy(dst, src)

	require.True(t, sudokutype.Equal(dst, src), "check equality after copy")
}

func TestEqualPartial(t *testing.T) {
	t.Parallel()

	src := &sudokutype.SudokuCommon{}
	dst := &sudokutype.SudokuCommon{}

	src.Set(8, 8, 9)
	src.Set(7, 7, 9)

	require.False(t, sudokutype.EqualPartial(dst, src), "check equality before copy")

	sudokutype.Copy(dst, src)

	require.True(t, sudokutype.EqualPartial(dst, src), "check equality after copy")

	src.Set(7, 7, 0)

	require.False(t, sudokutype.Equal(dst, src), "check equality after suppressing")
	require.True(t, sudokutype.EqualPartial(dst, src), "check partial equality after suppressing")
}

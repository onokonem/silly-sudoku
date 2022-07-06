package generator_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/generator"
	"silly-sudoku/internal/sudokutype"
	"silly-sudoku/internal/validator"
)

func TestMutateCommon(t *testing.T) {
	t.Parallel()

	s := &sudokutype.SudokuCommon{}
	sudokutype.BaseFill(s)

	generator.Randomize(s, rand.New(rand.NewSource(time.Now().UnixNano()))) //nolint:gosec // math/rand is ok

	ok, err := validator.Validate(s)
	require.NoError(t, err, "randomising common")
	require.True(t, ok, "randomising common")
}

func TestMutateBig(t *testing.T) {
	t.Parallel()

	s := &sudokutype.SudokuBig{}
	sudokutype.BaseFill(s)

	generator.Randomize(s, rand.New(rand.NewSource(time.Now().UnixNano()))) //nolint:gosec // math/rand is ok

	ok, err := validator.Validate(s)
	require.NoError(t, err, "randomising big")
	require.True(t, ok, "randomising big")
}

func TestMutateHuge(t *testing.T) {
	t.Parallel()

	s := &sudokutype.SudokuHuge{}
	sudokutype.BaseFill(s)

	generator.Randomize(s, rand.New(rand.NewSource(time.Now().UnixNano()))) //nolint:gosec // math/rand is ok

	ok, err := validator.Validate(s)
	require.NoError(t, err, "randomising huge")
	require.True(t, ok, "randomising huge")
}

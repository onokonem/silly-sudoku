package generator_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/generator"
	"silly-sudoku/internal/sudokutype"
)

const difficultyToTest float64 = 25.0 / 81.0

func TestSuppress(t *testing.T) {
	t.Parallel()

	original := &sudokutype.SudokuCommon{}

	generator.Generate(original)

	field := generator.Suppress(
		original,
		&sudokutype.SudokuCommon{},
		&sudokutype.SudokuCommon{},
		difficultyToTest,
	)

	require.False(t, sudokutype.Equal(original, field), "suppressing")
	require.True(t, sudokutype.EqualPartial(original, field), "suppressing")

	nonZero := 0

	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field); x++ {
			if field.Get(y, x) > 0 {
				nonZero++
			}
		}
	}

	require.Greater(t, nonZero, 0, "suppressing")
	require.Less(t, nonZero, len(field)*len(field), "suppressing")
}

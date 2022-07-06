package encdec_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/generator"
	"silly-sudoku/internal/handlers/internal/encdec"
	"silly-sudoku/internal/sudokutype"
)

func TestEncDec(t *testing.T) {
	t.Parallel()

	var (
		original     = &sudokutype.SudokuCommon{}
		fieldBack    = &sudokutype.SudokuCommon{}
		originalBack = &sudokutype.SudokuCommon{}
	)

	generator.Generate(original)

	field := generator.Suppress(
		original,
		&sudokutype.SudokuCommon{},
		&sudokutype.SudokuCommon{},
		0.5,
	)

	b, err := encdec.Enc(field, original)
	require.NoError(t, err, "encoding")

	err = encdec.Dec(b, fieldBack, originalBack)
	require.NoError(t, err, "decoding")
	require.Equal(t, field, fieldBack, "decoding")
	require.Equal(t, original, originalBack, "decoding")
}

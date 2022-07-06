package handlers_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/require"

	"silly-sudoku/internal/generator"
	"silly-sudoku/internal/handlers"
	"silly-sudoku/internal/oapi/models"
	"silly-sudoku/internal/oapi/restapi/operations/check"
	"silly-sudoku/internal/oapi/restapi/operations/get"
	"silly-sudoku/internal/solver"
	"silly-sudoku/internal/sudokutype"
)

func TestGetAndCheckOk(t *testing.T) {
	t.Parallel()

	field, original := getSudokuSolved(t, &sudokutype.SudokuCommon{})

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    handlers.FieldInternalToAPI(field),
				Original: original,
			},
		},
	)

	_, ok := checkRaw.(*check.CheckSudokuOK)
	require.True(t, ok, "checking sudoku")
}

func TestGetAndCheckBadField(t *testing.T) {
	t.Parallel()

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    models.Field{},
				Original: nil,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking bad field sudoku")
	require.Equal(t, int32(702), *checkRes.Payload.Code, "checking bad field sudoku")
}

func TestGetAndCheckBadRow(t *testing.T) {
	t.Parallel()

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    make(models.Field, sudokutype.SizeCommon),
				Original: nil,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking bad row sudoku")
	require.Equal(t, int32(702), *checkRes.Payload.Code, "checking bad row sudoku")
}

func TestGetAndCheckIncomplete(t *testing.T) {
	t.Parallel()

	field, original := getSudokuSolved(t, &sudokutype.SudokuCommon{})

	field.Set(0, 0, 0)

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    handlers.FieldInternalToAPI(field),
				Original: original,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking incomplete solution")
	require.Equal(t, int32(702), *checkRes.Payload.Code, "checking incomplete solution")
}

func TestGetAndCheckUnsolved(t *testing.T) {
	t.Parallel()

	field, original := getSudokuSolved(t, &sudokutype.SudokuCommon{})

	field.Set(0, 0, field.Get(0, 1))

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    handlers.FieldInternalToAPI(field),
				Original: original,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking bad solution")
	require.Equal(t, int32(702), *checkRes.Payload.Code, "checking bad solution")
}

func TestGetAndCheckDamaged(t *testing.T) {
	t.Parallel()

	field, original := getSudokuSolved(t, &sudokutype.SudokuCommon{})

	original = (*models.Original)(pointy.String(string(*original)[1:]))

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    handlers.FieldInternalToAPI(field),
				Original: original,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking bad original")
	require.Equal(t, int32(701), *checkRes.Payload.Code, "checking bad original")
}

func TestGetAndCheckIncompatible(t *testing.T) {
	t.Parallel()

	field, original := getSudokuSolved(t, &sudokutype.SudokuCommon{})

	generator.Randomize(field, rand.New(rand.NewSource(time.Now().UnixNano()))) //nolint:gosec // math/rand is ok

	checkRaw := handlers.CheckSudoku(
		check.CheckSudokuParams{ //nolint:exhaustruct // false positive
			ToCheck: &models.ToCheck{
				Field:    handlers.FieldInternalToAPI(field),
				Original: original,
			},
		},
	)

	checkRes, ok := checkRaw.(*check.CheckSudokuDefault)
	require.True(t, ok, "checking incompatible solution")
	require.Equal(t, int32(700), *checkRes.Payload.Code, "checking incompatible solution")
}

func getSudokuSolved[S sudokutype.Sudoku]( //nolint:ireturn // false positive
	t *testing.T,
	field S,
) (S, *models.Original) {
	t.Helper()

	getRaw := handlers.GetSudoku(
		get.GetSudokuParams{ //nolint:exhaustruct // false positive
			Size:       pointy.String("common"),
			Difficulty: pointy.String("common"),
		},
	)

	getRes, ok := getRaw.(*get.GetSudokuOK)
	require.True(t, ok, "getting sudoku")

	handlers.FieldAPIToInternal(getRes.Payload.Field, field)

	ok, err := solver.Solve(field)
	require.NoError(t, err, "solving field\n%s", sudokutype.Dump(field))
	require.True(t, ok, "solving field:\n%s", sudokutype.Dump(field))

	return field, getRes.Payload.Original
}

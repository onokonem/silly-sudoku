package handlers

import (
	"fmt"

	"github.com/openlyinc/pointy"

	"silly-sudoku/internal/handlers/internal/encdec"
	"silly-sudoku/internal/oapi/models"
	"silly-sudoku/internal/oapi/restapi/operations/check"
	"silly-sudoku/internal/sudokutype"
	"silly-sudoku/internal/validator"
)

//nolint:gomnd // false positives
func CheckSudoku(params check.CheckSudokuParams) check.CheckSudokuResponder { //nolint:ireturn //false positive
	switch len(params.ToCheck.Field) {
	case sudokutype.SizeCommon, sudokutype.SizeBig, sudokutype.SizeHuge:
	default:
		return check.NewCheckSudokuDefault(422).WithPayload(
			&models.Error{
				Code:    pointy.Int32(702),
				Message: pointy.String(fmt.Sprintf("unexpected field size: %d", len(params.ToCheck.Field))),
			},
		)
	}

	if err := checkInputFieldSize(params.ToCheck.Field); err != nil {
		return check.NewCheckSudokuDefault(422).WithPayload(err)
	}

	switch len(params.ToCheck.Field) {
	case sudokutype.SizeCommon:
		return checkSudoku(
			params,
			&sudokutype.SudokuCommon{},
			&sudokutype.SudokuCommon{},
			&sudokutype.SudokuCommon{},
		)
	case sudokutype.SizeBig:
		return checkSudoku(
			params,
			&sudokutype.SudokuBig{},
			&sudokutype.SudokuBig{},
			&sudokutype.SudokuBig{},
		)
	default:
		return checkSudoku(
			params,
			&sudokutype.SudokuHuge{},
			&sudokutype.SudokuHuge{},
			&sudokutype.SudokuHuge{},
		)
	}
}

//nolint:gomnd // false positives
func checkSudoku[S sudokutype.Sudoku]( //nolint:ireturn //false positive
	params check.CheckSudokuParams,
	solution S,
	field S,
	original S,
) check.CheckSudokuResponder {
	FieldAPIToInternal(params.ToCheck.Field, solution)

	ok, err := validator.Validate(solution)
	if err != nil {
		return check.NewCheckSudokuDefault(422).WithPayload(
			&models.Error{
				Code:    pointy.Int32(702),
				Message: pointy.String(fmt.Sprintf("validation failed: %+v", err)),
			},
		)
	}

	if !ok {
		return check.NewCheckSudokuDefault(422).WithPayload(
			&models.Error{
				Code:    pointy.Int32(702),
				Message: pointy.String("validation failed"),
			},
		)
	}

	if err = encdec.Dec([]byte(*params.ToCheck.Original), field, original); err != nil {
		return check.NewCheckSudokuDefault(400).WithPayload(
			&models.Error{
				Code:    pointy.Int32(701),
				Message: pointy.String(fmt.Sprintf("decryption failed: %+v", err)),
			},
		)
	}

	if !sudokutype.EqualPartial(solution, field) {
		return check.NewCheckSudokuDefault(400).WithPayload(
			&models.Error{
				Code:    pointy.Int32(700),
				Message: pointy.String("provided solution does not match to the original sudoku"),
			},
		)
	}

	return check.NewCheckSudokuOK().WithPayload(
		&models.Result{
			SameAsOriginal: sudokutype.Equal(original, solution),
		},
	)
}

//nolint:gomnd // false positives
func checkInputFieldSize(field models.Field) *models.Error {
	for y, row := range field {
		if len(row) != len(field) {
			return &models.Error{
				Code:    pointy.Int32(702),
				Message: pointy.String(fmt.Sprintf("unexpected row %d size: %d", y, len(row))),
			}
		}
	}

	return nil
}

// FieldAPIToInternal is converting the API level sudoku representation to to the internal one.
func FieldAPIToInternal[S sudokutype.Sudoku](in models.Field, out S) {
	for y, row := range in {
		for x, v := range row {
			out.Set(y, x, int8(v))
		}
	}
}

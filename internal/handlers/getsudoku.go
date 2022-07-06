package handlers

import (
	"fmt"

	"github.com/openlyinc/pointy"

	"silly-sudoku/internal/generator"
	"silly-sudoku/internal/handlers/internal/encdec"
	"silly-sudoku/internal/oapi/models"
	"silly-sudoku/internal/oapi/restapi/operations/get"
	"silly-sudoku/internal/sudokutype"
)

const (
	difficultyEasy      float64 = 35.0 / 81.0
	difficultyHard      float64 = 25.0 / 81.0
	difficultyNightmare float64 = 20.0 / 81.0
)

func GetSudoku(params get.GetSudokuParams) get.GetSudokuResponder { //nolint:ireturn //false positive
	var difficulty float64

	switch *params.Difficulty {
	case "common":
		difficulty = difficultyEasy
	case "hard":
		difficulty = difficultyHard
	default:
		difficulty = difficultyNightmare
	}

	switch *params.Size {
	case "common":
		return getSudoku(
			&sudokutype.SudokuCommon{},
			&sudokutype.SudokuCommon{},
			&sudokutype.SudokuCommon{},
			difficulty,
		)
	case "big":
		return getSudoku(
			&sudokutype.SudokuBig{},
			&sudokutype.SudokuBig{},
			&sudokutype.SudokuBig{},
			difficulty,
		)
	default:
		return getSudoku(
			&sudokutype.SudokuHuge{},
			&sudokutype.SudokuHuge{},
			&sudokutype.SudokuHuge{},
			difficulty,
		)
	}
}

//nolint:gomnd // false positives
func getSudoku[S sudokutype.Sudoku]( //nolint:ireturn //false positive
	field S,
	original S,
	tmp S,
	difficulty float64,
) get.GetSudokuResponder {
	generator.Generate(original)

	sudokutype.Copy(field, original)
	field = generator.Suppress(
		original,
		field,
		tmp,
		difficulty,
	)

	originalEnc, err := encdec.Enc(field, original)
	if err != nil {
		return get.NewGetSudokuDefault(500).WithPayload(
			&models.Error{
				Code:    pointy.Int32(500),
				Message: pointy.String(fmt.Sprintf("unexpected: %+v", err)),
			},
		)
	}

	return get.NewGetSudokuOK().WithPayload(
		&models.Sudoku{
			Field:    FieldInternalToAPI(field),
			Original: (*models.Original)(pointy.String(string(originalEnc))),
		},
	)
}

// FieldInternalToAPI is converting the internal  sudoku representation to the API level one.
func FieldInternalToAPI[S sudokutype.Sudoku](in S) models.Field {
	out := make(models.Field, len(in))

	for y := 0; y < len(in); y++ {
		out[y] = make(models.Row, len(in))
		for x := 0; x < len(in); x++ {
			out[y][x] = int32(in.Get(y, x))
		}
	}

	return out
}

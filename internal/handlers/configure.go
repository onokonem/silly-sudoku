package handlers

import (
	"silly-sudoku/internal/oapi/restapi/operations"
	"silly-sudoku/internal/oapi/restapi/operations/check"
	"silly-sudoku/internal/oapi/restapi/operations/get"
)

func Configure(api *operations.SudokuServerAPI) {
	api.GetGetSudokuHandler = get.GetSudokuHandlerFunc(GetSudoku)
	api.CheckCheckSudokuHandler = check.CheckSudokuHandlerFunc(CheckSudoku)
}

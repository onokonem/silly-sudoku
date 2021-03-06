// Code generated by go-swagger; DO NOT EDIT.

package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetSudokuParams creates a new GetSudokuParams object
// with the default values initialized.
func NewGetSudokuParams() GetSudokuParams {

	var (
		// initialize parameters with default values

		difficultyDefault = string("easy")
		sizeDefault       = string("common")
	)

	return GetSudokuParams{
		Difficulty: &difficultyDefault,

		Size: &sizeDefault,
	}
}

// GetSudokuParams contains all the bound params for the get sudoku operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSudoku
type GetSudokuParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	  Default: "easy"
	*/
	Difficulty *string
	/*
	  In: query
	  Default: "common"
	*/
	Size *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSudokuParams() beforehand.
func (o *GetSudokuParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDifficulty, qhkDifficulty, _ := qs.GetOK("difficulty")
	if err := o.bindDifficulty(qDifficulty, qhkDifficulty, route.Formats); err != nil {
		res = append(res, err)
	}

	qSize, qhkSize, _ := qs.GetOK("size")
	if err := o.bindSize(qSize, qhkSize, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDifficulty binds and validates parameter Difficulty from query.
func (o *GetSudokuParams) bindDifficulty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetSudokuParams()
		return nil
	}
	o.Difficulty = &raw

	if err := o.validateDifficulty(formats); err != nil {
		return err
	}

	return nil
}

// validateDifficulty carries on validations for parameter Difficulty
func (o *GetSudokuParams) validateDifficulty(formats strfmt.Registry) error {

	if err := validate.EnumCase("difficulty", "query", *o.Difficulty, []interface{}{"easy", "heavy", "nightmare"}, true); err != nil {
		return err
	}

	return nil
}

// bindSize binds and validates parameter Size from query.
func (o *GetSudokuParams) bindSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetSudokuParams()
		return nil
	}
	o.Size = &raw

	if err := o.validateSize(formats); err != nil {
		return err
	}

	return nil
}

// validateSize carries on validations for parameter Size
func (o *GetSudokuParams) validateSize(formats strfmt.Registry) error {

	if err := validate.EnumCase("size", "query", *o.Size, []interface{}{"common", "big", "huge"}, true); err != nil {
		return err
	}

	return nil
}

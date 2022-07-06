// Code generated by go-swagger; DO NOT EDIT.

package check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"silly-sudoku/internal/oapi/models"
)

// CheckSudokuOKCode is the HTTP code returned for type CheckSudokuOK
const CheckSudokuOKCode int = 200

/*CheckSudokuOK validation result

swagger:response checkSudokuOK
*/
type CheckSudokuOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewCheckSudokuOK creates CheckSudokuOK with default headers values
func NewCheckSudokuOK() *CheckSudokuOK {

	return &CheckSudokuOK{}
}

// WithPayload adds the payload to the check sudoku o k response
func (o *CheckSudokuOK) WithPayload(payload *models.Result) *CheckSudokuOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check sudoku o k response
func (o *CheckSudokuOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckSudokuOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CheckSudokuOK) CheckSudokuResponder() {}

/*CheckSudokuDefault generic error response

swagger:response checkSudokuDefault
*/
type CheckSudokuDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCheckSudokuDefault creates CheckSudokuDefault with default headers values
func NewCheckSudokuDefault(code int) *CheckSudokuDefault {
	if code <= 0 {
		code = 500
	}

	return &CheckSudokuDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the check sudoku default response
func (o *CheckSudokuDefault) WithStatusCode(code int) *CheckSudokuDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the check sudoku default response
func (o *CheckSudokuDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the check sudoku default response
func (o *CheckSudokuDefault) WithPayload(payload *models.Error) *CheckSudokuDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the check sudoku default response
func (o *CheckSudokuDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CheckSudokuDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *CheckSudokuDefault) CheckSudokuResponder() {}

type CheckSudokuNotImplementedResponder struct {
	middleware.Responder
}

func (*CheckSudokuNotImplementedResponder) CheckSudokuResponder() {}

func CheckSudokuNotImplemented() CheckSudokuResponder {
	return &CheckSudokuNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.CheckSudoku has not yet been implemented",
		),
	}
}

type CheckSudokuResponder interface {
	middleware.Responder
	CheckSudokuResponder()
}

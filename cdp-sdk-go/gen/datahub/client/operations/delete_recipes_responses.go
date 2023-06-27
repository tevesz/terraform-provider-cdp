// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// DeleteRecipesReader is a Reader for the DeleteRecipes structure.
type DeleteRecipesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecipesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecipesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteRecipesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteRecipesOK creates a DeleteRecipesOK with default headers values
func NewDeleteRecipesOK() *DeleteRecipesOK {
	return &DeleteRecipesOK{}
}

/*
DeleteRecipesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DeleteRecipesOK struct {
	Payload *models.DeleteRecipesResponse
}

// IsSuccess returns true when this delete recipes o k response has a 2xx status code
func (o *DeleteRecipesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete recipes o k response has a 3xx status code
func (o *DeleteRecipesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete recipes o k response has a 4xx status code
func (o *DeleteRecipesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete recipes o k response has a 5xx status code
func (o *DeleteRecipesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete recipes o k response a status code equal to that given
func (o *DeleteRecipesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete recipes o k response
func (o *DeleteRecipesOK) Code() int {
	return 200
}

func (o *DeleteRecipesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteRecipes][%d] deleteRecipesOK  %+v", 200, o.Payload)
}

func (o *DeleteRecipesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteRecipes][%d] deleteRecipesOK  %+v", 200, o.Payload)
}

func (o *DeleteRecipesOK) GetPayload() *models.DeleteRecipesResponse {
	return o.Payload
}

func (o *DeleteRecipesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteRecipesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecipesDefault creates a DeleteRecipesDefault with default headers values
func NewDeleteRecipesDefault(code int) *DeleteRecipesDefault {
	return &DeleteRecipesDefault{
		_statusCode: code,
	}
}

/*
DeleteRecipesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteRecipesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete recipes default response has a 2xx status code
func (o *DeleteRecipesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete recipes default response has a 3xx status code
func (o *DeleteRecipesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete recipes default response has a 4xx status code
func (o *DeleteRecipesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete recipes default response has a 5xx status code
func (o *DeleteRecipesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete recipes default response a status code equal to that given
func (o *DeleteRecipesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete recipes default response
func (o *DeleteRecipesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteRecipesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteRecipes][%d] deleteRecipes default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRecipesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteRecipes][%d] deleteRecipes default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteRecipesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRecipesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
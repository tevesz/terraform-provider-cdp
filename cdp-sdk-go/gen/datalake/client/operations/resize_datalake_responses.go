// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// ResizeDatalakeReader is a Reader for the ResizeDatalake structure.
type ResizeDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResizeDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResizeDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResizeDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResizeDatalakeOK creates a ResizeDatalakeOK with default headers values
func NewResizeDatalakeOK() *ResizeDatalakeOK {
	return &ResizeDatalakeOK{}
}

/*
ResizeDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid resize datalake request.
*/
type ResizeDatalakeOK struct {
	Payload *models.ResizeDatalakeResponse
}

// IsSuccess returns true when this resize datalake o k response has a 2xx status code
func (o *ResizeDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resize datalake o k response has a 3xx status code
func (o *ResizeDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resize datalake o k response has a 4xx status code
func (o *ResizeDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resize datalake o k response has a 5xx status code
func (o *ResizeDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resize datalake o k response a status code equal to that given
func (o *ResizeDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resize datalake o k response
func (o *ResizeDatalakeOK) Code() int {
	return 200
}

func (o *ResizeDatalakeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/resizeDatalake][%d] resizeDatalakeOK  %+v", 200, o.Payload)
}

func (o *ResizeDatalakeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/resizeDatalake][%d] resizeDatalakeOK  %+v", 200, o.Payload)
}

func (o *ResizeDatalakeOK) GetPayload() *models.ResizeDatalakeResponse {
	return o.Payload
}

func (o *ResizeDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResizeDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResizeDatalakeDefault creates a ResizeDatalakeDefault with default headers values
func NewResizeDatalakeDefault(code int) *ResizeDatalakeDefault {
	return &ResizeDatalakeDefault{
		_statusCode: code,
	}
}

/*
ResizeDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ResizeDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this resize datalake default response has a 2xx status code
func (o *ResizeDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resize datalake default response has a 3xx status code
func (o *ResizeDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resize datalake default response has a 4xx status code
func (o *ResizeDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resize datalake default response has a 5xx status code
func (o *ResizeDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resize datalake default response a status code equal to that given
func (o *ResizeDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resize datalake default response
func (o *ResizeDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *ResizeDatalakeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/resizeDatalake][%d] resizeDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *ResizeDatalakeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/resizeDatalake][%d] resizeDatalake default  %+v", o._statusCode, o.Payload)
}

func (o *ResizeDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResizeDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

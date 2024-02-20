// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// ListConnectedDataServicesReader is a Reader for the ListConnectedDataServices structure.
type ListConnectedDataServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConnectedDataServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConnectedDataServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListConnectedDataServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListConnectedDataServicesOK creates a ListConnectedDataServicesOK with default headers values
func NewListConnectedDataServicesOK() *ListConnectedDataServicesOK {
	return &ListConnectedDataServicesOK{}
}

/*
ListConnectedDataServicesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListConnectedDataServicesOK struct {
	Payload *models.ListConnectedDataServicesResponse
}

// IsSuccess returns true when this list connected data services o k response has a 2xx status code
func (o *ListConnectedDataServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list connected data services o k response has a 3xx status code
func (o *ListConnectedDataServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list connected data services o k response has a 4xx status code
func (o *ListConnectedDataServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list connected data services o k response has a 5xx status code
func (o *ListConnectedDataServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list connected data services o k response a status code equal to that given
func (o *ListConnectedDataServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list connected data services o k response
func (o *ListConnectedDataServicesOK) Code() int {
	return 200
}

func (o *ListConnectedDataServicesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/listConnectedDataServices][%d] listConnectedDataServicesOK  %+v", 200, o.Payload)
}

func (o *ListConnectedDataServicesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/listConnectedDataServices][%d] listConnectedDataServicesOK  %+v", 200, o.Payload)
}

func (o *ListConnectedDataServicesOK) GetPayload() *models.ListConnectedDataServicesResponse {
	return o.Payload
}

func (o *ListConnectedDataServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListConnectedDataServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConnectedDataServicesDefault creates a ListConnectedDataServicesDefault with default headers values
func NewListConnectedDataServicesDefault(code int) *ListConnectedDataServicesDefault {
	return &ListConnectedDataServicesDefault{
		_statusCode: code,
	}
}

/*
ListConnectedDataServicesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListConnectedDataServicesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list connected data services default response has a 2xx status code
func (o *ListConnectedDataServicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list connected data services default response has a 3xx status code
func (o *ListConnectedDataServicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list connected data services default response has a 4xx status code
func (o *ListConnectedDataServicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list connected data services default response has a 5xx status code
func (o *ListConnectedDataServicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list connected data services default response a status code equal to that given
func (o *ListConnectedDataServicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list connected data services default response
func (o *ListConnectedDataServicesDefault) Code() int {
	return o._statusCode
}

func (o *ListConnectedDataServicesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/listConnectedDataServices][%d] listConnectedDataServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListConnectedDataServicesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/listConnectedDataServices][%d] listConnectedDataServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListConnectedDataServicesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConnectedDataServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
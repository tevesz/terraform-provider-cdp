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

// SyncAllUsersReader is a Reader for the SyncAllUsers structure.
type SyncAllUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncAllUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncAllUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncAllUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncAllUsersOK creates a SyncAllUsersOK with default headers values
func NewSyncAllUsersOK() *SyncAllUsersOK {
	return &SyncAllUsersOK{}
}

/*
SyncAllUsersOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SyncAllUsersOK struct {
	Payload *models.SyncAllUsersResponse
}

// IsSuccess returns true when this sync all users o k response has a 2xx status code
func (o *SyncAllUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync all users o k response has a 3xx status code
func (o *SyncAllUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync all users o k response has a 4xx status code
func (o *SyncAllUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync all users o k response has a 5xx status code
func (o *SyncAllUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync all users o k response a status code equal to that given
func (o *SyncAllUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync all users o k response
func (o *SyncAllUsersOK) Code() int {
	return 200
}

func (o *SyncAllUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncAllUsers][%d] syncAllUsersOK  %+v", 200, o.Payload)
}

func (o *SyncAllUsersOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncAllUsers][%d] syncAllUsersOK  %+v", 200, o.Payload)
}

func (o *SyncAllUsersOK) GetPayload() *models.SyncAllUsersResponse {
	return o.Payload
}

func (o *SyncAllUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyncAllUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncAllUsersDefault creates a SyncAllUsersDefault with default headers values
func NewSyncAllUsersDefault(code int) *SyncAllUsersDefault {
	return &SyncAllUsersDefault{
		_statusCode: code,
	}
}

/*
SyncAllUsersDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SyncAllUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this sync all users default response has a 2xx status code
func (o *SyncAllUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync all users default response has a 3xx status code
func (o *SyncAllUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync all users default response has a 4xx status code
func (o *SyncAllUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync all users default response has a 5xx status code
func (o *SyncAllUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync all users default response a status code equal to that given
func (o *SyncAllUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync all users default response
func (o *SyncAllUsersDefault) Code() int {
	return o._statusCode
}

func (o *SyncAllUsersDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncAllUsers][%d] syncAllUsers default  %+v", o._statusCode, o.Payload)
}

func (o *SyncAllUsersDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/syncAllUsers][%d] syncAllUsers default  %+v", o._statusCode, o.Payload)
}

func (o *SyncAllUsersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SyncAllUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

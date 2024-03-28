// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// ListRestoreSnapshotsReader is a Reader for the ListRestoreSnapshots structure.
type ListRestoreSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRestoreSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRestoreSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRestoreSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRestoreSnapshotsOK creates a ListRestoreSnapshotsOK with default headers values
func NewListRestoreSnapshotsOK() *ListRestoreSnapshotsOK {
	return &ListRestoreSnapshotsOK{}
}

/*
ListRestoreSnapshotsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListRestoreSnapshotsOK struct {
	Payload *models.ListRestoreSnapshotsResponse
}

// IsSuccess returns true when this list restore snapshots o k response has a 2xx status code
func (o *ListRestoreSnapshotsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list restore snapshots o k response has a 3xx status code
func (o *ListRestoreSnapshotsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list restore snapshots o k response has a 4xx status code
func (o *ListRestoreSnapshotsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list restore snapshots o k response has a 5xx status code
func (o *ListRestoreSnapshotsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list restore snapshots o k response a status code equal to that given
func (o *ListRestoreSnapshotsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list restore snapshots o k response
func (o *ListRestoreSnapshotsOK) Code() int {
	return 200
}

func (o *ListRestoreSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listRestoreSnapshots][%d] listRestoreSnapshotsOK  %+v", 200, o.Payload)
}

func (o *ListRestoreSnapshotsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listRestoreSnapshots][%d] listRestoreSnapshotsOK  %+v", 200, o.Payload)
}

func (o *ListRestoreSnapshotsOK) GetPayload() *models.ListRestoreSnapshotsResponse {
	return o.Payload
}

func (o *ListRestoreSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRestoreSnapshotsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRestoreSnapshotsDefault creates a ListRestoreSnapshotsDefault with default headers values
func NewListRestoreSnapshotsDefault(code int) *ListRestoreSnapshotsDefault {
	return &ListRestoreSnapshotsDefault{
		_statusCode: code,
	}
}

/*
ListRestoreSnapshotsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListRestoreSnapshotsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list restore snapshots default response has a 2xx status code
func (o *ListRestoreSnapshotsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list restore snapshots default response has a 3xx status code
func (o *ListRestoreSnapshotsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list restore snapshots default response has a 4xx status code
func (o *ListRestoreSnapshotsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list restore snapshots default response has a 5xx status code
func (o *ListRestoreSnapshotsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list restore snapshots default response a status code equal to that given
func (o *ListRestoreSnapshotsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list restore snapshots default response
func (o *ListRestoreSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *ListRestoreSnapshotsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listRestoreSnapshots][%d] listRestoreSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoreSnapshotsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/opdb/listRestoreSnapshots][%d] listRestoreSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoreSnapshotsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListRestoreSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
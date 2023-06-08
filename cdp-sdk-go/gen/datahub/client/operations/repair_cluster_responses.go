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

// RepairClusterReader is a Reader for the RepairCluster structure.
type RepairClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepairClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRepairClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepairClusterOK creates a RepairClusterOK with default headers values
func NewRepairClusterOK() *RepairClusterOK {
	return &RepairClusterOK{}
}

/*
RepairClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RepairClusterOK struct {
	Payload models.RepairClusterResponse
}

// IsSuccess returns true when this repair cluster o k response has a 2xx status code
func (o *RepairClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repair cluster o k response has a 3xx status code
func (o *RepairClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repair cluster o k response has a 4xx status code
func (o *RepairClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repair cluster o k response has a 5xx status code
func (o *RepairClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repair cluster o k response a status code equal to that given
func (o *RepairClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repair cluster o k response
func (o *RepairClusterOK) Code() int {
	return 200
}

func (o *RepairClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/repairCluster][%d] repairClusterOK  %+v", 200, o.Payload)
}

func (o *RepairClusterOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/repairCluster][%d] repairClusterOK  %+v", 200, o.Payload)
}

func (o *RepairClusterOK) GetPayload() models.RepairClusterResponse {
	return o.Payload
}

func (o *RepairClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepairClusterDefault creates a RepairClusterDefault with default headers values
func NewRepairClusterDefault(code int) *RepairClusterDefault {
	return &RepairClusterDefault{
		_statusCode: code,
	}
}

/*
RepairClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RepairClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this repair cluster default response has a 2xx status code
func (o *RepairClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this repair cluster default response has a 3xx status code
func (o *RepairClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this repair cluster default response has a 4xx status code
func (o *RepairClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this repair cluster default response has a 5xx status code
func (o *RepairClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this repair cluster default response a status code equal to that given
func (o *RepairClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the repair cluster default response
func (o *RepairClusterDefault) Code() int {
	return o._statusCode
}

func (o *RepairClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/repairCluster][%d] repairCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RepairClusterDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/repairCluster][%d] repairCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RepairClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RepairClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

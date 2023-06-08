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

// DeleteCustomConfigurationsReader is a Reader for the DeleteCustomConfigurations structure.
type DeleteCustomConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteCustomConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteCustomConfigurationsOK creates a DeleteCustomConfigurationsOK with default headers values
func NewDeleteCustomConfigurationsOK() *DeleteCustomConfigurationsOK {
	return &DeleteCustomConfigurationsOK{}
}

/*
DeleteCustomConfigurationsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DeleteCustomConfigurationsOK struct {
	Payload *models.DeleteCustomConfigurationsResponse
}

// IsSuccess returns true when this delete custom configurations o k response has a 2xx status code
func (o *DeleteCustomConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete custom configurations o k response has a 3xx status code
func (o *DeleteCustomConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete custom configurations o k response has a 4xx status code
func (o *DeleteCustomConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete custom configurations o k response has a 5xx status code
func (o *DeleteCustomConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete custom configurations o k response a status code equal to that given
func (o *DeleteCustomConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete custom configurations o k response
func (o *DeleteCustomConfigurationsOK) Code() int {
	return 200
}

func (o *DeleteCustomConfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCustomConfigurations][%d] deleteCustomConfigurationsOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomConfigurationsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCustomConfigurations][%d] deleteCustomConfigurationsOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomConfigurationsOK) GetPayload() *models.DeleteCustomConfigurationsResponse {
	return o.Payload
}

func (o *DeleteCustomConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteCustomConfigurationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomConfigurationsDefault creates a DeleteCustomConfigurationsDefault with default headers values
func NewDeleteCustomConfigurationsDefault(code int) *DeleteCustomConfigurationsDefault {
	return &DeleteCustomConfigurationsDefault{
		_statusCode: code,
	}
}

/*
DeleteCustomConfigurationsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteCustomConfigurationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete custom configurations default response has a 2xx status code
func (o *DeleteCustomConfigurationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete custom configurations default response has a 3xx status code
func (o *DeleteCustomConfigurationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete custom configurations default response has a 4xx status code
func (o *DeleteCustomConfigurationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete custom configurations default response has a 5xx status code
func (o *DeleteCustomConfigurationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete custom configurations default response a status code equal to that given
func (o *DeleteCustomConfigurationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete custom configurations default response
func (o *DeleteCustomConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCustomConfigurationsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCustomConfigurations][%d] deleteCustomConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomConfigurationsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCustomConfigurations][%d] deleteCustomConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteCustomConfigurationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCustomConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

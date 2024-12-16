// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OperationStatus Status of an operation.
//
// swagger:model OperationStatus
type OperationStatus string

func NewOperationStatus(value OperationStatus) *OperationStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OperationStatus.
func (m OperationStatus) Pointer() *OperationStatus {
	return &m
}

const (

	// OperationStatusNEVERRUN captures enum value "NEVER_RUN"
	OperationStatusNEVERRUN OperationStatus = "NEVER_RUN"

	// OperationStatusREQUESTED captures enum value "REQUESTED"
	OperationStatusREQUESTED OperationStatus = "REQUESTED"

	// OperationStatusREJECTED captures enum value "REJECTED"
	OperationStatusREJECTED OperationStatus = "REJECTED"

	// OperationStatusRUNNING captures enum value "RUNNING"
	OperationStatusRUNNING OperationStatus = "RUNNING"

	// OperationStatusCOMPLETED captures enum value "COMPLETED"
	OperationStatusCOMPLETED OperationStatus = "COMPLETED"

	// OperationStatusFAILED captures enum value "FAILED"
	OperationStatusFAILED OperationStatus = "FAILED"

	// OperationStatusTIMEDOUT captures enum value "TIMEDOUT"
	OperationStatusTIMEDOUT OperationStatus = "TIMEDOUT"
)

// for schema
var operationStatusEnum []interface{}

func init() {
	var res []OperationStatus
	if err := json.Unmarshal([]byte(`["NEVER_RUN","REQUESTED","REJECTED","RUNNING","COMPLETED","FAILED","TIMEDOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operationStatusEnum = append(operationStatusEnum, v)
	}
}

func (m OperationStatus) validateOperationStatusEnum(path, location string, value OperationStatus) error {
	if err := validate.EnumCase(path, location, value, operationStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operation status
func (m OperationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperationStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this operation status based on context it is used
func (m OperationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
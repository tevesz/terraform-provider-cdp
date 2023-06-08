// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncComponentVersionsFromCmRequest Datalake sync CM component versions request.
//
// swagger:model SyncComponentVersionsFromCmRequest
type SyncComponentVersionsFromCmRequest struct {

	// The name or CRN of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`
}

// Validate validates this sync component versions from cm request
func (m *SyncComponentVersionsFromCmRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncComponentVersionsFromCmRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sync component versions from cm request based on context it is used
func (m *SyncComponentVersionsFromCmRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyncComponentVersionsFromCmRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncComponentVersionsFromCmRequest) UnmarshalBinary(b []byte) error {
	var res SyncComponentVersionsFromCmRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

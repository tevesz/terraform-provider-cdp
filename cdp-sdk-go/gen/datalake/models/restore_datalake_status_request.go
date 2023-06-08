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

// RestoreDatalakeStatusRequest Request object to get the status of a restore operation.
//
// swagger:model RestoreDatalakeStatusRequest
type RestoreDatalakeStatusRequest struct {

	// The name of the Data Lake for which the most recent restore status will be retrieved.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// Unique identifier of the restore operation performed.
	RestoreID string `json:"restoreId,omitempty"`
}

// Validate validates this restore datalake status request
func (m *RestoreDatalakeStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreDatalakeStatusRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore datalake status request based on context it is used
func (m *RestoreDatalakeStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreDatalakeStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreDatalakeStatusRequest) UnmarshalBinary(b []byte) error {
	var res RestoreDatalakeStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

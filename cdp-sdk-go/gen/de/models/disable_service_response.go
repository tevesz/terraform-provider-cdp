// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisableServiceResponse DisableService response object.
//
// swagger:model DisableServiceResponse
type DisableServiceResponse struct {

	// Status of deletion
	Status string `json:"status,omitempty"`
}

// Validate validates this disable service response
func (m *DisableServiceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disable service response based on context it is used
func (m *DisableServiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisableServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisableServiceResponse) UnmarshalBinary(b []byte) error {
	var res DisableServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
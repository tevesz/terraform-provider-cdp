// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HealthCheck The result of a health check.
//
// swagger:model HealthCheck
type HealthCheck struct {

	// The name of service health check.
	Name string `json:"name,omitempty"`

	// The service health check summary.
	Summary string `json:"summary,omitempty"`
}

// Validate validates this health check
func (m *HealthCheck) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this health check based on context it is used
func (m *HealthCheck) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HealthCheck) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthCheck) UnmarshalBinary(b []byte) error {
	var res HealthCheck
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

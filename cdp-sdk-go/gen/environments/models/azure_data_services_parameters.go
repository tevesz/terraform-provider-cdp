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

// AzureDataServicesParameters Azure-specific Data Service parameters response.
//
// swagger:model AzureDataServicesParameters
type AzureDataServicesParameters struct {

	// User-assigned managed identity used by the AKS control plane.
	// Required: true
	SharedManagedIdentity *string `json:"sharedManagedIdentity"`
}

// Validate validates this azure data services parameters
func (m *AzureDataServicesParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSharedManagedIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureDataServicesParameters) validateSharedManagedIdentity(formats strfmt.Registry) error {

	if err := validate.Required("sharedManagedIdentity", "body", m.SharedManagedIdentity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure data services parameters based on context it is used
func (m *AzureDataServicesParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureDataServicesParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureDataServicesParameters) UnmarshalBinary(b []byte) error {
	var res AzureDataServicesParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
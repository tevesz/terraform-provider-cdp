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

// ListCertificatesRequest The request of listing SHA-1 fingerprints of certificates in Global Trust Store
//
// swagger:model ListCertificatesRequest
type ListCertificatesRequest struct {

	// The name or CRN of the database.
	// Required: true
	Database *string `json:"database"`

	// The name or CRN of the environment.
	// Required: true
	Environment *string `json:"environment"`
}

// Validate validates this list certificates request
func (m *ListCertificatesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCertificatesRequest) validateDatabase(formats strfmt.Registry) error {

	if err := validate.Required("database", "body", m.Database); err != nil {
		return err
	}

	return nil
}

func (m *ListCertificatesRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list certificates request based on context it is used
func (m *ListCertificatesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListCertificatesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCertificatesRequest) UnmarshalBinary(b []byte) error {
	var res ListCertificatesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

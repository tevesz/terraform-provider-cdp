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

// UploadCertificateRequest The request of uploading a custom certificate to Global Trust Store
//
// swagger:model UploadCertificateRequest
type UploadCertificateRequest struct {

	// Custom certificate in PEM format
	// Required: true
	Certificate *string `json:"certificate"`

	// The name or CRN of the database.
	// Required: true
	Database *string `json:"database"`

	// The name or CRN of the environment.
	// Required: true
	Environment *string `json:"environment"`
}

// Validate validates this upload certificate request
func (m *UploadCertificateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

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

func (m *UploadCertificateRequest) validateCertificate(formats strfmt.Registry) error {

	if err := validate.Required("certificate", "body", m.Certificate); err != nil {
		return err
	}

	return nil
}

func (m *UploadCertificateRequest) validateDatabase(formats strfmt.Registry) error {

	if err := validate.Required("database", "body", m.Database); err != nil {
		return err
	}

	return nil
}

func (m *UploadCertificateRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upload certificate request based on context it is used
func (m *UploadCertificateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadCertificateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadCertificateRequest) UnmarshalBinary(b []byte) error {
	var res UploadCertificateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

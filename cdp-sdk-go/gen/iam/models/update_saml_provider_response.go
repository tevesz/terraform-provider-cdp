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

// UpdateSamlProviderResponse Response object for an updating SAML provider request.
//
// swagger:model UpdateSamlProviderResponse
type UpdateSamlProviderResponse struct {

	// The SAML provider.
	// Required: true
	SamlProvider *SamlProvider `json:"samlProvider"`
}

// Validate validates this update saml provider response
func (m *UpdateSamlProviderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamlProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSamlProviderResponse) validateSamlProvider(formats strfmt.Registry) error {

	if err := validate.Required("samlProvider", "body", m.SamlProvider); err != nil {
		return err
	}

	if m.SamlProvider != nil {
		if err := m.SamlProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("samlProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("samlProvider")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update saml provider response based on the context it is used
func (m *UpdateSamlProviderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamlProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSamlProviderResponse) contextValidateSamlProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.SamlProvider != nil {
		if err := m.SamlProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("samlProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("samlProvider")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSamlProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSamlProviderResponse) UnmarshalBinary(b []byte) error {
	var res UpdateSamlProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
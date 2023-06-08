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

// DescribeSamlProviderResponse Response object for a describe SAML request.
//
// swagger:model DescribeSamlProviderResponse
type DescribeSamlProviderResponse struct {

	// The SAML provider.
	// Required: true
	SamlProvider *SamlProvider `json:"samlProvider"`
}

// Validate validates this describe saml provider response
func (m *DescribeSamlProviderResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamlProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeSamlProviderResponse) validateSamlProvider(formats strfmt.Registry) error {

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

// ContextValidate validate this describe saml provider response based on the context it is used
func (m *DescribeSamlProviderResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamlProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeSamlProviderResponse) contextValidateSamlProvider(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DescribeSamlProviderResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeSamlProviderResponse) UnmarshalBinary(b []byte) error {
	var res DescribeSamlProviderResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

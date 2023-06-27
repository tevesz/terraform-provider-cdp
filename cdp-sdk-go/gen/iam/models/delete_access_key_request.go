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

// DeleteAccessKeyRequest Request object for a delete access key request.
//
// swagger:model DeleteAccessKeyRequest
type DeleteAccessKeyRequest struct {

	// The ID of the access key.
	// Required: true
	AccessKeyID *string `json:"accessKeyId"`
}

// Validate validates this delete access key request
func (m *DeleteAccessKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteAccessKeyRequest) validateAccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("accessKeyId", "body", m.AccessKeyID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete access key request based on context it is used
func (m *DeleteAccessKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteAccessKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteAccessKeyRequest) UnmarshalBinary(b []byte) error {
	var res DeleteAccessKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
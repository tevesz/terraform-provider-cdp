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

// ListWorkspaceAccessResponse Response object for the ListWorkspaceAccess method.
//
// swagger:model ListWorkspaceAccessResponse
type ListWorkspaceAccessResponse struct {

	// The list of users that have access.
	// Required: true
	Users []string `json:"users"`
}

// Validate validates this list workspace access response
func (m *ListWorkspaceAccessResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListWorkspaceAccessResponse) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list workspace access response based on context it is used
func (m *ListWorkspaceAccessResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListWorkspaceAccessResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListWorkspaceAccessResponse) UnmarshalBinary(b []byte) error {
	var res ListWorkspaceAccessResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
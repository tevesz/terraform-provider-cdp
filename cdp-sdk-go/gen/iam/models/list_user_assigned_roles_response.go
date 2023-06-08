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

// ListUserAssignedRolesResponse Response object for a list user assigned roles request.
//
// swagger:model ListUserAssignedRolesResponse
type ListUserAssignedRolesResponse struct {

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`

	// The role CRNs assigned to the user.
	// Required: true
	RoleCrns []string `json:"roleCrns"`
}

// Validate validates this list user assigned roles response
func (m *ListUserAssignedRolesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleCrns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListUserAssignedRolesResponse) validateRoleCrns(formats strfmt.Registry) error {

	if err := validate.Required("roleCrns", "body", m.RoleCrns); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list user assigned roles response based on context it is used
func (m *ListUserAssignedRolesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListUserAssignedRolesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListUserAssignedRolesResponse) UnmarshalBinary(b []byte) error {
	var res ListUserAssignedRolesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// ResourceAssignee Information about the resource role assignee for the resource.
//
// swagger:model ResourceAssignee
type ResourceAssignee struct {

	// The CRN of the assignee.
	// Required: true
	AssigneeCrn *string `json:"assigneeCrn"`

	// The assigned resource role CRN.
	// Required: true
	ResourceRoleCrn *string `json:"resourceRoleCrn"`
}

// Validate validates this resource assignee
func (m *ResourceAssignee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssigneeCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceRoleCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceAssignee) validateAssigneeCrn(formats strfmt.Registry) error {

	if err := validate.Required("assigneeCrn", "body", m.AssigneeCrn); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAssignee) validateResourceRoleCrn(formats strfmt.Registry) error {

	if err := validate.Required("resourceRoleCrn", "body", m.ResourceRoleCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resource assignee based on context it is used
func (m *ResourceAssignee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAssignee) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAssignee) UnmarshalBinary(b []byte) error {
	var res ResourceAssignee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

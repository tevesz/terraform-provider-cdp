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

// ModifyClusterInstanceGroupRequest Request object for ModifyClusterInstanceGroup.
//
// swagger:model ModifyClusterInstanceGroupRequest
type ModifyClusterInstanceGroupRequest struct {

	// The name of the instance group of the workspace cluster to modify.
	// Required: true
	InstanceGroupName *string `json:"instanceGroupName"`

	// The desired instance type of the workspace cluster instance group.
	// Required: true
	InstanceType *string `json:"instanceType"`

	// The desired autoscaling min of the workspace cluster instance group.
	// Required: true
	Max *int32 `json:"max"`

	// The desired autoscaling min of the workspace cluster instance group.
	// Required: true
	Min *int32 `json:"min"`

	// The CRN of the workspace cluster to modify.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`
}

// Validate validates this modify cluster instance group request
func (m *ModifyClusterInstanceGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModifyClusterInstanceGroupRequest) validateInstanceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupName", "body", m.InstanceGroupName); err != nil {
		return err
	}

	return nil
}

func (m *ModifyClusterInstanceGroupRequest) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *ModifyClusterInstanceGroupRequest) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *ModifyClusterInstanceGroupRequest) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

func (m *ModifyClusterInstanceGroupRequest) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this modify cluster instance group request based on context it is used
func (m *ModifyClusterInstanceGroupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModifyClusterInstanceGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModifyClusterInstanceGroupRequest) UnmarshalBinary(b []byte) error {
	var res ModifyClusterInstanceGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
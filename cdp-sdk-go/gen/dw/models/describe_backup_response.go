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

// DescribeBackupResponse Response object for the describe backup request.
//
// swagger:model DescribeBackupResponse
type DescribeBackupResponse struct {

	// The backup details.
	// Required: true
	Backup *Backup `json:"backup"`
}

// Validate validates this describe backup response
func (m *DescribeBackupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeBackupResponse) validateBackup(formats strfmt.Registry) error {

	if err := validate.Required("backup", "body", m.Backup); err != nil {
		return err
	}

	if m.Backup != nil {
		if err := m.Backup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe backup response based on the context it is used
func (m *DescribeBackupResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeBackupResponse) contextValidateBackup(ctx context.Context, formats strfmt.Registry) error {

	if m.Backup != nil {
		if err := m.Backup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeBackupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeBackupResponse) UnmarshalBinary(b []byte) error {
	var res DescribeBackupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

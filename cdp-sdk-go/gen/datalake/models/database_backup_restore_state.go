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

// DatabaseBackupRestoreState The state of the database backup/restore operation.
//
// swagger:model DatabaseBackupRestoreState
type DatabaseBackupRestoreState struct {

	// The status of the database backup/restore.
	// Required: true
	Database *BackupRestoreOperationStatus `json:"database"`
}

// Validate validates this database backup restore state
func (m *DatabaseBackupRestoreState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackupRestoreState) validateDatabase(formats strfmt.Registry) error {

	if err := validate.Required("database", "body", m.Database); err != nil {
		return err
	}

	if m.Database != nil {
		if err := m.Database.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this database backup restore state based on the context it is used
func (m *DatabaseBackupRestoreState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseBackupRestoreState) contextValidateDatabase(ctx context.Context, formats strfmt.Registry) error {

	if m.Database != nil {
		if err := m.Database.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("database")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("database")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseBackupRestoreState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseBackupRestoreState) UnmarshalBinary(b []byte) error {
	var res DatabaseBackupRestoreState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
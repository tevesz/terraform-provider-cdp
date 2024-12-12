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

// RestoreSnapshotRequest Restore Snapshot Request.
//
// swagger:model RestoreSnapshotRequest
type RestoreSnapshotRequest struct {

	// The name of the original database.
	// Required: true
	DatabaseName *string `json:"databaseName"`

	// The name of the original environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// The name of the snapshot.
	// Required: true
	SnapshotName *string `json:"snapshotName"`

	// The name of the target database where the snapshot should be restored.
	// Required: true
	TargetDatabaseName *string `json:"targetDatabaseName"`

	// The name of the target environment where the snapshot should be restored.
	// Required: true
	TargetEnvironmentName *string `json:"targetEnvironmentName"`
}

// Validate validates this restore snapshot request
func (m *RestoreSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetDatabaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreSnapshotRequest) validateDatabaseName(formats strfmt.Registry) error {

	if err := validate.Required("databaseName", "body", m.DatabaseName); err != nil {
		return err
	}

	return nil
}

func (m *RestoreSnapshotRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *RestoreSnapshotRequest) validateSnapshotName(formats strfmt.Registry) error {

	if err := validate.Required("snapshotName", "body", m.SnapshotName); err != nil {
		return err
	}

	return nil
}

func (m *RestoreSnapshotRequest) validateTargetDatabaseName(formats strfmt.Registry) error {

	if err := validate.Required("targetDatabaseName", "body", m.TargetDatabaseName); err != nil {
		return err
	}

	return nil
}

func (m *RestoreSnapshotRequest) validateTargetEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("targetEnvironmentName", "body", m.TargetEnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restore snapshot request based on context it is used
func (m *RestoreSnapshotRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res RestoreSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

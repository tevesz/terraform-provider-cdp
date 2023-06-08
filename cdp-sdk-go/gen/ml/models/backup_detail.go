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

// BackupDetail Backup Detail response object for listing backups.
//
// swagger:model BackupDetail
type BackupDetail struct {

	// The CRN of the backup snapshot.
	BackupCrn string `json:"backupCrn,omitempty"`

	// The name of the backup snapshot.
	BackupName string `json:"backupName,omitempty"`

	// The status of the backup.
	BackupStatus string `json:"backupStatus,omitempty"`

	// The creation time of the backup snapshot.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// The CRN of the creator.
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// The version of the backed-up workspace at the time of backup.
	WorkspaceVersionAtBackup string `json:"workspaceVersionAtBackup,omitempty"`
}

// Validate validates this backup detail
func (m *BackupDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupDetail) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup detail based on context it is used
func (m *BackupDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupDetail) UnmarshalBinary(b []byte) error {
	var res BackupDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

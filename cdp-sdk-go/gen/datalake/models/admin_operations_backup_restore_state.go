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

// AdminOperationsBackupRestoreState The state of Cloudera Manager admin operations.
//
// swagger:model AdminOperationsBackupRestoreState
type AdminOperationsBackupRestoreState struct {

	// Validate storage permissions before running a backup/restore.
	// Required: true
	PrecheckStoragePermission *BackupRestoreOperationStatus `json:"precheckStoragePermission"`

	// Run the ranger audit collection validation in the backup precheck.
	// Required: true
	RangerAuditCollectionValidation *BackupRestoreOperationStatus `json:"rangerAuditCollectionValidation"`

	// The status of the start services operation this is triggered after the backup/restore is complete.
	// Required: true
	StartServices *BackupRestoreOperationStatus `json:"startServices"`

	// The status of the stop services operation that is triggered before the backup/restore is started.
	// Required: true
	StopServices *BackupRestoreOperationStatus `json:"stopServices"`
}

// Validate validates this admin operations backup restore state
func (m *AdminOperationsBackupRestoreState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrecheckStoragePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRangerAuditCollectionValidation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminOperationsBackupRestoreState) validatePrecheckStoragePermission(formats strfmt.Registry) error {

	if err := validate.Required("precheckStoragePermission", "body", m.PrecheckStoragePermission); err != nil {
		return err
	}

	if m.PrecheckStoragePermission != nil {
		if err := m.PrecheckStoragePermission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("precheckStoragePermission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("precheckStoragePermission")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) validateRangerAuditCollectionValidation(formats strfmt.Registry) error {

	if err := validate.Required("rangerAuditCollectionValidation", "body", m.RangerAuditCollectionValidation); err != nil {
		return err
	}

	if m.RangerAuditCollectionValidation != nil {
		if err := m.RangerAuditCollectionValidation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rangerAuditCollectionValidation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rangerAuditCollectionValidation")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) validateStartServices(formats strfmt.Registry) error {

	if err := validate.Required("startServices", "body", m.StartServices); err != nil {
		return err
	}

	if m.StartServices != nil {
		if err := m.StartServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startServices")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) validateStopServices(formats strfmt.Registry) error {

	if err := validate.Required("stopServices", "body", m.StopServices); err != nil {
		return err
	}

	if m.StopServices != nil {
		if err := m.StopServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stopServices")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this admin operations backup restore state based on the context it is used
func (m *AdminOperationsBackupRestoreState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrecheckStoragePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRangerAuditCollectionValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStopServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminOperationsBackupRestoreState) contextValidatePrecheckStoragePermission(ctx context.Context, formats strfmt.Registry) error {

	if m.PrecheckStoragePermission != nil {
		if err := m.PrecheckStoragePermission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("precheckStoragePermission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("precheckStoragePermission")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) contextValidateRangerAuditCollectionValidation(ctx context.Context, formats strfmt.Registry) error {

	if m.RangerAuditCollectionValidation != nil {
		if err := m.RangerAuditCollectionValidation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rangerAuditCollectionValidation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rangerAuditCollectionValidation")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) contextValidateStartServices(ctx context.Context, formats strfmt.Registry) error {

	if m.StartServices != nil {
		if err := m.StartServices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("startServices")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOperationsBackupRestoreState) contextValidateStopServices(ctx context.Context, formats strfmt.Registry) error {

	if m.StopServices != nil {
		if err := m.StopServices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stopServices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stopServices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminOperationsBackupRestoreState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminOperationsBackupRestoreState) UnmarshalBinary(b []byte) error {
	var res AdminOperationsBackupRestoreState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
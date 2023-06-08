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

// AuditEvent Audit event descibes an performed or performing in a given workspace.
//
// swagger:model AuditEvent
type AuditEvent struct {

	// Action the user has generated.
	// Required: true
	Action *string `json:"action"`

	// Time at creation of event.
	// Required: true
	// Format: date-time
	CreatedDate *strfmt.DateTime `json:"createdDate"`

	// Unique request ID to keep track of event.
	// Required: true
	RequestID *string `json:"requestID"`

	// UserCrn to track which user has caused the event.
	// Required: true
	UserCrn *string `json:"userCrn"`

	// Workspace crn where the event observed.
	// Required: true
	WorkspaceCrn *string `json:"workspaceCrn"`
}

// Validate validates this audit event
func (m *AuditEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserCrn(formats); err != nil {
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

func (m *AuditEvent) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *AuditEvent) validateCreatedDate(formats strfmt.Registry) error {

	if err := validate.Required("createdDate", "body", m.CreatedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEvent) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("requestID", "body", m.RequestID); err != nil {
		return err
	}

	return nil
}

func (m *AuditEvent) validateUserCrn(formats strfmt.Registry) error {

	if err := validate.Required("userCrn", "body", m.UserCrn); err != nil {
		return err
	}

	return nil
}

func (m *AuditEvent) validateWorkspaceCrn(formats strfmt.Registry) error {

	if err := validate.Required("workspaceCrn", "body", m.WorkspaceCrn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this audit event based on context it is used
func (m *AuditEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuditEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEvent) UnmarshalBinary(b []byte) error {
	var res AuditEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

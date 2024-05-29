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

// ServiceSummary Summary of a CDE service.
//
// swagger:model ServiceSummary
type ServiceSummary struct {

	// Cluster Id of the CDE Service.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Email Address of the CDE creator.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Timestamp of service enabling.
	EnablingTime string `json:"enablingTime,omitempty"`

	// CDP Environment Name.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Name of the CDE Service.
	// Required: true
	Name *string `json:"name"`

	// Status of the CDE service.
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this service summary
func (m *ServiceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceSummary) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceSummary) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceSummary) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this service summary based on context it is used
func (m *ServiceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceSummary) UnmarshalBinary(b []byte) error {
	var res ServiceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
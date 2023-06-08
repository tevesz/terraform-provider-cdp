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

// SetTelemetryFeaturesRequest Request object to enable environment level telemetry features.
//
// swagger:model SetTelemetryFeaturesRequest
type SetTelemetryFeaturesRequest struct {

	// The name or CRN of the environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Flag to enable environment level deployment log collection.
	ReportDeploymentLogs bool `json:"reportDeploymentLogs,omitempty"`

	// Flag to enable environment level workload analytics.
	WorkloadAnalytics bool `json:"workloadAnalytics,omitempty"`
}

// Validate validates this set telemetry features request
func (m *SetTelemetryFeaturesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetTelemetryFeaturesRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set telemetry features request based on context it is used
func (m *SetTelemetryFeaturesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetTelemetryFeaturesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetTelemetryFeaturesRequest) UnmarshalBinary(b []byte) error {
	var res SetTelemetryFeaturesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartDatalakeVerticalScalingResponse The response object for Data Lake vertical scaling.
//
// swagger:model StartDatalakeVerticalScalingResponse
type StartDatalakeVerticalScalingResponse struct {

	// The result of the operation.
	Result string `json:"result,omitempty"`
}

// Validate validates this start datalake vertical scaling response
func (m *StartDatalakeVerticalScalingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start datalake vertical scaling response based on context it is used
func (m *StartDatalakeVerticalScalingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartDatalakeVerticalScalingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartDatalakeVerticalScalingResponse) UnmarshalBinary(b []byte) error {
	var res StartDatalakeVerticalScalingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
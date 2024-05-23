// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeMlServingAppRequest Request object for the DescribeMlServingApp method.
//
// swagger:model DescribeMlServingAppRequest
type DescribeMlServingAppRequest struct {

	// The CRN Of the App.
	AppCrn string `json:"appCrn,omitempty"`

	// The name of the App. If this is supplied and appCrn is not supplied, either environmentCrn or environmentName is required.
	AppName string `json:"appName,omitempty"`

	// The CRN of the environment associated with this App. If appName is supplied and appCrn is not supplied, either environmentCrn or environmentName is required.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The name of the environment associated with this App. If appName is supplied and appCrn is not supplied, either environmentCrn or environmentName is required.
	EnvironmentName string `json:"environmentName,omitempty"`
}

// Validate validates this describe ml serving app request
func (m *DescribeMlServingAppRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe ml serving app request based on context it is used
func (m *DescribeMlServingAppRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeMlServingAppRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeMlServingAppRequest) UnmarshalBinary(b []byte) error {
	var res DescribeMlServingAppRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
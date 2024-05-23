// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GrantMlServingAppAccessRequest Request object for the GrantMlServingAppAccess method.
//
// swagger:model GrantMlServingAppAccessRequest
type GrantMlServingAppAccessRequest struct {

	// The cloud provider user id, such as ARN, which will be granted access to the Model Serving App's Kubernetes cluster.
	Identifier string `json:"identifier,omitempty"`

	// The CRN of the Model Serving App to grant access to.
	ResourceCrn string `json:"resourceCrn,omitempty"`
}

// Validate validates this grant ml serving app access request
func (m *GrantMlServingAppAccessRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this grant ml serving app access request based on context it is used
func (m *GrantMlServingAppAccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GrantMlServingAppAccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GrantMlServingAppAccessRequest) UnmarshalBinary(b []byte) error {
	var res GrantMlServingAppAccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
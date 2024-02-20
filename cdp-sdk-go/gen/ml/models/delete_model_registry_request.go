// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteModelRegistryRequest Request for deleting model registry.
//
// swagger:model DeleteModelRegistryRequest
type DeleteModelRegistryRequest struct {

	// Force delete a model registry workspace even if errors occur during deletion. Force delete removes the guarantee that resources in your cloud account will be cleaned up. By default, force is set to false.
	Force bool `json:"force,omitempty"`

	// Deprecated, please use modelRegistryCrn. The environment CRN of the model registry.
	ID string `json:"id,omitempty"`

	// CRN of the model registry to be deleted.
	ModelRegistryCrn string `json:"modelRegistryCrn,omitempty"`

	// Deprecated. The workspace CRN of the model registry (Public cloud only).
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`
}

// Validate validates this delete model registry request
func (m *DeleteModelRegistryRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete model registry request based on context it is used
func (m *DeleteModelRegistryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteModelRegistryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteModelRegistryRequest) UnmarshalBinary(b []byte) error {
	var res DeleteModelRegistryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
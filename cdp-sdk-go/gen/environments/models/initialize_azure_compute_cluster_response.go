// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InitializeAzureComputeClusterResponse Response object for an initialize Azure default compute cluster request.
//
// swagger:model InitializeAzureComputeClusterResponse
type InitializeAzureComputeClusterResponse struct {

	// The id of the related operation.
	OperationID string `json:"operationId,omitempty"`
}

// Validate validates this initialize azure compute cluster response
func (m *InitializeAzureComputeClusterResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this initialize azure compute cluster response based on context it is used
func (m *InitializeAzureComputeClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InitializeAzureComputeClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitializeAzureComputeClusterResponse) UnmarshalBinary(b []byte) error {
	var res InitializeAzureComputeClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
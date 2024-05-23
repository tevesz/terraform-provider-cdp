// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListMlServingAppsRequest Request object for the ListMlServingApps method.
//
// swagger:model ListMlServingAppsRequest
type ListMlServingAppsRequest struct {

	// Additional query options.
	QueryOptions []string `json:"queryOptions"`
}

// Validate validates this list ml serving apps request
func (m *ListMlServingAppsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list ml serving apps request based on context it is used
func (m *ListMlServingAppsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListMlServingAppsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMlServingAppsRequest) UnmarshalBinary(b []byte) error {
	var res ListMlServingAppsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
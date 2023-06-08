// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateServerSettingItem A DWX server setting update.
//
// swagger:model UpdateServerSettingItem
type UpdateServerSettingItem struct {

	// The identifier of the setting.
	ConfigurationKey string `json:"configurationKey,omitempty"`

	// The state of the setting.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update server setting item
func (m *UpdateServerSettingItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update server setting item based on context it is used
func (m *UpdateServerSettingItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateServerSettingItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateServerSettingItem) UnmarshalBinary(b []byte) error {
	var res UpdateServerSettingItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

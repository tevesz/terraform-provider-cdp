// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AttachedVolume The attached volume configuration.
//
// swagger:model AttachedVolume
type AttachedVolume struct {

	// The number of volumes.
	Count int32 `json:"count,omitempty"`

	// The size of each volume in GB.
	Size int64 `json:"size,omitempty"`

	// The type of the volumes.
	VolumeType string `json:"volumeType,omitempty"`
}

// Validate validates this attached volume
func (m *AttachedVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this attached volume based on context it is used
func (m *AttachedVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AttachedVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AttachedVolume) UnmarshalBinary(b []byte) error {
	var res AttachedVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

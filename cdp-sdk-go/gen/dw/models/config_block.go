// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigBlock Set values for a configuration file of a service.
//
// swagger:model ConfigBlock
type ConfigBlock struct {

	// Configuration file to update.
	ConfigFile string `json:"configFile,omitempty"`

	// Set values for text format configuration file e.g. TEXT, JSON, YAML, XML typed file.
	Content string `json:"content,omitempty"`

	// Set values for key-value format configuration file e.g. ENV, PROPERTIES, FLAGFILE, HADOOP_XML typed file.
	KeyValue map[string]string `json:"keyValue,omitempty"`
}

// Validate validates this config block
func (m *ConfigBlock) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config block based on context it is used
func (m *ConfigBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBlock) UnmarshalBinary(b []byte) error {
	var res ConfigBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

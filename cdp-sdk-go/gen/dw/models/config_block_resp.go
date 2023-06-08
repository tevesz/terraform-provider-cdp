// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigBlockResp A piece of configuration stored in the same place (e.g. same file or environment variables).
//
// swagger:model ConfigBlockResp
type ConfigBlockResp struct {

	// Contents of a ConfigBlock.
	Content *ConfigContentResp `json:"content,omitempty"`

	// Format of ConfigBlock.
	// Enum: [HADOOP_XML PROPERTIES TEXT JSON BINARY ENV FLAGFILE]
	Format string `json:"format,omitempty"`

	// ID of the ConfigBlock. Unique within an ApplicationConfig.
	ID string `json:"id,omitempty"`
}

// Validate validates this config block resp
func (m *ConfigBlockResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBlockResp) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

var configBlockRespTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HADOOP_XML","PROPERTIES","TEXT","JSON","BINARY","ENV","FLAGFILE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configBlockRespTypeFormatPropEnum = append(configBlockRespTypeFormatPropEnum, v)
	}
}

const (

	// ConfigBlockRespFormatHADOOPXML captures enum value "HADOOP_XML"
	ConfigBlockRespFormatHADOOPXML string = "HADOOP_XML"

	// ConfigBlockRespFormatPROPERTIES captures enum value "PROPERTIES"
	ConfigBlockRespFormatPROPERTIES string = "PROPERTIES"

	// ConfigBlockRespFormatTEXT captures enum value "TEXT"
	ConfigBlockRespFormatTEXT string = "TEXT"

	// ConfigBlockRespFormatJSON captures enum value "JSON"
	ConfigBlockRespFormatJSON string = "JSON"

	// ConfigBlockRespFormatBINARY captures enum value "BINARY"
	ConfigBlockRespFormatBINARY string = "BINARY"

	// ConfigBlockRespFormatENV captures enum value "ENV"
	ConfigBlockRespFormatENV string = "ENV"

	// ConfigBlockRespFormatFLAGFILE captures enum value "FLAGFILE"
	ConfigBlockRespFormatFLAGFILE string = "FLAGFILE"
)

// prop value enum
func (m *ConfigBlockResp) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configBlockRespTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfigBlockResp) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config block resp based on the context it is used
func (m *ConfigBlockResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBlockResp) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBlockResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBlockResp) UnmarshalBinary(b []byte) error {
	var res ConfigBlockResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

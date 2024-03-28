// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetK8sCertJKSResponse The response object for the getK8sCertJKS method.
//
// swagger:model GetK8sCertJKSResponse
type GetK8sCertJKSResponse struct {

	// Base64 encoded Java truststore.
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// The password for the received Java truststore.
	Password string `json:"password,omitempty"`
}

// Validate validates this get k8s cert j k s response
func (m *GetK8sCertJKSResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get k8s cert j k s response based on context it is used
func (m *GetK8sCertJKSResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetK8sCertJKSResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetK8sCertJKSResponse) UnmarshalBinary(b []byte) error {
	var res GetK8sCertJKSResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
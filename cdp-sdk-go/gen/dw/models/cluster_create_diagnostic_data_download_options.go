// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterCreateDiagnosticDataDownloadOptions Flags that denote which diagnostics to include for the cluster.
//
// swagger:model ClusterCreateDiagnosticDataDownloadOptions
type ClusterCreateDiagnosticDataDownloadOptions struct {

	// Include cluster info.
	IncludeClusterInfo *bool `json:"includeClusterInfo,omitempty"`

	// Include Istio system.
	IncludeIstioSystem *bool `json:"includeIstioSystem,omitempty"`

	// Include Kube system.
	IncludeKubeSystem *bool `json:"includeKubeSystem,omitempty"`

	// Include shared services.
	IncludeSharedServices *bool `json:"includeSharedServices,omitempty"`
}

// Validate validates this cluster create diagnostic data download options
func (m *ClusterCreateDiagnosticDataDownloadOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster create diagnostic data download options based on context it is used
func (m *ClusterCreateDiagnosticDataDownloadOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateDiagnosticDataDownloadOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateDiagnosticDataDownloadOptions) UnmarshalBinary(b []byte) error {
	var res ClusterCreateDiagnosticDataDownloadOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

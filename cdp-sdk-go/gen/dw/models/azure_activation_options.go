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

// AzureActivationOptions Options for activating an Azure environment.
//
// swagger:model AzureActivationOptions
type AzureActivationOptions struct {

	// Azure compute instance types that the environment is restricted to use. This affects the creation of virtual warehouses where this restriction will apply. Select an instance type that meets your computing, memory, networking, or storage needs. As of now, only a single instance type can be listed. Use describe-allowed-instance-types to see currently possible values and the default value used for the case it is not provided.
	ComputeInstanceTypes []string `json:"computeInstanceTypes"`

	// Enables Azure Availability Zones for the cluster deployment.
	EnableAZ bool `json:"enableAZ,omitempty"`

	// Enable Azure Private AKS mode.
	EnablePrivateAks *bool `json:"enablePrivateAks,omitempty"`

	// Enables private SQL for the cluster deployment.
	EnablePrivateSQL *bool `json:"enablePrivateSQL,omitempty"`

	// Whether to enable spot instances for Virtual warehouses. It cannot be updated later. If the AzureActivationOptions is not provided it defaults to false.
	EnableSpotInstances *bool `json:"enableSpotInstances,omitempty"`

	// Enable uptime SLA for Kubernetes API server. This option is deprecated and will be removed upon the next release of the DWX Public Cloud. The uptime SLA for the Kubernetes API server will be always enabled.
	EnableUptimeSLA *bool `json:"enableUptimeSLA,omitempty"`

	// Enable monitoring of Azure Kubernetes Service (AKS) cluster. Workspace ID for Azure log analytics.
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId,omitempty"`

	// Network outbound type. This setting controls the egress traffic for cluster nodes in Azure Kubernetes Service. Please refer to the following AKS documentation on the Azure portal. https://learn.microsoft.com/en-us/azure/aks/egress-outboundtype, https://learn.microsoft.com/en-us/azure/aks/nat-gateway
	// Enum: [LoadBalancer UserAssignedNATGateway UserDefinedRouting]
	OutboundType string `json:"outboundType,omitempty"`

	// Private DNS zone AKS resource ID.
	PrivateDNSZoneAKS string `json:"privateDNSZoneAKS,omitempty"`

	// ID of Azure subnet where the cluster should be deployed.
	SubnetID string `json:"subnetId,omitempty"`

	// Resource ID of the managed identity used by AKS. It is a mandatory parameter for Azure cluster creation.
	UserAssignedManagedIdentity string `json:"userAssignedManagedIdentity,omitempty"`
}

// Validate validates this azure activation options
func (m *AzureActivationOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutboundType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var azureActivationOptionsTypeOutboundTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LoadBalancer","UserAssignedNATGateway","UserDefinedRouting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureActivationOptionsTypeOutboundTypePropEnum = append(azureActivationOptionsTypeOutboundTypePropEnum, v)
	}
}

const (

	// AzureActivationOptionsOutboundTypeLoadBalancer captures enum value "LoadBalancer"
	AzureActivationOptionsOutboundTypeLoadBalancer string = "LoadBalancer"

	// AzureActivationOptionsOutboundTypeUserAssignedNATGateway captures enum value "UserAssignedNATGateway"
	AzureActivationOptionsOutboundTypeUserAssignedNATGateway string = "UserAssignedNATGateway"

	// AzureActivationOptionsOutboundTypeUserDefinedRouting captures enum value "UserDefinedRouting"
	AzureActivationOptionsOutboundTypeUserDefinedRouting string = "UserDefinedRouting"
)

// prop value enum
func (m *AzureActivationOptions) validateOutboundTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureActivationOptionsTypeOutboundTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureActivationOptions) validateOutboundType(formats strfmt.Registry) error {
	if swag.IsZero(m.OutboundType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutboundTypeEnum("outboundType", "body", m.OutboundType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure activation options based on context it is used
func (m *AzureActivationOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureActivationOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureActivationOptions) UnmarshalBinary(b []byte) error {
	var res AzureActivationOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

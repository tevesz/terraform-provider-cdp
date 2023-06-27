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

// CreatePrivateEnvironmentRequest Request object for a create private cloud environment request.
//
// swagger:model CreatePrivateEnvironmentRequest
type CreatePrivateEnvironmentRequest struct {

	// The address of the Cloudera Manager managing the Datalake cluster.
	// Required: true
	Address *string `json:"address"`

	// A string (text or json) used to authenticate to the Cloudera Manager.
	// Required: true
	AuthenticationToken *string `json:"authenticationToken"`

	// How to interpret the authenticationToken field. Defaults to CLEARTEXT_PASSWORD.
	// Enum: [CLEARTEXT_PASSWORD]
	AuthenticationTokenType string `json:"authenticationTokenType,omitempty"`

	// The name of the cluster(s) to use as a Datalake for the environment.
	// Required: true
	ClusterNames []string `json:"clusterNames"`

	// An description of the environment.
	Description string `json:"description,omitempty"`

	// docker pull secrets for the K8s cluster. This is expected to be a docker config json.
	DockerConfigJSON string `json:"dockerConfigJson,omitempty"`

	// docker user pass
	DockerUserPass *CreatePrivateEnvironmentRequestDockerUserPass `json:"dockerUserPass,omitempty"`

	// default domain suffix to work workload applications to use.
	Domain string `json:"domain,omitempty"`

	// The name of the environment. Must contain only lowercase letters, numbers and hyphens.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Name of credentials holding kubeconfig for access to the kubernetes cluster paired with this Environment.
	KubeConfig string `json:"kubeConfig,omitempty"`

	// Prefix for all namespaces created by Cloudera Data Platform within this cluster.
	NamespacePrefix *string `json:"namespacePrefix,omitempty"`

	// the K8s cluster type used for the environment.
	Platform *string `json:"platform,omitempty"`

	// An existing storage class on this kubernetes cluster. If not specified, the default storage class will be used.
	StorageClass string `json:"storageClass,omitempty"`

	// User name for accessing the Cloudera Manager.
	// Required: true
	User *string `json:"user"`
}

// Validate validates this create private environment request
func (m *CreatePrivateEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationTokenType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDockerUserPass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateAuthenticationToken(formats strfmt.Registry) error {

	if err := validate.Required("authenticationToken", "body", m.AuthenticationToken); err != nil {
		return err
	}

	return nil
}

var createPrivateEnvironmentRequestTypeAuthenticationTokenTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLEARTEXT_PASSWORD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createPrivateEnvironmentRequestTypeAuthenticationTokenTypePropEnum = append(createPrivateEnvironmentRequestTypeAuthenticationTokenTypePropEnum, v)
	}
}

const (

	// CreatePrivateEnvironmentRequestAuthenticationTokenTypeCLEARTEXTPASSWORD captures enum value "CLEARTEXT_PASSWORD"
	CreatePrivateEnvironmentRequestAuthenticationTokenTypeCLEARTEXTPASSWORD string = "CLEARTEXT_PASSWORD"
)

// prop value enum
func (m *CreatePrivateEnvironmentRequest) validateAuthenticationTokenTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createPrivateEnvironmentRequestTypeAuthenticationTokenTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateAuthenticationTokenType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationTokenType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTokenTypeEnum("authenticationTokenType", "body", m.AuthenticationTokenType); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateClusterNames(formats strfmt.Registry) error {

	if err := validate.Required("clusterNames", "body", m.ClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateDockerUserPass(formats strfmt.Registry) error {
	if swag.IsZero(m.DockerUserPass) { // not required
		return nil
	}

	if m.DockerUserPass != nil {
		if err := m.DockerUserPass.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerUserPass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dockerUserPass")
			}
			return err
		}
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequest) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create private environment request based on the context it is used
func (m *CreatePrivateEnvironmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDockerUserPass(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePrivateEnvironmentRequest) contextValidateDockerUserPass(ctx context.Context, formats strfmt.Registry) error {

	if m.DockerUserPass != nil {
		if err := m.DockerUserPass.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dockerUserPass")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dockerUserPass")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreatePrivateEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePrivateEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res CreatePrivateEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreatePrivateEnvironmentRequestDockerUserPass Alternative to dockerConfigJson.
//
// swagger:model CreatePrivateEnvironmentRequestDockerUserPass
type CreatePrivateEnvironmentRequestDockerUserPass struct {

	// Docker email.
	Email string `json:"email,omitempty"`

	// Docker password.
	// Required: true
	Password *string `json:"password"`

	// Docker Registry FQDN.
	// Required: true
	Server *string `json:"server"`

	// Docker username.
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this create private environment request docker user pass
func (m *CreatePrivateEnvironmentRequestDockerUserPass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreatePrivateEnvironmentRequestDockerUserPass) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("dockerUserPass"+"."+"password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequestDockerUserPass) validateServer(formats strfmt.Registry) error {

	if err := validate.Required("dockerUserPass"+"."+"server", "body", m.Server); err != nil {
		return err
	}

	return nil
}

func (m *CreatePrivateEnvironmentRequestDockerUserPass) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("dockerUserPass"+"."+"username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create private environment request docker user pass based on context it is used
func (m *CreatePrivateEnvironmentRequestDockerUserPass) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatePrivateEnvironmentRequestDockerUserPass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatePrivateEnvironmentRequestDockerUserPass) UnmarshalBinary(b []byte) error {
	var res CreatePrivateEnvironmentRequestDockerUserPass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
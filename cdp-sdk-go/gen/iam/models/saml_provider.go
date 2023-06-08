// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SamlProvider Information used to connect a CDP account to an external identity provider.
//
// swagger:model SamlProvider
type SamlProvider struct {

	// The Service Provider SAML metadata specific to this CDP SAML provider. This field will only be set for createSamlProvider and describeSamlProvider API calls.
	CdpSpMetadata string `json:"cdpSpMetadata,omitempty"`

	// The date when this SAML provider record was created.
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// CRN of the SAML provider in CDP.
	// Required: true
	Crn *string `json:"crn"`

	// Whether SCIM is enabled on this SAML provider. System for Cross-domain Identity Management (SCIM) version 2.0 is a standard for automating the provisioning of user and group identity information from identity provider to CDP. It is omitted for Cloudera for Government.
	EnableScim bool `json:"enableScim,omitempty"`

	// Whether to generate users' workload username by email or by identity provider user ID (SAML NameID).
	GenerateWorkloadUsernameByEmail bool `json:"generateWorkloadUsernameByEmail,omitempty"`

	// The original metadata that was passed while creating the SAML provider connector. This field will not be set when the SAML provider does not have metadata. This field will not be set for listSamlProviders API response.
	SamlMetadataDocument string `json:"samlMetadataDocument,omitempty"`

	// The unique ID of the saml provider.
	// Required: true
	SamlProviderID *string `json:"samlProviderId"`

	// Name of the SAML provider.
	// Required: true
	SamlProviderName *string `json:"samlProviderName"`

	// The SCIM URL if SCIM is enabled. It is omitted for Cloudera for Government.
	ScimURL string `json:"scimUrl,omitempty"`

	// Whether users federated with this SAML provider will have their group membership synchronized. Group membership can be passed using the https://cdp.cloudera.com/SAML/Attributes/groups SAML assertion.
	// Required: true
	SyncGroupsOnLogin *bool `json:"syncGroupsOnLogin"`
}

// Validate validates this saml provider
func (m *SamlProvider) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamlProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamlProviderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncGroupsOnLogin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlProvider) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SamlProvider) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *SamlProvider) validateSamlProviderID(formats strfmt.Registry) error {

	if err := validate.Required("samlProviderId", "body", m.SamlProviderID); err != nil {
		return err
	}

	return nil
}

func (m *SamlProvider) validateSamlProviderName(formats strfmt.Registry) error {

	if err := validate.Required("samlProviderName", "body", m.SamlProviderName); err != nil {
		return err
	}

	return nil
}

func (m *SamlProvider) validateSyncGroupsOnLogin(formats strfmt.Registry) error {

	if err := validate.Required("syncGroupsOnLogin", "body", m.SyncGroupsOnLogin); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this saml provider based on context it is used
func (m *SamlProvider) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SamlProvider) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlProvider) UnmarshalBinary(b []byte) error {
	var res SamlProvider
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

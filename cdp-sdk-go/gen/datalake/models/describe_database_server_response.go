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

// DescribeDatabaseServerResponse Response object for obtaining Database Server details.
//
// swagger:model DescribeDatabaseServerResponse
type DescribeDatabaseServerResponse struct {

	// CRN of the cluster of the database server
	ClusterCrn string `json:"clusterCrn,omitempty"`

	// Creation date of the database server
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// CRN of the database server
	Crn string `json:"crn,omitempty"`

	// Name of the database vendor (MYSQL, POSTGRES, ...)
	DatabaseVendor string `json:"databaseVendor,omitempty"`

	// Display name of the database vendor (MySQL, PostgreSQL, ...)
	DatabaseVendorDisplayName string `json:"databaseVendorDisplayName,omitempty"`

	// Description of the database server
	Description string `json:"description,omitempty"`

	// CRN of the environment of the database server
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// Host of the database server
	Host string `json:"host,omitempty"`

	// Name of the database server
	Name string `json:"name,omitempty"`

	// Port of the database server
	Port int32 `json:"port,omitempty"`

	// Ownership status of the database server
	// Enum: [UNKNOWN SERVICE_MANAGED USER_MANAGED]
	ResourceStatus string `json:"resourceStatus,omitempty"`

	// SSL configuration of the database server
	SslConfig *DatabaseServerSslConfig `json:"sslConfig,omitempty"`

	// Status of the database server stack
	// Enum: [REQUESTED CREATE_IN_PROGRESS AVAILABLE UPDATE_IN_PROGRESS UPDATE_REQUESTED UPDATE_FAILED CREATE_FAILED ENABLE_SECURITY_FAILED DELETE_REQUESTED PRE_DELETE_IN_PROGRESS DELETE_IN_PROGRESS DELETE_FAILED DELETE_COMPLETED STOPPED STOP_REQUESTED START_REQUESTED STOP_IN_PROGRESS START_IN_PROGRESS START_FAILED STOP_FAILED WAIT_FOR_SYNC MAINTENANCE_MODE_ENABLED UNKNOWN]
	Status string `json:"status,omitempty"`

	// Additional status information about the database server stack
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this describe database server response
func (m *DescribeDatabaseServerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSslConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDatabaseServerResponse) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var describeDatabaseServerResponseTypeResourceStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","SERVICE_MANAGED","USER_MANAGED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		describeDatabaseServerResponseTypeResourceStatusPropEnum = append(describeDatabaseServerResponseTypeResourceStatusPropEnum, v)
	}
}

const (

	// DescribeDatabaseServerResponseResourceStatusUNKNOWN captures enum value "UNKNOWN"
	DescribeDatabaseServerResponseResourceStatusUNKNOWN string = "UNKNOWN"

	// DescribeDatabaseServerResponseResourceStatusSERVICEMANAGED captures enum value "SERVICE_MANAGED"
	DescribeDatabaseServerResponseResourceStatusSERVICEMANAGED string = "SERVICE_MANAGED"

	// DescribeDatabaseServerResponseResourceStatusUSERMANAGED captures enum value "USER_MANAGED"
	DescribeDatabaseServerResponseResourceStatusUSERMANAGED string = "USER_MANAGED"
)

// prop value enum
func (m *DescribeDatabaseServerResponse) validateResourceStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, describeDatabaseServerResponseTypeResourceStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DescribeDatabaseServerResponse) validateResourceStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceStatusEnum("resourceStatus", "body", m.ResourceStatus); err != nil {
		return err
	}

	return nil
}

func (m *DescribeDatabaseServerResponse) validateSslConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SslConfig) { // not required
		return nil
	}

	if m.SslConfig != nil {
		if err := m.SslConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sslConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sslConfig")
			}
			return err
		}
	}

	return nil
}

var describeDatabaseServerResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","DELETE_REQUESTED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC","MAINTENANCE_MODE_ENABLED","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		describeDatabaseServerResponseTypeStatusPropEnum = append(describeDatabaseServerResponseTypeStatusPropEnum, v)
	}
}

const (

	// DescribeDatabaseServerResponseStatusREQUESTED captures enum value "REQUESTED"
	DescribeDatabaseServerResponseStatusREQUESTED string = "REQUESTED"

	// DescribeDatabaseServerResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusAVAILABLE captures enum value "AVAILABLE"
	DescribeDatabaseServerResponseStatusAVAILABLE string = "AVAILABLE"

	// DescribeDatabaseServerResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	DescribeDatabaseServerResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"

	// DescribeDatabaseServerResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	DescribeDatabaseServerResponseStatusUPDATEFAILED string = "UPDATE_FAILED"

	// DescribeDatabaseServerResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	DescribeDatabaseServerResponseStatusCREATEFAILED string = "CREATE_FAILED"

	// DescribeDatabaseServerResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	DescribeDatabaseServerResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"

	// DescribeDatabaseServerResponseStatusDELETEREQUESTED captures enum value "DELETE_REQUESTED"
	DescribeDatabaseServerResponseStatusDELETEREQUESTED string = "DELETE_REQUESTED"

	// DescribeDatabaseServerResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	DescribeDatabaseServerResponseStatusDELETEFAILED string = "DELETE_FAILED"

	// DescribeDatabaseServerResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	DescribeDatabaseServerResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"

	// DescribeDatabaseServerResponseStatusSTOPPED captures enum value "STOPPED"
	DescribeDatabaseServerResponseStatusSTOPPED string = "STOPPED"

	// DescribeDatabaseServerResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	DescribeDatabaseServerResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"

	// DescribeDatabaseServerResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	DescribeDatabaseServerResponseStatusSTARTREQUESTED string = "START_REQUESTED"

	// DescribeDatabaseServerResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	DescribeDatabaseServerResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"

	// DescribeDatabaseServerResponseStatusSTARTFAILED captures enum value "START_FAILED"
	DescribeDatabaseServerResponseStatusSTARTFAILED string = "START_FAILED"

	// DescribeDatabaseServerResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	DescribeDatabaseServerResponseStatusSTOPFAILED string = "STOP_FAILED"

	// DescribeDatabaseServerResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	DescribeDatabaseServerResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"

	// DescribeDatabaseServerResponseStatusMAINTENANCEMODEENABLED captures enum value "MAINTENANCE_MODE_ENABLED"
	DescribeDatabaseServerResponseStatusMAINTENANCEMODEENABLED string = "MAINTENANCE_MODE_ENABLED"

	// DescribeDatabaseServerResponseStatusUNKNOWN captures enum value "UNKNOWN"
	DescribeDatabaseServerResponseStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *DescribeDatabaseServerResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, describeDatabaseServerResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DescribeDatabaseServerResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this describe database server response based on the context it is used
func (m *DescribeDatabaseServerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSslConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDatabaseServerResponse) contextValidateSslConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SslConfig != nil {

		if swag.IsZero(m.SslConfig) { // not required
			return nil
		}

		if err := m.SslConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sslConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sslConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDatabaseServerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDatabaseServerResponse) UnmarshalBinary(b []byte) error {
	var res DescribeDatabaseServerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
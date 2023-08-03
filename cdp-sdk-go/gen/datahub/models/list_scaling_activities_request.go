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

// ListScalingActivitiesRequest Request object for List scaling Activities request.
//
// swagger:model ListScalingActivitiesRequest
type ListScalingActivitiesRequest struct {

	// The name or CRN of the cluster.
	// Required: true
	Cluster *string `json:"cluster"`

	// Duration in minutes for which we want all the scaling activities for the cluster. Either duration or start and end time needs to be provided.
	Duration int64 `json:"duration,omitempty"`

	// End time value in epoch millisecond until which we want to get all the scaling activities. Need to specify start time with it.
	EndTime int64 `json:"endTime,omitempty"`

	// Flag that decides whether to return only failed scaling activities or return all scaling activities in a given duration or a specific time interval.
	OnlyFailedScalingActivities bool `json:"onlyFailedScalingActivities,omitempty"`

	// The size of the page for getting scaling activities.
	// Maximum: 100
	// Minimum: 1
	PageSize int32 `json:"pageSize,omitempty"`

	// Start time value in epoch millisecond from which we want to get all the scaling activities. Need to specify end time with it.
	StartTime int64 `json:"startTime,omitempty"`

	// A token to specify where to start paginating. This is the nextToken from a previously truncated response.
	StartingToken string `json:"startingToken,omitempty"`
}

// Validate validates this list scaling activities request
func (m *ListScalingActivitiesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListScalingActivitiesRequest) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

func (m *ListScalingActivitiesRequest) validatePageSize(formats strfmt.Registry) error {
	if swag.IsZero(m.PageSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("pageSize", "body", int64(m.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "body", int64(m.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list scaling activities request based on context it is used
func (m *ListScalingActivitiesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListScalingActivitiesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListScalingActivitiesRequest) UnmarshalBinary(b []byte) error {
	var res ListScalingActivitiesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
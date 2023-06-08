// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// NewUpdateAutoScaleRulesParams creates a new UpdateAutoScaleRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAutoScaleRulesParams() *UpdateAutoScaleRulesParams {
	return &UpdateAutoScaleRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAutoScaleRulesParamsWithTimeout creates a new UpdateAutoScaleRulesParams object
// with the ability to set a timeout on a request.
func NewUpdateAutoScaleRulesParamsWithTimeout(timeout time.Duration) *UpdateAutoScaleRulesParams {
	return &UpdateAutoScaleRulesParams{
		timeout: timeout,
	}
}

// NewUpdateAutoScaleRulesParamsWithContext creates a new UpdateAutoScaleRulesParams object
// with the ability to set a context for a request.
func NewUpdateAutoScaleRulesParamsWithContext(ctx context.Context) *UpdateAutoScaleRulesParams {
	return &UpdateAutoScaleRulesParams{
		Context: ctx,
	}
}

// NewUpdateAutoScaleRulesParamsWithHTTPClient creates a new UpdateAutoScaleRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAutoScaleRulesParamsWithHTTPClient(client *http.Client) *UpdateAutoScaleRulesParams {
	return &UpdateAutoScaleRulesParams{
		HTTPClient: client,
	}
}

/*
UpdateAutoScaleRulesParams contains all the parameters to send to the API endpoint

	for the update auto scale rules operation.

	Typically these are written to a http.Request.
*/
type UpdateAutoScaleRulesParams struct {

	// Input.
	Input *models.UpdateAutoScaleRulesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update auto scale rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAutoScaleRulesParams) WithDefaults() *UpdateAutoScaleRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update auto scale rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAutoScaleRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) WithTimeout(timeout time.Duration) *UpdateAutoScaleRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) WithContext(ctx context.Context) *UpdateAutoScaleRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) WithHTTPClient(client *http.Client) *UpdateAutoScaleRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) WithInput(input *models.UpdateAutoScaleRulesRequest) *UpdateAutoScaleRulesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update auto scale rules params
func (o *UpdateAutoScaleRulesParams) SetInput(input *models.UpdateAutoScaleRulesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAutoScaleRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

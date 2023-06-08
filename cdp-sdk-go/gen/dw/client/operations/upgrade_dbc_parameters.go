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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewUpgradeDbcParams creates a new UpgradeDbcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeDbcParams() *UpgradeDbcParams {
	return &UpgradeDbcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeDbcParamsWithTimeout creates a new UpgradeDbcParams object
// with the ability to set a timeout on a request.
func NewUpgradeDbcParamsWithTimeout(timeout time.Duration) *UpgradeDbcParams {
	return &UpgradeDbcParams{
		timeout: timeout,
	}
}

// NewUpgradeDbcParamsWithContext creates a new UpgradeDbcParams object
// with the ability to set a context for a request.
func NewUpgradeDbcParamsWithContext(ctx context.Context) *UpgradeDbcParams {
	return &UpgradeDbcParams{
		Context: ctx,
	}
}

// NewUpgradeDbcParamsWithHTTPClient creates a new UpgradeDbcParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeDbcParamsWithHTTPClient(client *http.Client) *UpgradeDbcParams {
	return &UpgradeDbcParams{
		HTTPClient: client,
	}
}

/*
UpgradeDbcParams contains all the parameters to send to the API endpoint

	for the upgrade dbc operation.

	Typically these are written to a http.Request.
*/
type UpgradeDbcParams struct {

	// Input.
	Input *models.UpgradeDbcRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade dbc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeDbcParams) WithDefaults() *UpgradeDbcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade dbc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeDbcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade dbc params
func (o *UpgradeDbcParams) WithTimeout(timeout time.Duration) *UpgradeDbcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade dbc params
func (o *UpgradeDbcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade dbc params
func (o *UpgradeDbcParams) WithContext(ctx context.Context) *UpgradeDbcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade dbc params
func (o *UpgradeDbcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade dbc params
func (o *UpgradeDbcParams) WithHTTPClient(client *http.Client) *UpgradeDbcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade dbc params
func (o *UpgradeDbcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upgrade dbc params
func (o *UpgradeDbcParams) WithInput(input *models.UpgradeDbcRequest) *UpgradeDbcParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upgrade dbc params
func (o *UpgradeDbcParams) SetInput(input *models.UpgradeDbcRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeDbcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

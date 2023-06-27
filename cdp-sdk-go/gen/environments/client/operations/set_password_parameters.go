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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewSetPasswordParams creates a new SetPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetPasswordParams() *SetPasswordParams {
	return &SetPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetPasswordParamsWithTimeout creates a new SetPasswordParams object
// with the ability to set a timeout on a request.
func NewSetPasswordParamsWithTimeout(timeout time.Duration) *SetPasswordParams {
	return &SetPasswordParams{
		timeout: timeout,
	}
}

// NewSetPasswordParamsWithContext creates a new SetPasswordParams object
// with the ability to set a context for a request.
func NewSetPasswordParamsWithContext(ctx context.Context) *SetPasswordParams {
	return &SetPasswordParams{
		Context: ctx,
	}
}

// NewSetPasswordParamsWithHTTPClient creates a new SetPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetPasswordParamsWithHTTPClient(client *http.Client) *SetPasswordParams {
	return &SetPasswordParams{
		HTTPClient: client,
	}
}

/*
SetPasswordParams contains all the parameters to send to the API endpoint

	for the set password operation.

	Typically these are written to a http.Request.
*/
type SetPasswordParams struct {

	// Input.
	Input *models.SetPasswordRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPasswordParams) WithDefaults() *SetPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set password params
func (o *SetPasswordParams) WithTimeout(timeout time.Duration) *SetPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set password params
func (o *SetPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set password params
func (o *SetPasswordParams) WithContext(ctx context.Context) *SetPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set password params
func (o *SetPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set password params
func (o *SetPasswordParams) WithHTTPClient(client *http.Client) *SetPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set password params
func (o *SetPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the set password params
func (o *SetPasswordParams) WithInput(input *models.SetPasswordRequest) *SetPasswordParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the set password params
func (o *SetPasswordParams) SetInput(input *models.SetPasswordRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SetPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
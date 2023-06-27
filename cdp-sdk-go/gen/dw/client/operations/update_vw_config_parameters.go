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

// NewUpdateVwConfigParams creates a new UpdateVwConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVwConfigParams() *UpdateVwConfigParams {
	return &UpdateVwConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVwConfigParamsWithTimeout creates a new UpdateVwConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateVwConfigParamsWithTimeout(timeout time.Duration) *UpdateVwConfigParams {
	return &UpdateVwConfigParams{
		timeout: timeout,
	}
}

// NewUpdateVwConfigParamsWithContext creates a new UpdateVwConfigParams object
// with the ability to set a context for a request.
func NewUpdateVwConfigParamsWithContext(ctx context.Context) *UpdateVwConfigParams {
	return &UpdateVwConfigParams{
		Context: ctx,
	}
}

// NewUpdateVwConfigParamsWithHTTPClient creates a new UpdateVwConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVwConfigParamsWithHTTPClient(client *http.Client) *UpdateVwConfigParams {
	return &UpdateVwConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateVwConfigParams contains all the parameters to send to the API endpoint

	for the update vw config operation.

	Typically these are written to a http.Request.
*/
type UpdateVwConfigParams struct {

	// Input.
	Input *models.UpdateVwConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update vw config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVwConfigParams) WithDefaults() *UpdateVwConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update vw config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVwConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update vw config params
func (o *UpdateVwConfigParams) WithTimeout(timeout time.Duration) *UpdateVwConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update vw config params
func (o *UpdateVwConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update vw config params
func (o *UpdateVwConfigParams) WithContext(ctx context.Context) *UpdateVwConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update vw config params
func (o *UpdateVwConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update vw config params
func (o *UpdateVwConfigParams) WithHTTPClient(client *http.Client) *UpdateVwConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update vw config params
func (o *UpdateVwConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update vw config params
func (o *UpdateVwConfigParams) WithInput(input *models.UpdateVwConfigRequest) *UpdateVwConfigParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update vw config params
func (o *UpdateVwConfigParams) SetInput(input *models.UpdateVwConfigRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVwConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
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

// NewUpscaleFreeipaParams creates a new UpscaleFreeipaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpscaleFreeipaParams() *UpscaleFreeipaParams {
	return &UpscaleFreeipaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpscaleFreeipaParamsWithTimeout creates a new UpscaleFreeipaParams object
// with the ability to set a timeout on a request.
func NewUpscaleFreeipaParamsWithTimeout(timeout time.Duration) *UpscaleFreeipaParams {
	return &UpscaleFreeipaParams{
		timeout: timeout,
	}
}

// NewUpscaleFreeipaParamsWithContext creates a new UpscaleFreeipaParams object
// with the ability to set a context for a request.
func NewUpscaleFreeipaParamsWithContext(ctx context.Context) *UpscaleFreeipaParams {
	return &UpscaleFreeipaParams{
		Context: ctx,
	}
}

// NewUpscaleFreeipaParamsWithHTTPClient creates a new UpscaleFreeipaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpscaleFreeipaParamsWithHTTPClient(client *http.Client) *UpscaleFreeipaParams {
	return &UpscaleFreeipaParams{
		HTTPClient: client,
	}
}

/*
UpscaleFreeipaParams contains all the parameters to send to the API endpoint

	for the upscale freeipa operation.

	Typically these are written to a http.Request.
*/
type UpscaleFreeipaParams struct {

	// Input.
	Input *models.UpscaleFreeipaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upscale freeipa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpscaleFreeipaParams) WithDefaults() *UpscaleFreeipaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upscale freeipa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpscaleFreeipaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upscale freeipa params
func (o *UpscaleFreeipaParams) WithTimeout(timeout time.Duration) *UpscaleFreeipaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upscale freeipa params
func (o *UpscaleFreeipaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upscale freeipa params
func (o *UpscaleFreeipaParams) WithContext(ctx context.Context) *UpscaleFreeipaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upscale freeipa params
func (o *UpscaleFreeipaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upscale freeipa params
func (o *UpscaleFreeipaParams) WithHTTPClient(client *http.Client) *UpscaleFreeipaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upscale freeipa params
func (o *UpscaleFreeipaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upscale freeipa params
func (o *UpscaleFreeipaParams) WithInput(input *models.UpscaleFreeipaRequest) *UpscaleFreeipaParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upscale freeipa params
func (o *UpscaleFreeipaParams) SetInput(input *models.UpscaleFreeipaRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpscaleFreeipaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
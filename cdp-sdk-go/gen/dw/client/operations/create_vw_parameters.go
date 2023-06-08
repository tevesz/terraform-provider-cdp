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

// NewCreateVwParams creates a new CreateVwParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVwParams() *CreateVwParams {
	return &CreateVwParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVwParamsWithTimeout creates a new CreateVwParams object
// with the ability to set a timeout on a request.
func NewCreateVwParamsWithTimeout(timeout time.Duration) *CreateVwParams {
	return &CreateVwParams{
		timeout: timeout,
	}
}

// NewCreateVwParamsWithContext creates a new CreateVwParams object
// with the ability to set a context for a request.
func NewCreateVwParamsWithContext(ctx context.Context) *CreateVwParams {
	return &CreateVwParams{
		Context: ctx,
	}
}

// NewCreateVwParamsWithHTTPClient creates a new CreateVwParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVwParamsWithHTTPClient(client *http.Client) *CreateVwParams {
	return &CreateVwParams{
		HTTPClient: client,
	}
}

/*
CreateVwParams contains all the parameters to send to the API endpoint

	for the create vw operation.

	Typically these are written to a http.Request.
*/
type CreateVwParams struct {

	// Input.
	Input *models.CreateVwRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVwParams) WithDefaults() *CreateVwParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVwParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create vw params
func (o *CreateVwParams) WithTimeout(timeout time.Duration) *CreateVwParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vw params
func (o *CreateVwParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vw params
func (o *CreateVwParams) WithContext(ctx context.Context) *CreateVwParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vw params
func (o *CreateVwParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vw params
func (o *CreateVwParams) WithHTTPClient(client *http.Client) *CreateVwParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vw params
func (o *CreateVwParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create vw params
func (o *CreateVwParams) WithInput(input *models.CreateVwRequest) *CreateVwParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create vw params
func (o *CreateVwParams) SetInput(input *models.CreateVwRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVwParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

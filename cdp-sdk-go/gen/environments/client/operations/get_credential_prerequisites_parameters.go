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

// NewGetCredentialPrerequisitesParams creates a new GetCredentialPrerequisitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCredentialPrerequisitesParams() *GetCredentialPrerequisitesParams {
	return &GetCredentialPrerequisitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCredentialPrerequisitesParamsWithTimeout creates a new GetCredentialPrerequisitesParams object
// with the ability to set a timeout on a request.
func NewGetCredentialPrerequisitesParamsWithTimeout(timeout time.Duration) *GetCredentialPrerequisitesParams {
	return &GetCredentialPrerequisitesParams{
		timeout: timeout,
	}
}

// NewGetCredentialPrerequisitesParamsWithContext creates a new GetCredentialPrerequisitesParams object
// with the ability to set a context for a request.
func NewGetCredentialPrerequisitesParamsWithContext(ctx context.Context) *GetCredentialPrerequisitesParams {
	return &GetCredentialPrerequisitesParams{
		Context: ctx,
	}
}

// NewGetCredentialPrerequisitesParamsWithHTTPClient creates a new GetCredentialPrerequisitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCredentialPrerequisitesParamsWithHTTPClient(client *http.Client) *GetCredentialPrerequisitesParams {
	return &GetCredentialPrerequisitesParams{
		HTTPClient: client,
	}
}

/*
GetCredentialPrerequisitesParams contains all the parameters to send to the API endpoint

	for the get credential prerequisites operation.

	Typically these are written to a http.Request.
*/
type GetCredentialPrerequisitesParams struct {

	// Input.
	Input *models.GetCredentialPrerequisitesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get credential prerequisites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialPrerequisitesParams) WithDefaults() *GetCredentialPrerequisitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get credential prerequisites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCredentialPrerequisitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) WithTimeout(timeout time.Duration) *GetCredentialPrerequisitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) WithContext(ctx context.Context) *GetCredentialPrerequisitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) WithHTTPClient(client *http.Client) *GetCredentialPrerequisitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) WithInput(input *models.GetCredentialPrerequisitesRequest) *GetCredentialPrerequisitesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get credential prerequisites params
func (o *GetCredentialPrerequisitesParams) SetInput(input *models.GetCredentialPrerequisitesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetCredentialPrerequisitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

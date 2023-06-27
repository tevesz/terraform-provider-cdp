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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewRenewCertificateParams creates a new RenewCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRenewCertificateParams() *RenewCertificateParams {
	return &RenewCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRenewCertificateParamsWithTimeout creates a new RenewCertificateParams object
// with the ability to set a timeout on a request.
func NewRenewCertificateParamsWithTimeout(timeout time.Duration) *RenewCertificateParams {
	return &RenewCertificateParams{
		timeout: timeout,
	}
}

// NewRenewCertificateParamsWithContext creates a new RenewCertificateParams object
// with the ability to set a context for a request.
func NewRenewCertificateParamsWithContext(ctx context.Context) *RenewCertificateParams {
	return &RenewCertificateParams{
		Context: ctx,
	}
}

// NewRenewCertificateParamsWithHTTPClient creates a new RenewCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRenewCertificateParamsWithHTTPClient(client *http.Client) *RenewCertificateParams {
	return &RenewCertificateParams{
		HTTPClient: client,
	}
}

/*
RenewCertificateParams contains all the parameters to send to the API endpoint

	for the renew certificate operation.

	Typically these are written to a http.Request.
*/
type RenewCertificateParams struct {

	// Input.
	Input *models.RenewCertificateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the renew certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenewCertificateParams) WithDefaults() *RenewCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the renew certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RenewCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the renew certificate params
func (o *RenewCertificateParams) WithTimeout(timeout time.Duration) *RenewCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the renew certificate params
func (o *RenewCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the renew certificate params
func (o *RenewCertificateParams) WithContext(ctx context.Context) *RenewCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the renew certificate params
func (o *RenewCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the renew certificate params
func (o *RenewCertificateParams) WithHTTPClient(client *http.Client) *RenewCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the renew certificate params
func (o *RenewCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the renew certificate params
func (o *RenewCertificateParams) WithInput(input *models.RenewCertificateRequest) *RenewCertificateParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the renew certificate params
func (o *RenewCertificateParams) SetInput(input *models.RenewCertificateRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RenewCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
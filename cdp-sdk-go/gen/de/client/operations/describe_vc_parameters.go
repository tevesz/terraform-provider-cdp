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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// NewDescribeVcParams creates a new DescribeVcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeVcParams() *DescribeVcParams {
	return &DescribeVcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeVcParamsWithTimeout creates a new DescribeVcParams object
// with the ability to set a timeout on a request.
func NewDescribeVcParamsWithTimeout(timeout time.Duration) *DescribeVcParams {
	return &DescribeVcParams{
		timeout: timeout,
	}
}

// NewDescribeVcParamsWithContext creates a new DescribeVcParams object
// with the ability to set a context for a request.
func NewDescribeVcParamsWithContext(ctx context.Context) *DescribeVcParams {
	return &DescribeVcParams{
		Context: ctx,
	}
}

// NewDescribeVcParamsWithHTTPClient creates a new DescribeVcParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeVcParamsWithHTTPClient(client *http.Client) *DescribeVcParams {
	return &DescribeVcParams{
		HTTPClient: client,
	}
}

/*
DescribeVcParams contains all the parameters to send to the API endpoint

	for the describe vc operation.

	Typically these are written to a http.Request.
*/
type DescribeVcParams struct {

	// Input.
	Input *models.DescribeVcRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe vc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeVcParams) WithDefaults() *DescribeVcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe vc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeVcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe vc params
func (o *DescribeVcParams) WithTimeout(timeout time.Duration) *DescribeVcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe vc params
func (o *DescribeVcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe vc params
func (o *DescribeVcParams) WithContext(ctx context.Context) *DescribeVcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe vc params
func (o *DescribeVcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe vc params
func (o *DescribeVcParams) WithHTTPClient(client *http.Client) *DescribeVcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe vc params
func (o *DescribeVcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the describe vc params
func (o *DescribeVcParams) WithInput(input *models.DescribeVcRequest) *DescribeVcParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the describe vc params
func (o *DescribeVcParams) SetInput(input *models.DescribeVcRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeVcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// NewAddInstanceGroupsMlServingAppParams creates a new AddInstanceGroupsMlServingAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddInstanceGroupsMlServingAppParams() *AddInstanceGroupsMlServingAppParams {
	return &AddInstanceGroupsMlServingAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddInstanceGroupsMlServingAppParamsWithTimeout creates a new AddInstanceGroupsMlServingAppParams object
// with the ability to set a timeout on a request.
func NewAddInstanceGroupsMlServingAppParamsWithTimeout(timeout time.Duration) *AddInstanceGroupsMlServingAppParams {
	return &AddInstanceGroupsMlServingAppParams{
		timeout: timeout,
	}
}

// NewAddInstanceGroupsMlServingAppParamsWithContext creates a new AddInstanceGroupsMlServingAppParams object
// with the ability to set a context for a request.
func NewAddInstanceGroupsMlServingAppParamsWithContext(ctx context.Context) *AddInstanceGroupsMlServingAppParams {
	return &AddInstanceGroupsMlServingAppParams{
		Context: ctx,
	}
}

// NewAddInstanceGroupsMlServingAppParamsWithHTTPClient creates a new AddInstanceGroupsMlServingAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddInstanceGroupsMlServingAppParamsWithHTTPClient(client *http.Client) *AddInstanceGroupsMlServingAppParams {
	return &AddInstanceGroupsMlServingAppParams{
		HTTPClient: client,
	}
}

/*
AddInstanceGroupsMlServingAppParams contains all the parameters to send to the API endpoint

	for the add instance groups ml serving app operation.

	Typically these are written to a http.Request.
*/
type AddInstanceGroupsMlServingAppParams struct {

	// Input.
	Input *models.AddInstanceGroupsMlServingAppRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add instance groups ml serving app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddInstanceGroupsMlServingAppParams) WithDefaults() *AddInstanceGroupsMlServingAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add instance groups ml serving app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddInstanceGroupsMlServingAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) WithTimeout(timeout time.Duration) *AddInstanceGroupsMlServingAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) WithContext(ctx context.Context) *AddInstanceGroupsMlServingAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) WithHTTPClient(client *http.Client) *AddInstanceGroupsMlServingAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) WithInput(input *models.AddInstanceGroupsMlServingAppRequest) *AddInstanceGroupsMlServingAppParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the add instance groups ml serving app params
func (o *AddInstanceGroupsMlServingAppParams) SetInput(input *models.AddInstanceGroupsMlServingAppRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *AddInstanceGroupsMlServingAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
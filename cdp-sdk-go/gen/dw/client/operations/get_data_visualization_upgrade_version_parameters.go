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

// NewGetDataVisualizationUpgradeVersionParams creates a new GetDataVisualizationUpgradeVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataVisualizationUpgradeVersionParams() *GetDataVisualizationUpgradeVersionParams {
	return &GetDataVisualizationUpgradeVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataVisualizationUpgradeVersionParamsWithTimeout creates a new GetDataVisualizationUpgradeVersionParams object
// with the ability to set a timeout on a request.
func NewGetDataVisualizationUpgradeVersionParamsWithTimeout(timeout time.Duration) *GetDataVisualizationUpgradeVersionParams {
	return &GetDataVisualizationUpgradeVersionParams{
		timeout: timeout,
	}
}

// NewGetDataVisualizationUpgradeVersionParamsWithContext creates a new GetDataVisualizationUpgradeVersionParams object
// with the ability to set a context for a request.
func NewGetDataVisualizationUpgradeVersionParamsWithContext(ctx context.Context) *GetDataVisualizationUpgradeVersionParams {
	return &GetDataVisualizationUpgradeVersionParams{
		Context: ctx,
	}
}

// NewGetDataVisualizationUpgradeVersionParamsWithHTTPClient creates a new GetDataVisualizationUpgradeVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataVisualizationUpgradeVersionParamsWithHTTPClient(client *http.Client) *GetDataVisualizationUpgradeVersionParams {
	return &GetDataVisualizationUpgradeVersionParams{
		HTTPClient: client,
	}
}

/*
GetDataVisualizationUpgradeVersionParams contains all the parameters to send to the API endpoint

	for the get data visualization upgrade version operation.

	Typically these are written to a http.Request.
*/
type GetDataVisualizationUpgradeVersionParams struct {

	// Input.
	Input *models.GetDataVisualizationUpgradeVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data visualization upgrade version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataVisualizationUpgradeVersionParams) WithDefaults() *GetDataVisualizationUpgradeVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data visualization upgrade version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataVisualizationUpgradeVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) WithTimeout(timeout time.Duration) *GetDataVisualizationUpgradeVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) WithContext(ctx context.Context) *GetDataVisualizationUpgradeVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) WithHTTPClient(client *http.Client) *GetDataVisualizationUpgradeVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) WithInput(input *models.GetDataVisualizationUpgradeVersionRequest) *GetDataVisualizationUpgradeVersionParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get data visualization upgrade version params
func (o *GetDataVisualizationUpgradeVersionParams) SetInput(input *models.GetDataVisualizationUpgradeVersionRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataVisualizationUpgradeVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

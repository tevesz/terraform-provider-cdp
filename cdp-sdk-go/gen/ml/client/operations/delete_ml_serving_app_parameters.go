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

// NewDeleteMlServingAppParams creates a new DeleteMlServingAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMlServingAppParams() *DeleteMlServingAppParams {
	return &DeleteMlServingAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMlServingAppParamsWithTimeout creates a new DeleteMlServingAppParams object
// with the ability to set a timeout on a request.
func NewDeleteMlServingAppParamsWithTimeout(timeout time.Duration) *DeleteMlServingAppParams {
	return &DeleteMlServingAppParams{
		timeout: timeout,
	}
}

// NewDeleteMlServingAppParamsWithContext creates a new DeleteMlServingAppParams object
// with the ability to set a context for a request.
func NewDeleteMlServingAppParamsWithContext(ctx context.Context) *DeleteMlServingAppParams {
	return &DeleteMlServingAppParams{
		Context: ctx,
	}
}

// NewDeleteMlServingAppParamsWithHTTPClient creates a new DeleteMlServingAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMlServingAppParamsWithHTTPClient(client *http.Client) *DeleteMlServingAppParams {
	return &DeleteMlServingAppParams{
		HTTPClient: client,
	}
}

/*
DeleteMlServingAppParams contains all the parameters to send to the API endpoint

	for the delete ml serving app operation.

	Typically these are written to a http.Request.
*/
type DeleteMlServingAppParams struct {

	// Input.
	Input *models.DeleteMlServingAppRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete ml serving app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMlServingAppParams) WithDefaults() *DeleteMlServingAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete ml serving app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMlServingAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete ml serving app params
func (o *DeleteMlServingAppParams) WithTimeout(timeout time.Duration) *DeleteMlServingAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete ml serving app params
func (o *DeleteMlServingAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete ml serving app params
func (o *DeleteMlServingAppParams) WithContext(ctx context.Context) *DeleteMlServingAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete ml serving app params
func (o *DeleteMlServingAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete ml serving app params
func (o *DeleteMlServingAppParams) WithHTTPClient(client *http.Client) *DeleteMlServingAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete ml serving app params
func (o *DeleteMlServingAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete ml serving app params
func (o *DeleteMlServingAppParams) WithInput(input *models.DeleteMlServingAppRequest) *DeleteMlServingAppParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete ml serving app params
func (o *DeleteMlServingAppParams) SetInput(input *models.DeleteMlServingAppRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMlServingAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
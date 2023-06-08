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

// NewUpdateOrchestratorStateParams creates a new UpdateOrchestratorStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrchestratorStateParams() *UpdateOrchestratorStateParams {
	return &UpdateOrchestratorStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrchestratorStateParamsWithTimeout creates a new UpdateOrchestratorStateParams object
// with the ability to set a timeout on a request.
func NewUpdateOrchestratorStateParamsWithTimeout(timeout time.Duration) *UpdateOrchestratorStateParams {
	return &UpdateOrchestratorStateParams{
		timeout: timeout,
	}
}

// NewUpdateOrchestratorStateParamsWithContext creates a new UpdateOrchestratorStateParams object
// with the ability to set a context for a request.
func NewUpdateOrchestratorStateParamsWithContext(ctx context.Context) *UpdateOrchestratorStateParams {
	return &UpdateOrchestratorStateParams{
		Context: ctx,
	}
}

// NewUpdateOrchestratorStateParamsWithHTTPClient creates a new UpdateOrchestratorStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrchestratorStateParamsWithHTTPClient(client *http.Client) *UpdateOrchestratorStateParams {
	return &UpdateOrchestratorStateParams{
		HTTPClient: client,
	}
}

/*
UpdateOrchestratorStateParams contains all the parameters to send to the API endpoint

	for the update orchestrator state operation.

	Typically these are written to a http.Request.
*/
type UpdateOrchestratorStateParams struct {

	// Input.
	Input *models.UpdateOrchestratorStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update orchestrator state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrchestratorStateParams) WithDefaults() *UpdateOrchestratorStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update orchestrator state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrchestratorStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) WithTimeout(timeout time.Duration) *UpdateOrchestratorStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) WithContext(ctx context.Context) *UpdateOrchestratorStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) WithHTTPClient(client *http.Client) *UpdateOrchestratorStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) WithInput(input *models.UpdateOrchestratorStateRequest) *UpdateOrchestratorStateParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update orchestrator state params
func (o *UpdateOrchestratorStateParams) SetInput(input *models.UpdateOrchestratorStateRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrchestratorStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

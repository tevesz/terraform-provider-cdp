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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewDeleteSSHPublicKeyParams creates a new DeleteSSHPublicKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSSHPublicKeyParams() *DeleteSSHPublicKeyParams {
	return &DeleteSSHPublicKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSSHPublicKeyParamsWithTimeout creates a new DeleteSSHPublicKeyParams object
// with the ability to set a timeout on a request.
func NewDeleteSSHPublicKeyParamsWithTimeout(timeout time.Duration) *DeleteSSHPublicKeyParams {
	return &DeleteSSHPublicKeyParams{
		timeout: timeout,
	}
}

// NewDeleteSSHPublicKeyParamsWithContext creates a new DeleteSSHPublicKeyParams object
// with the ability to set a context for a request.
func NewDeleteSSHPublicKeyParamsWithContext(ctx context.Context) *DeleteSSHPublicKeyParams {
	return &DeleteSSHPublicKeyParams{
		Context: ctx,
	}
}

// NewDeleteSSHPublicKeyParamsWithHTTPClient creates a new DeleteSSHPublicKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSSHPublicKeyParamsWithHTTPClient(client *http.Client) *DeleteSSHPublicKeyParams {
	return &DeleteSSHPublicKeyParams{
		HTTPClient: client,
	}
}

/*
DeleteSSHPublicKeyParams contains all the parameters to send to the API endpoint

	for the delete Ssh public key operation.

	Typically these are written to a http.Request.
*/
type DeleteSSHPublicKeyParams struct {

	// Input.
	Input *models.DeleteSSHPublicKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Ssh public key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSSHPublicKeyParams) WithDefaults() *DeleteSSHPublicKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Ssh public key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSSHPublicKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) WithTimeout(timeout time.Duration) *DeleteSSHPublicKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) WithContext(ctx context.Context) *DeleteSSHPublicKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) WithHTTPClient(client *http.Client) *DeleteSSHPublicKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) WithInput(input *models.DeleteSSHPublicKeyRequest) *DeleteSSHPublicKeyParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete Ssh public key params
func (o *DeleteSSHPublicKeyParams) SetInput(input *models.DeleteSSHPublicKeyRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSSHPublicKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

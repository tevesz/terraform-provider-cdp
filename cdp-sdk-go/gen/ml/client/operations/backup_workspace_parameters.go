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

// NewBackupWorkspaceParams creates a new BackupWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupWorkspaceParams() *BackupWorkspaceParams {
	return &BackupWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupWorkspaceParamsWithTimeout creates a new BackupWorkspaceParams object
// with the ability to set a timeout on a request.
func NewBackupWorkspaceParamsWithTimeout(timeout time.Duration) *BackupWorkspaceParams {
	return &BackupWorkspaceParams{
		timeout: timeout,
	}
}

// NewBackupWorkspaceParamsWithContext creates a new BackupWorkspaceParams object
// with the ability to set a context for a request.
func NewBackupWorkspaceParamsWithContext(ctx context.Context) *BackupWorkspaceParams {
	return &BackupWorkspaceParams{
		Context: ctx,
	}
}

// NewBackupWorkspaceParamsWithHTTPClient creates a new BackupWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupWorkspaceParamsWithHTTPClient(client *http.Client) *BackupWorkspaceParams {
	return &BackupWorkspaceParams{
		HTTPClient: client,
	}
}

/*
BackupWorkspaceParams contains all the parameters to send to the API endpoint

	for the backup workspace operation.

	Typically these are written to a http.Request.
*/
type BackupWorkspaceParams struct {

	// Input.
	Input *models.BackupWorkspaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupWorkspaceParams) WithDefaults() *BackupWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup workspace params
func (o *BackupWorkspaceParams) WithTimeout(timeout time.Duration) *BackupWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup workspace params
func (o *BackupWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup workspace params
func (o *BackupWorkspaceParams) WithContext(ctx context.Context) *BackupWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup workspace params
func (o *BackupWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup workspace params
func (o *BackupWorkspaceParams) WithHTTPClient(client *http.Client) *BackupWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup workspace params
func (o *BackupWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the backup workspace params
func (o *BackupWorkspaceParams) WithInput(input *models.BackupWorkspaceRequest) *BackupWorkspaceParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the backup workspace params
func (o *BackupWorkspaceParams) SetInput(input *models.BackupWorkspaceRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *BackupWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
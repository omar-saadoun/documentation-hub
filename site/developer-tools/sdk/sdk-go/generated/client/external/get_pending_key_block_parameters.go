// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPendingKeyBlockParams creates a new GetPendingKeyBlockParams object
// with the default values initialized.
func NewGetPendingKeyBlockParams() *GetPendingKeyBlockParams {

	return &GetPendingKeyBlockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPendingKeyBlockParamsWithTimeout creates a new GetPendingKeyBlockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPendingKeyBlockParamsWithTimeout(timeout time.Duration) *GetPendingKeyBlockParams {

	return &GetPendingKeyBlockParams{

		timeout: timeout,
	}
}

// NewGetPendingKeyBlockParamsWithContext creates a new GetPendingKeyBlockParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPendingKeyBlockParamsWithContext(ctx context.Context) *GetPendingKeyBlockParams {

	return &GetPendingKeyBlockParams{

		Context: ctx,
	}
}

// NewGetPendingKeyBlockParamsWithHTTPClient creates a new GetPendingKeyBlockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPendingKeyBlockParamsWithHTTPClient(client *http.Client) *GetPendingKeyBlockParams {

	return &GetPendingKeyBlockParams{
		HTTPClient: client,
	}
}

/*GetPendingKeyBlockParams contains all the parameters to send to the API endpoint
for the get pending key block operation typically these are written to a http.Request
*/
type GetPendingKeyBlockParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pending key block params
func (o *GetPendingKeyBlockParams) WithTimeout(timeout time.Duration) *GetPendingKeyBlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pending key block params
func (o *GetPendingKeyBlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pending key block params
func (o *GetPendingKeyBlockParams) WithContext(ctx context.Context) *GetPendingKeyBlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pending key block params
func (o *GetPendingKeyBlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pending key block params
func (o *GetPendingKeyBlockParams) WithHTTPClient(client *http.Client) *GetPendingKeyBlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pending key block params
func (o *GetPendingKeyBlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPendingKeyBlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

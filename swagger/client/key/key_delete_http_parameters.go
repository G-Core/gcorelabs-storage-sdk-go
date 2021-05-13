// Code generated by go-swagger; DO NOT EDIT.

package key

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
	"github.com/go-openapi/swag"
)

// NewKeyDeleteHTTPParams creates a new KeyDeleteHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKeyDeleteHTTPParams() *KeyDeleteHTTPParams {
	return &KeyDeleteHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKeyDeleteHTTPParamsWithTimeout creates a new KeyDeleteHTTPParams object
// with the ability to set a timeout on a request.
func NewKeyDeleteHTTPParamsWithTimeout(timeout time.Duration) *KeyDeleteHTTPParams {
	return &KeyDeleteHTTPParams{
		timeout: timeout,
	}
}

// NewKeyDeleteHTTPParamsWithContext creates a new KeyDeleteHTTPParams object
// with the ability to set a context for a request.
func NewKeyDeleteHTTPParamsWithContext(ctx context.Context) *KeyDeleteHTTPParams {
	return &KeyDeleteHTTPParams{
		Context: ctx,
	}
}

// NewKeyDeleteHTTPParamsWithHTTPClient creates a new KeyDeleteHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewKeyDeleteHTTPParamsWithHTTPClient(client *http.Client) *KeyDeleteHTTPParams {
	return &KeyDeleteHTTPParams{
		HTTPClient: client,
	}
}

/* KeyDeleteHTTPParams contains all the parameters to send to the API endpoint
   for the key delete Http operation.

   Typically these are written to a http.Request.
*/
type KeyDeleteHTTPParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the key delete Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyDeleteHTTPParams) WithDefaults() *KeyDeleteHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the key delete Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KeyDeleteHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the key delete Http params
func (o *KeyDeleteHTTPParams) WithTimeout(timeout time.Duration) *KeyDeleteHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the key delete Http params
func (o *KeyDeleteHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the key delete Http params
func (o *KeyDeleteHTTPParams) WithContext(ctx context.Context) *KeyDeleteHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the key delete Http params
func (o *KeyDeleteHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the key delete Http params
func (o *KeyDeleteHTTPParams) WithHTTPClient(client *http.Client) *KeyDeleteHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the key delete Http params
func (o *KeyDeleteHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the key delete Http params
func (o *KeyDeleteHTTPParams) WithID(id int64) *KeyDeleteHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the key delete Http params
func (o *KeyDeleteHTTPParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *KeyDeleteHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
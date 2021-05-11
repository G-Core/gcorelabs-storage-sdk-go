// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewStorageGetHTTPParams creates a new StorageGetHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageGetHTTPParams() *StorageGetHTTPParams {
	return &StorageGetHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageGetHTTPParamsWithTimeout creates a new StorageGetHTTPParams object
// with the ability to set a timeout on a request.
func NewStorageGetHTTPParamsWithTimeout(timeout time.Duration) *StorageGetHTTPParams {
	return &StorageGetHTTPParams{
		timeout: timeout,
	}
}

// NewStorageGetHTTPParamsWithContext creates a new StorageGetHTTPParams object
// with the ability to set a context for a request.
func NewStorageGetHTTPParamsWithContext(ctx context.Context) *StorageGetHTTPParams {
	return &StorageGetHTTPParams{
		Context: ctx,
	}
}

// NewStorageGetHTTPParamsWithHTTPClient creates a new StorageGetHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageGetHTTPParamsWithHTTPClient(client *http.Client) *StorageGetHTTPParams {
	return &StorageGetHTTPParams{
		HTTPClient: client,
	}
}

/* StorageGetHTTPParams contains all the parameters to send to the API endpoint
   for the storage get Http operation.

   Typically these are written to a http.Request.
*/
type StorageGetHTTPParams struct {

	// ID.
	//
	// Format: int64
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage get Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageGetHTTPParams) WithDefaults() *StorageGetHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage get Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageGetHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the storage get Http params
func (o *StorageGetHTTPParams) WithTimeout(timeout time.Duration) *StorageGetHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage get Http params
func (o *StorageGetHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage get Http params
func (o *StorageGetHTTPParams) WithContext(ctx context.Context) *StorageGetHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage get Http params
func (o *StorageGetHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage get Http params
func (o *StorageGetHTTPParams) WithHTTPClient(client *http.Client) *StorageGetHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage get Http params
func (o *StorageGetHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the storage get Http params
func (o *StorageGetHTTPParams) WithID(id int64) *StorageGetHTTPParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the storage get Http params
func (o *StorageGetHTTPParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *StorageGetHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package notifications

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
)

// NewEventsMetaHTTPParams creates a new EventsMetaHTTPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEventsMetaHTTPParams() *EventsMetaHTTPParams {
	return &EventsMetaHTTPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEventsMetaHTTPParamsWithTimeout creates a new EventsMetaHTTPParams object
// with the ability to set a timeout on a request.
func NewEventsMetaHTTPParamsWithTimeout(timeout time.Duration) *EventsMetaHTTPParams {
	return &EventsMetaHTTPParams{
		timeout: timeout,
	}
}

// NewEventsMetaHTTPParamsWithContext creates a new EventsMetaHTTPParams object
// with the ability to set a context for a request.
func NewEventsMetaHTTPParamsWithContext(ctx context.Context) *EventsMetaHTTPParams {
	return &EventsMetaHTTPParams{
		Context: ctx,
	}
}

// NewEventsMetaHTTPParamsWithHTTPClient creates a new EventsMetaHTTPParams object
// with the ability to set a custom HTTPClient for a request.
func NewEventsMetaHTTPParamsWithHTTPClient(client *http.Client) *EventsMetaHTTPParams {
	return &EventsMetaHTTPParams{
		HTTPClient: client,
	}
}

/* EventsMetaHTTPParams contains all the parameters to send to the API endpoint
   for the events meta Http operation.

   Typically these are written to a http.Request.
*/
type EventsMetaHTTPParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the events meta Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventsMetaHTTPParams) WithDefaults() *EventsMetaHTTPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the events meta Http params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EventsMetaHTTPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the events meta Http params
func (o *EventsMetaHTTPParams) WithTimeout(timeout time.Duration) *EventsMetaHTTPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the events meta Http params
func (o *EventsMetaHTTPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the events meta Http params
func (o *EventsMetaHTTPParams) WithContext(ctx context.Context) *EventsMetaHTTPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the events meta Http params
func (o *EventsMetaHTTPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the events meta Http params
func (o *EventsMetaHTTPParams) WithHTTPClient(client *http.Client) *EventsMetaHTTPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the events meta Http params
func (o *EventsMetaHTTPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *EventsMetaHTTPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

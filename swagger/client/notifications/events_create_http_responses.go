// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

// EventsCreateHTTPReader is a Reader for the EventsCreateHTTP structure.
type EventsCreateHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsCreateHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventsCreateHTTPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEventsCreateHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEventsCreateHTTPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEventsCreateHTTPConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEventsCreateHTTPOK creates a EventsCreateHTTPOK with default headers values
func NewEventsCreateHTTPOK() *EventsCreateHTTPOK {
	return &EventsCreateHTTPOK{}
}

/* EventsCreateHTTPOK describes a response with status code 200, with default header values.

EventsResponse
*/
type EventsCreateHTTPOK struct {
	Payload *models.Events
}

func (o *EventsCreateHTTPOK) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/events][%d] eventsCreateHttpOK  %+v", 200, o.Payload)
}
func (o *EventsCreateHTTPOK) GetPayload() *models.Events {
	return o.Payload
}

func (o *EventsCreateHTTPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsCreateHTTPBadRequest creates a EventsCreateHTTPBadRequest with default headers values
func NewEventsCreateHTTPBadRequest() *EventsCreateHTTPBadRequest {
	return &EventsCreateHTTPBadRequest{}
}

/* EventsCreateHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type EventsCreateHTTPBadRequest struct {
	Payload *models.ErrResponse
}

func (o *EventsCreateHTTPBadRequest) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/events][%d] eventsCreateHttpBadRequest  %+v", 400, o.Payload)
}
func (o *EventsCreateHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsCreateHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsCreateHTTPUnauthorized creates a EventsCreateHTTPUnauthorized with default headers values
func NewEventsCreateHTTPUnauthorized() *EventsCreateHTTPUnauthorized {
	return &EventsCreateHTTPUnauthorized{}
}

/* EventsCreateHTTPUnauthorized describes a response with status code 401, with default header values.

ErrResponse
*/
type EventsCreateHTTPUnauthorized struct {
	Payload *models.ErrResponse
}

func (o *EventsCreateHTTPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/events][%d] eventsCreateHttpUnauthorized  %+v", 401, o.Payload)
}
func (o *EventsCreateHTTPUnauthorized) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsCreateHTTPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsCreateHTTPConflict creates a EventsCreateHTTPConflict with default headers values
func NewEventsCreateHTTPConflict() *EventsCreateHTTPConflict {
	return &EventsCreateHTTPConflict{}
}

/* EventsCreateHTTPConflict describes a response with status code 409, with default header values.

ErrResponse
*/
type EventsCreateHTTPConflict struct {
	Payload *models.ErrResponse
}

func (o *EventsCreateHTTPConflict) Error() string {
	return fmt.Sprintf("[POST /notifications/v1/events][%d] eventsCreateHttpConflict  %+v", 409, o.Payload)
}
func (o *EventsCreateHTTPConflict) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsCreateHTTPConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

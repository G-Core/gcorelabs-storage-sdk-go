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

// EventsGetHTTPReader is a Reader for the EventsGetHTTP structure.
type EventsGetHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EventsGetHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEventsGetHTTPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEventsGetHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEventsGetHTTPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEventsGetHTTPConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEventsGetHTTPOK creates a EventsGetHTTPOK with default headers values
func NewEventsGetHTTPOK() *EventsGetHTTPOK {
	return &EventsGetHTTPOK{}
}

/* EventsGetHTTPOK describes a response with status code 200, with default header values.

EventsResponse
*/
type EventsGetHTTPOK struct {
	Payload *models.Events
}

func (o *EventsGetHTTPOK) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/events][%d] eventsGetHttpOK  %+v", 200, o.Payload)
}
func (o *EventsGetHTTPOK) GetPayload() *models.Events {
	return o.Payload
}

func (o *EventsGetHTTPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsGetHTTPBadRequest creates a EventsGetHTTPBadRequest with default headers values
func NewEventsGetHTTPBadRequest() *EventsGetHTTPBadRequest {
	return &EventsGetHTTPBadRequest{}
}

/* EventsGetHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type EventsGetHTTPBadRequest struct {
	Payload *models.ErrResponse
}

func (o *EventsGetHTTPBadRequest) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/events][%d] eventsGetHttpBadRequest  %+v", 400, o.Payload)
}
func (o *EventsGetHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsGetHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsGetHTTPUnauthorized creates a EventsGetHTTPUnauthorized with default headers values
func NewEventsGetHTTPUnauthorized() *EventsGetHTTPUnauthorized {
	return &EventsGetHTTPUnauthorized{}
}

/* EventsGetHTTPUnauthorized describes a response with status code 401, with default header values.

ErrResponse
*/
type EventsGetHTTPUnauthorized struct {
	Payload *models.ErrResponse
}

func (o *EventsGetHTTPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/events][%d] eventsGetHttpUnauthorized  %+v", 401, o.Payload)
}
func (o *EventsGetHTTPUnauthorized) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsGetHTTPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEventsGetHTTPConflict creates a EventsGetHTTPConflict with default headers values
func NewEventsGetHTTPConflict() *EventsGetHTTPConflict {
	return &EventsGetHTTPConflict{}
}

/* EventsGetHTTPConflict describes a response with status code 409, with default header values.

ErrResponse
*/
type EventsGetHTTPConflict struct {
	Payload *models.ErrResponse
}

func (o *EventsGetHTTPConflict) Error() string {
	return fmt.Sprintf("[GET /notifications/v1/events][%d] eventsGetHttpConflict  %+v", 409, o.Payload)
}
func (o *EventsGetHTTPConflict) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *EventsGetHTTPConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
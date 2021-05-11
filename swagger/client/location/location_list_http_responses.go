// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

// LocationListHTTPReader is a Reader for the LocationListHTTP structure.
type LocationListHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocationListHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocationListHTTPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLocationListHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLocationListHTTPOK creates a LocationListHTTPOK with default headers values
func NewLocationListHTTPOK() *LocationListHTTPOK {
	return &LocationListHTTPOK{}
}

/* LocationListHTTPOK describes a response with status code 200, with default header values.

clientLocationRes
*/
type LocationListHTTPOK struct {
	Payload []*models.ClientLocationRes
}

func (o *LocationListHTTPOK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/location][%d] locationListHttpOK  %+v", 200, o.Payload)
}
func (o *LocationListHTTPOK) GetPayload() []*models.ClientLocationRes {
	return o.Payload
}

func (o *LocationListHTTPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocationListHTTPBadRequest creates a LocationListHTTPBadRequest with default headers values
func NewLocationListHTTPBadRequest() *LocationListHTTPBadRequest {
	return &LocationListHTTPBadRequest{}
}

/* LocationListHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type LocationListHTTPBadRequest struct {
	Payload *models.ErrResponse
}

func (o *LocationListHTTPBadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/location][%d] locationListHttpBadRequest  %+v", 400, o.Payload)
}
func (o *LocationListHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *LocationListHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

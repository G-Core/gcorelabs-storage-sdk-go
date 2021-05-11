// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

// StorageUpdateHTTPReader is a Reader for the StorageUpdateHTTP structure.
type StorageUpdateHTTPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageUpdateHTTPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageUpdateHTTPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageUpdateHTTPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStorageUpdateHTTPOK creates a StorageUpdateHTTPOK with default headers values
func NewStorageUpdateHTTPOK() *StorageUpdateHTTPOK {
	return &StorageUpdateHTTPOK{}
}

/* StorageUpdateHTTPOK describes a response with status code 200, with default header values.

Storage
*/
type StorageUpdateHTTPOK struct {
	Payload *models.Storage
}

func (o *StorageUpdateHTTPOK) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/storage/{id}][%d] storageUpdateHttpOK  %+v", 200, o.Payload)
}
func (o *StorageUpdateHTTPOK) GetPayload() *models.Storage {
	return o.Payload
}

func (o *StorageUpdateHTTPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Storage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageUpdateHTTPBadRequest creates a StorageUpdateHTTPBadRequest with default headers values
func NewStorageUpdateHTTPBadRequest() *StorageUpdateHTTPBadRequest {
	return &StorageUpdateHTTPBadRequest{}
}

/* StorageUpdateHTTPBadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageUpdateHTTPBadRequest struct {
	Payload *models.ErrResponse
}

func (o *StorageUpdateHTTPBadRequest) Error() string {
	return fmt.Sprintf("[POST /provisioning/v1/storage/{id}][%d] storageUpdateHttpBadRequest  %+v", 400, o.Payload)
}
func (o *StorageUpdateHTTPBadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageUpdateHTTPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StorageUpdateHTTPBody storage update HTTP body
swagger:model StorageUpdateHTTPBody
*/
type StorageUpdateHTTPBody struct {

	// expires
	Expires string `json:"expires,omitempty"`

	// server alias
	ServerAlias string `json:"server_alias,omitempty"`
}

// Validate validates this storage update HTTP body
func (o *StorageUpdateHTTPBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage update HTTP body based on context it is used
func (o *StorageUpdateHTTPBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StorageUpdateHTTPBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StorageUpdateHTTPBody) UnmarshalBinary(b []byte) error {
	var res StorageUpdateHTTPBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

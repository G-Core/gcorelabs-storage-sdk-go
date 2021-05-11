// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/G-Core/gcorelabs-storage-sdk-go/swagger/models"
)

// StorageListHTTPV1Reader is a Reader for the StorageListHTTPV1 structure.
type StorageListHTTPV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageListHTTPV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageListHTTPV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStorageListHTTPV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStorageListHTTPV1OK creates a StorageListHTTPV1OK with default headers values
func NewStorageListHTTPV1OK() *StorageListHTTPV1OK {
	return &StorageListHTTPV1OK{}
}

/* StorageListHTTPV1OK describes a response with status code 200, with default header values.

Storage
*/
type StorageListHTTPV1OK struct {
	Payload []*models.Storage
}

func (o *StorageListHTTPV1OK) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/storage][%d] storageListHttpV1OK  %+v", 200, o.Payload)
}
func (o *StorageListHTTPV1OK) GetPayload() []*models.Storage {
	return o.Payload
}

func (o *StorageListHTTPV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageListHTTPV1BadRequest creates a StorageListHTTPV1BadRequest with default headers values
func NewStorageListHTTPV1BadRequest() *StorageListHTTPV1BadRequest {
	return &StorageListHTTPV1BadRequest{}
}

/* StorageListHTTPV1BadRequest describes a response with status code 400, with default header values.

ErrResponse
*/
type StorageListHTTPV1BadRequest struct {
	Payload *models.ErrResponse
}

func (o *StorageListHTTPV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /provisioning/v1/storage][%d] storageListHttpV1BadRequest  %+v", 400, o.Payload)
}
func (o *StorageListHTTPV1BadRequest) GetPayload() *models.ErrResponse {
	return o.Payload
}

func (o *StorageListHTTPV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

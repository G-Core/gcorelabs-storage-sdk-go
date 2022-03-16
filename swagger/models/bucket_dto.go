// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BucketDto BucketDto for response
//
// swagger:model BucketDto
type BucketDto struct {

	// lifecycle
	Lifecycle int64 `json:"lifecycle,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this bucket dto
func (m *BucketDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bucket dto based on context it is used
func (m *BucketDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BucketDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BucketDto) UnmarshalBinary(b []byte) error {
	var res BucketDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventField EventField details
//
// swagger:model EventField
type EventField struct {

	// Name of field in root event structure
	// Example: storage_name
	Name string `json:"name,omitempty"`

	// Type of field
	// Example: number, bool, text
	Type string `json:"type,omitempty"`
}

// Validate validates this event field
func (m *EventField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event field based on context it is used
func (m *EventField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventField) UnmarshalBinary(b []byte) error {
	var res EventField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

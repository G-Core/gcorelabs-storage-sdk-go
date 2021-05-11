// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClientLocationRes client location res
//
// swagger:model clientLocationRes
type ClientLocationRes struct {

	// address
	Address string `json:"address,omitempty"`

	// allow for new storage
	// Enum: [deny allow]
	AllowForNewStorage string `json:"allow_for_new_storage,omitempty"`

	// Id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: s-ed1 / s-ws1 / ams / sin / fra / mia / etc.
	// Enum: [s-ed1 s-ws1 ams sin fra mia]
	Name string `json:"name,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this client location res
func (m *ClientLocationRes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowForNewStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clientLocationResTypeAllowForNewStoragePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","allow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientLocationResTypeAllowForNewStoragePropEnum = append(clientLocationResTypeAllowForNewStoragePropEnum, v)
	}
}

const (

	// ClientLocationResAllowForNewStorageDeny captures enum value "deny"
	ClientLocationResAllowForNewStorageDeny string = "deny"

	// ClientLocationResAllowForNewStorageAllow captures enum value "allow"
	ClientLocationResAllowForNewStorageAllow string = "allow"
)

// prop value enum
func (m *ClientLocationRes) validateAllowForNewStorageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientLocationResTypeAllowForNewStoragePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientLocationRes) validateAllowForNewStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.AllowForNewStorage) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllowForNewStorageEnum("allow_for_new_storage", "body", m.AllowForNewStorage); err != nil {
		return err
	}

	return nil
}

var clientLocationResTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["s-ed1","s-ws1","ams","sin","fra","mia"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientLocationResTypeNamePropEnum = append(clientLocationResTypeNamePropEnum, v)
	}
}

const (

	// ClientLocationResNameSDashEd1 captures enum value "s-ed1"
	ClientLocationResNameSDashEd1 string = "s-ed1"

	// ClientLocationResNameSDashWs1 captures enum value "s-ws1"
	ClientLocationResNameSDashWs1 string = "s-ws1"

	// ClientLocationResNameAms captures enum value "ams"
	ClientLocationResNameAms string = "ams"

	// ClientLocationResNameSin captures enum value "sin"
	ClientLocationResNameSin string = "sin"

	// ClientLocationResNameFra captures enum value "fra"
	ClientLocationResNameFra string = "fra"

	// ClientLocationResNameMia captures enum value "mia"
	ClientLocationResNameMia string = "mia"
)

// prop value enum
func (m *ClientLocationRes) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientLocationResTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientLocationRes) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this client location res based on context it is used
func (m *ClientLocationRes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientLocationRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientLocationRes) UnmarshalBinary(b []byte) error {
	var res ClientLocationRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

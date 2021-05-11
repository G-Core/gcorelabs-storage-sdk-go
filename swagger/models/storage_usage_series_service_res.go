// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StorageUsageSeriesServiceRes storage usage series service res
//
// swagger:model StorageUsageSeriesServiceRes
type StorageUsageSeriesServiceRes struct {

	// clients
	Clients map[string]ClientStats `json:"clients,omitempty"`
}

// Validate validates this storage usage series service res
func (m *StorageUsageSeriesServiceRes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUsageSeriesServiceRes) validateClients(formats strfmt.Registry) error {
	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	for k := range m.Clients {

		if err := validate.Required("clients"+"."+k, "body", m.Clients[k]); err != nil {
			return err
		}
		if val, ok := m.Clients[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this storage usage series service res based on the context it is used
func (m *StorageUsageSeriesServiceRes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageUsageSeriesServiceRes) contextValidateClients(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Clients {

		if val, ok := m.Clients[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageUsageSeriesServiceRes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageUsageSeriesServiceRes) UnmarshalBinary(b []byte) error {
	var res StorageUsageSeriesServiceRes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

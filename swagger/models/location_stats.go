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

// LocationStats location stats
//
// swagger:model LocationStats
type LocationStats struct {

	// name
	Name string `json:"name,omitempty"`

	// requests in sum
	RequestsInSum uint64 `json:"requests_in_sum,omitempty"`

	// requests out edges sum
	RequestsOutEdgesSum uint64 `json:"requests_out_edges_sum,omitempty"`

	// requests out wo edges sum
	RequestsOutWoEdgesSum uint64 `json:"requests_out_wo_edges_sum,omitempty"`

	// requests sum
	RequestsSum uint64 `json:"requests_sum,omitempty"`

	// size sum max
	SizeSumMax uint64 `json:"size_sum_max,omitempty"`

	// size sum mean
	SizeSumMean uint64 `json:"size_sum_mean,omitempty"`

	// storages
	Storages map[string]StorageStats `json:"storages,omitempty"`

	// traffic in sum
	TrafficInSum uint64 `json:"traffic_in_sum,omitempty"`

	// traffic out edges sum
	TrafficOutEdgesSum uint64 `json:"traffic_out_edges_sum,omitempty"`

	// traffic out wo edges sum
	TrafficOutWoEdgesSum uint64 `json:"traffic_out_wo_edges_sum,omitempty"`

	// traffic sum
	TrafficSum uint64 `json:"traffic_sum,omitempty"`
}

// Validate validates this location stats
func (m *LocationStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStorages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationStats) validateStorages(formats strfmt.Registry) error {
	if swag.IsZero(m.Storages) { // not required
		return nil
	}

	for k := range m.Storages {

		if err := validate.Required("storages"+"."+k, "body", m.Storages[k]); err != nil {
			return err
		}
		if val, ok := m.Storages[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this location stats based on the context it is used
func (m *LocationStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocationStats) contextValidateStorages(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Storages {

		if val, ok := m.Storages[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LocationStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationStats) UnmarshalBinary(b []byte) error {
	var res LocationStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserRecord A User Record
// swagger:model UserRecord
type UserRecord struct {

	// first name
	// Required: true
	FirstName *string `json:"firstName"`

	// groups
	// Required: true
	// Min Items: 1
	Groups []string `json:"groups"`

	// last name
	// Required: true
	LastName *string `json:"lastName"`

	// user Id
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this user record
func (m *UserRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRecord) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *UserRecord) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	iGroupsSize := int64(len(m.Groups))

	if err := validate.MinItems("groups", "body", iGroupsSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *UserRecord) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *UserRecord) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRecord) UnmarshalBinary(b []byte) error {
	var res UserRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

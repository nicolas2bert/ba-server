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

// User user
// swagger:model User
type User struct {

	// flickr secret token
	// Required: true
	FlickrSecretToken *string `json:"flickrSecretToken"`

	// flickr token
	// Required: true
	FlickrToken *string `json:"flickrToken"`

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlickrSecretToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlickrToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateFlickrSecretToken(formats strfmt.Registry) error {

	if err := validate.Required("flickrSecretToken", "body", m.FlickrSecretToken); err != nil {
		return err
	}

	return nil
}

func (m *User) validateFlickrToken(formats strfmt.Registry) error {

	if err := validate.Required("flickrToken", "body", m.FlickrToken); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RefreshTokenHookResponse RefreshTokenHookResponse is the response body received from the refresh token hook.
//
// swagger:model refreshTokenHookResponse
type RefreshTokenHookResponse struct {

	// session
	Session *ConsentRequestSession `json:"session,omitempty"`
}

// Validate validates this refresh token hook response
func (m *RefreshTokenHookResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshTokenHookResponse) validateSession(formats strfmt.Registry) error {
	if swag.IsZero(m.Session) { // not required
		return nil
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this refresh token hook response based on the context it is used
func (m *RefreshTokenHookResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshTokenHookResponse) contextValidateSession(ctx context.Context, formats strfmt.Registry) error {

	if m.Session != nil {
		if err := m.Session.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RefreshTokenHookResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefreshTokenHookResponse) UnmarshalBinary(b []byte) error {
	var res RefreshTokenHookResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

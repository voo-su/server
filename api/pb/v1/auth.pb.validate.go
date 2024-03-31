// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/auth.proto

package api_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AuthLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthLoginRequestMultiError, or nil if none found.
func (m *AuthLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Platform

	if len(errors) > 0 {
		return AuthLoginRequestMultiError(errors)
	}

	return nil
}

// AuthLoginRequestMultiError is an error wrapping multiple validation errors
// returned by AuthLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginRequestMultiError) AllErrors() []error { return m }

// AuthLoginRequestValidationError is the validation error returned by
// AuthLoginRequest.Validate if the designated constraints aren't met.
type AuthLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginRequestValidationError) ErrorName() string { return "AuthLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginRequestValidationError{}

// Validate checks the field values on AuthLoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthLoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthLoginResponseMultiError, or nil if none found.
func (m *AuthLoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for ExpiresIn

	if len(errors) > 0 {
		return AuthLoginResponseMultiError(errors)
	}

	return nil
}

// AuthLoginResponseMultiError is an error wrapping multiple validation errors
// returned by AuthLoginResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginResponseMultiError) AllErrors() []error { return m }

// AuthLoginResponseValidationError is the validation error returned by
// AuthLoginResponse.Validate if the designated constraints aren't met.
type AuthLoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginResponseValidationError) ErrorName() string {
	return "AuthLoginResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthLoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginResponseValidationError{}

// Validate checks the field values on AuthVerifyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthVerifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthVerifyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthVerifyRequestMultiError, or nil if none found.
func (m *AuthVerifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthVerifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for Code

	if len(errors) > 0 {
		return AuthVerifyRequestMultiError(errors)
	}

	return nil
}

// AuthVerifyRequestMultiError is an error wrapping multiple validation errors
// returned by AuthVerifyRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthVerifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthVerifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthVerifyRequestMultiError) AllErrors() []error { return m }

// AuthVerifyRequestValidationError is the validation error returned by
// AuthVerifyRequest.Validate if the designated constraints aren't met.
type AuthVerifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthVerifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthVerifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthVerifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthVerifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthVerifyRequestValidationError) ErrorName() string {
	return "AuthVerifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthVerifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthVerifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthVerifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthVerifyRequestValidationError{}

// Validate checks the field values on AuthVerifyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthVerifyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthVerifyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthVerifyResponseMultiError, or nil if none found.
func (m *AuthVerifyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthVerifyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for AccessToken

	// no validation rules for ExpiresIn

	if len(errors) > 0 {
		return AuthVerifyResponseMultiError(errors)
	}

	return nil
}

// AuthVerifyResponseMultiError is an error wrapping multiple validation errors
// returned by AuthVerifyResponse.ValidateAll() if the designated constraints
// aren't met.
type AuthVerifyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthVerifyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthVerifyResponseMultiError) AllErrors() []error { return m }

// AuthVerifyResponseValidationError is the validation error returned by
// AuthVerifyResponse.Validate if the designated constraints aren't met.
type AuthVerifyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthVerifyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthVerifyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthVerifyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthVerifyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthVerifyResponseValidationError) ErrorName() string {
	return "AuthVerifyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthVerifyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthVerifyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthVerifyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthVerifyResponseValidationError{}

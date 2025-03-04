// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager/dashboard.proto

package manager_pb

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

// Validate checks the field values on ManagerDashboardRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ManagerDashboardRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ManagerDashboardRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ManagerDashboardRequestMultiError, or nil if none found.
func (m *ManagerDashboardRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ManagerDashboardRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ManagerDashboardRequestMultiError(errors)
	}

	return nil
}

// ManagerDashboardRequestMultiError is an error wrapping multiple validation
// errors returned by ManagerDashboardRequest.ValidateAll() if the designated
// constraints aren't met.
type ManagerDashboardRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ManagerDashboardRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ManagerDashboardRequestMultiError) AllErrors() []error { return m }

// ManagerDashboardRequestValidationError is the validation error returned by
// ManagerDashboardRequest.Validate if the designated constraints aren't met.
type ManagerDashboardRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ManagerDashboardRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ManagerDashboardRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ManagerDashboardRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ManagerDashboardRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ManagerDashboardRequestValidationError) ErrorName() string {
	return "ManagerDashboardRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ManagerDashboardRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sManagerDashboardRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ManagerDashboardRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ManagerDashboardRequestValidationError{}

// Validate checks the field values on ManagerDashboardResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ManagerDashboardResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ManagerDashboardResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ManagerDashboardResponseMultiError, or nil if none found.
func (m *ManagerDashboardResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ManagerDashboardResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Users

	// no validation rules for Bots

	// no validation rules for TotalMessages

	// no validation rules for GroupChats

	// no validation rules for GroupMessages

	// no validation rules for PrivateMessages

	if len(errors) > 0 {
		return ManagerDashboardResponseMultiError(errors)
	}

	return nil
}

// ManagerDashboardResponseMultiError is an error wrapping multiple validation
// errors returned by ManagerDashboardResponse.ValidateAll() if the designated
// constraints aren't met.
type ManagerDashboardResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ManagerDashboardResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ManagerDashboardResponseMultiError) AllErrors() []error { return m }

// ManagerDashboardResponseValidationError is the validation error returned by
// ManagerDashboardResponse.Validate if the designated constraints aren't met.
type ManagerDashboardResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ManagerDashboardResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ManagerDashboardResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ManagerDashboardResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ManagerDashboardResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ManagerDashboardResponseValidationError) ErrorName() string {
	return "ManagerDashboardResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ManagerDashboardResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sManagerDashboardResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ManagerDashboardResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ManagerDashboardResponseValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bot/message.proto

package bot_pb

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

// Validate checks the field values on MessageSendRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageSendRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageSendRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageSendRequestMultiError, or nil if none found.
func (m *MessageSendRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageSendRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	// no validation rules for Text

	if len(errors) > 0 {
		return MessageSendRequestMultiError(errors)
	}

	return nil
}

// MessageSendRequestMultiError is an error wrapping multiple validation errors
// returned by MessageSendRequest.ValidateAll() if the designated constraints
// aren't met.
type MessageSendRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageSendRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageSendRequestMultiError) AllErrors() []error { return m }

// MessageSendRequestValidationError is the validation error returned by
// MessageSendRequest.Validate if the designated constraints aren't met.
type MessageSendRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageSendRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageSendRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageSendRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageSendRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageSendRequestValidationError) ErrorName() string {
	return "MessageSendRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MessageSendRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageSendRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageSendRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageSendRequestValidationError{}

// Validate checks the field values on MessageSendResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageSendResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageSendResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageSendResponseMultiError, or nil if none found.
func (m *MessageSendResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageSendResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MessageSendResponseMultiError(errors)
	}

	return nil
}

// MessageSendResponseMultiError is an error wrapping multiple validation
// errors returned by MessageSendResponse.ValidateAll() if the designated
// constraints aren't met.
type MessageSendResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageSendResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageSendResponseMultiError) AllErrors() []error { return m }

// MessageSendResponseValidationError is the validation error returned by
// MessageSendResponse.Validate if the designated constraints aren't met.
type MessageSendResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageSendResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageSendResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageSendResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageSendResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageSendResponseValidationError) ErrorName() string {
	return "MessageSendResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MessageSendResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageSendResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageSendResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageSendResponseValidationError{}

// Validate checks the field values on MessageChatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageChatsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageChatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageChatsRequestMultiError, or nil if none found.
func (m *MessageChatsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageChatsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return MessageChatsRequestMultiError(errors)
	}

	return nil
}

// MessageChatsRequestMultiError is an error wrapping multiple validation
// errors returned by MessageChatsRequest.ValidateAll() if the designated
// constraints aren't met.
type MessageChatsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageChatsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageChatsRequestMultiError) AllErrors() []error { return m }

// MessageChatsRequestValidationError is the validation error returned by
// MessageChatsRequest.Validate if the designated constraints aren't met.
type MessageChatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageChatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageChatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageChatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageChatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageChatsRequestValidationError) ErrorName() string {
	return "MessageChatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MessageChatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageChatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageChatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageChatsRequestValidationError{}

// Validate checks the field values on MessageChatsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageChatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageChatsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageChatsResponseMultiError, or nil if none found.
func (m *MessageChatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageChatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageChatsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageChatsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageChatsResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MessageChatsResponseMultiError(errors)
	}

	return nil
}

// MessageChatsResponseMultiError is an error wrapping multiple validation
// errors returned by MessageChatsResponse.ValidateAll() if the designated
// constraints aren't met.
type MessageChatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageChatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageChatsResponseMultiError) AllErrors() []error { return m }

// MessageChatsResponseValidationError is the validation error returned by
// MessageChatsResponse.Validate if the designated constraints aren't met.
type MessageChatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageChatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageChatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageChatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageChatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageChatsResponseValidationError) ErrorName() string {
	return "MessageChatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MessageChatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageChatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageChatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageChatsResponseValidationError{}

// Validate checks the field values on MessageChatsResponse_Item with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MessageChatsResponse_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageChatsResponse_Item with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageChatsResponse_ItemMultiError, or nil if none found.
func (m *MessageChatsResponse_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageChatsResponse_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return MessageChatsResponse_ItemMultiError(errors)
	}

	return nil
}

// MessageChatsResponse_ItemMultiError is an error wrapping multiple validation
// errors returned by MessageChatsResponse_Item.ValidateAll() if the
// designated constraints aren't met.
type MessageChatsResponse_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageChatsResponse_ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageChatsResponse_ItemMultiError) AllErrors() []error { return m }

// MessageChatsResponse_ItemValidationError is the validation error returned by
// MessageChatsResponse_Item.Validate if the designated constraints aren't met.
type MessageChatsResponse_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageChatsResponse_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageChatsResponse_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageChatsResponse_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageChatsResponse_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageChatsResponse_ItemValidationError) ErrorName() string {
	return "MessageChatsResponse_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e MessageChatsResponse_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageChatsResponse_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageChatsResponse_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageChatsResponse_ItemValidationError{}

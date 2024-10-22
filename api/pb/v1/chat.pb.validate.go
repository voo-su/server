// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/chat.proto

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

// Validate checks the field values on ChatCreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ChatCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatCreateRequestMultiError, or nil if none found.
func (m *ChatCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DialogType

	// no validation rules for ReceiverId

	if len(errors) > 0 {
		return ChatCreateRequestMultiError(errors)
	}

	return nil
}

// ChatCreateRequestMultiError is an error wrapping multiple validation errors
// returned by ChatCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type ChatCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatCreateRequestMultiError) AllErrors() []error { return m }

// ChatCreateRequestValidationError is the validation error returned by
// ChatCreateRequest.Validate if the designated constraints aren't met.
type ChatCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatCreateRequestValidationError) ErrorName() string {
	return "ChatCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatCreateRequestValidationError{}

// Validate checks the field values on ChatCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatCreateResponseMultiError, or nil if none found.
func (m *ChatCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for DialogType

	// no validation rules for ReceiverId

	// no validation rules for IsTop

	// no validation rules for IsDisturb

	// no validation rules for IsOnline

	// no validation rules for IsBot

	// no validation rules for Username

	// no validation rules for Name

	// no validation rules for Surname

	// no validation rules for Avatar

	// no validation rules for UnreadNum

	// no validation rules for MsgText

	// no validation rules for UpdatedAt

	// no validation rules for RemarkName

	if len(errors) > 0 {
		return ChatCreateResponseMultiError(errors)
	}

	return nil
}

// ChatCreateResponseMultiError is an error wrapping multiple validation errors
// returned by ChatCreateResponse.ValidateAll() if the designated constraints
// aren't met.
type ChatCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatCreateResponseMultiError) AllErrors() []error { return m }

// ChatCreateResponseValidationError is the validation error returned by
// ChatCreateResponse.Validate if the designated constraints aren't met.
type ChatCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatCreateResponseValidationError) ErrorName() string {
	return "ChatCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatCreateResponseValidationError{}

// Validate checks the field values on ChatItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatItemMultiError, or nil
// if none found.
func (m *ChatItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for DialogType

	// no validation rules for ReceiverId

	// no validation rules for Username

	// no validation rules for Avatar

	// no validation rules for Name

	// no validation rules for Surname

	// no validation rules for UnreadNum

	// no validation rules for MsgText

	// no validation rules for UpdatedAt

	// no validation rules for IsTop

	// no validation rules for IsDisturb

	// no validation rules for IsOnline

	// no validation rules for IsBot

	// no validation rules for Remark

	if len(errors) > 0 {
		return ChatItemMultiError(errors)
	}

	return nil
}

// ChatItemMultiError is an error wrapping multiple validation errors returned
// by ChatItem.ValidateAll() if the designated constraints aren't met.
type ChatItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatItemMultiError) AllErrors() []error { return m }

// ChatItemValidationError is the validation error returned by
// ChatItem.Validate if the designated constraints aren't met.
type ChatItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatItemValidationError) ErrorName() string { return "ChatItemValidationError" }

// Error satisfies the builtin error interface
func (e ChatItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatItemValidationError{}

// Validate checks the field values on ChatListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ChatListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatListResponseMultiError, or nil if none found.
func (m *ChatListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatListResponse) validate(all bool) error {
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
					errors = append(errors, ChatListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ChatListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ChatListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ChatListResponseMultiError(errors)
	}

	return nil
}

// ChatListResponseMultiError is an error wrapping multiple validation errors
// returned by ChatListResponse.ValidateAll() if the designated constraints
// aren't met.
type ChatListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatListResponseMultiError) AllErrors() []error { return m }

// ChatListResponseValidationError is the validation error returned by
// ChatListResponse.Validate if the designated constraints aren't met.
type ChatListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatListResponseValidationError) ErrorName() string { return "ChatListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ChatListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatListResponseValidationError{}

// Validate checks the field values on ChatDeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ChatDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatDeleteRequestMultiError, or nil if none found.
func (m *ChatDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ListId

	if len(errors) > 0 {
		return ChatDeleteRequestMultiError(errors)
	}

	return nil
}

// ChatDeleteRequestMultiError is an error wrapping multiple validation errors
// returned by ChatDeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type ChatDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatDeleteRequestMultiError) AllErrors() []error { return m }

// ChatDeleteRequestValidationError is the validation error returned by
// ChatDeleteRequest.Validate if the designated constraints aren't met.
type ChatDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatDeleteRequestValidationError) ErrorName() string {
	return "ChatDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatDeleteRequestValidationError{}

// Validate checks the field values on ChatDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatDeleteResponseMultiError, or nil if none found.
func (m *ChatDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatDeleteResponseMultiError(errors)
	}

	return nil
}

// ChatDeleteResponseMultiError is an error wrapping multiple validation errors
// returned by ChatDeleteResponse.ValidateAll() if the designated constraints
// aren't met.
type ChatDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatDeleteResponseMultiError) AllErrors() []error { return m }

// ChatDeleteResponseValidationError is the validation error returned by
// ChatDeleteResponse.Validate if the designated constraints aren't met.
type ChatDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatDeleteResponseValidationError) ErrorName() string {
	return "ChatDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatDeleteResponseValidationError{}

// Validate checks the field values on ChatTopRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatTopRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatTopRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatTopRequestMultiError,
// or nil if none found.
func (m *ChatTopRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatTopRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ListId

	// no validation rules for Type

	if len(errors) > 0 {
		return ChatTopRequestMultiError(errors)
	}

	return nil
}

// ChatTopRequestMultiError is an error wrapping multiple validation errors
// returned by ChatTopRequest.ValidateAll() if the designated constraints
// aren't met.
type ChatTopRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatTopRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatTopRequestMultiError) AllErrors() []error { return m }

// ChatTopRequestValidationError is the validation error returned by
// ChatTopRequest.Validate if the designated constraints aren't met.
type ChatTopRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatTopRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatTopRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatTopRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatTopRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatTopRequestValidationError) ErrorName() string { return "ChatTopRequestValidationError" }

// Error satisfies the builtin error interface
func (e ChatTopRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatTopRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatTopRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatTopRequestValidationError{}

// Validate checks the field values on ChatTopResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ChatTopResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatTopResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatTopResponseMultiError, or nil if none found.
func (m *ChatTopResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatTopResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatTopResponseMultiError(errors)
	}

	return nil
}

// ChatTopResponseMultiError is an error wrapping multiple validation errors
// returned by ChatTopResponse.ValidateAll() if the designated constraints
// aren't met.
type ChatTopResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatTopResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatTopResponseMultiError) AllErrors() []error { return m }

// ChatTopResponseValidationError is the validation error returned by
// ChatTopResponse.Validate if the designated constraints aren't met.
type ChatTopResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatTopResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatTopResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatTopResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatTopResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatTopResponseValidationError) ErrorName() string { return "ChatTopResponseValidationError" }

// Error satisfies the builtin error interface
func (e ChatTopResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatTopResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatTopResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatTopResponseValidationError{}

// Validate checks the field values on ChatDisturbRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatDisturbRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatDisturbRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatDisturbRequestMultiError, or nil if none found.
func (m *ChatDisturbRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatDisturbRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DialogType

	// no validation rules for ReceiverId

	// no validation rules for IsDisturb

	if len(errors) > 0 {
		return ChatDisturbRequestMultiError(errors)
	}

	return nil
}

// ChatDisturbRequestMultiError is an error wrapping multiple validation errors
// returned by ChatDisturbRequest.ValidateAll() if the designated constraints
// aren't met.
type ChatDisturbRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatDisturbRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatDisturbRequestMultiError) AllErrors() []error { return m }

// ChatDisturbRequestValidationError is the validation error returned by
// ChatDisturbRequest.Validate if the designated constraints aren't met.
type ChatDisturbRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatDisturbRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatDisturbRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatDisturbRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatDisturbRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatDisturbRequestValidationError) ErrorName() string {
	return "ChatDisturbRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatDisturbRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatDisturbRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatDisturbRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatDisturbRequestValidationError{}

// Validate checks the field values on ChatDisturbResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatDisturbResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatDisturbResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatDisturbResponseMultiError, or nil if none found.
func (m *ChatDisturbResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatDisturbResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatDisturbResponseMultiError(errors)
	}

	return nil
}

// ChatDisturbResponseMultiError is an error wrapping multiple validation
// errors returned by ChatDisturbResponse.ValidateAll() if the designated
// constraints aren't met.
type ChatDisturbResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatDisturbResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatDisturbResponseMultiError) AllErrors() []error { return m }

// ChatDisturbResponseValidationError is the validation error returned by
// ChatDisturbResponse.Validate if the designated constraints aren't met.
type ChatDisturbResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatDisturbResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatDisturbResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatDisturbResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatDisturbResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatDisturbResponseValidationError) ErrorName() string {
	return "ChatDisturbResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatDisturbResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatDisturbResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatDisturbResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatDisturbResponseValidationError{}

// Validate checks the field values on ChatClearUnreadNumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatClearUnreadNumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatClearUnreadNumRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatClearUnreadNumRequestMultiError, or nil if none found.
func (m *ChatClearUnreadNumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatClearUnreadNumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DialogType

	// no validation rules for ReceiverId

	if len(errors) > 0 {
		return ChatClearUnreadNumRequestMultiError(errors)
	}

	return nil
}

// ChatClearUnreadNumRequestMultiError is an error wrapping multiple validation
// errors returned by ChatClearUnreadNumRequest.ValidateAll() if the
// designated constraints aren't met.
type ChatClearUnreadNumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatClearUnreadNumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatClearUnreadNumRequestMultiError) AllErrors() []error { return m }

// ChatClearUnreadNumRequestValidationError is the validation error returned by
// ChatClearUnreadNumRequest.Validate if the designated constraints aren't met.
type ChatClearUnreadNumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatClearUnreadNumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatClearUnreadNumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatClearUnreadNumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatClearUnreadNumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatClearUnreadNumRequestValidationError) ErrorName() string {
	return "ChatClearUnreadNumRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChatClearUnreadNumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatClearUnreadNumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatClearUnreadNumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatClearUnreadNumRequestValidationError{}

// Validate checks the field values on ChatClearUnreadNumResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ChatClearUnreadNumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatClearUnreadNumResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ChatClearUnreadNumResponseMultiError, or nil if none found.
func (m *ChatClearUnreadNumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatClearUnreadNumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ChatClearUnreadNumResponseMultiError(errors)
	}

	return nil
}

// ChatClearUnreadNumResponseMultiError is an error wrapping multiple
// validation errors returned by ChatClearUnreadNumResponse.ValidateAll() if
// the designated constraints aren't met.
type ChatClearUnreadNumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatClearUnreadNumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatClearUnreadNumResponseMultiError) AllErrors() []error { return m }

// ChatClearUnreadNumResponseValidationError is the validation error returned
// by ChatClearUnreadNumResponse.Validate if the designated constraints aren't met.
type ChatClearUnreadNumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatClearUnreadNumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatClearUnreadNumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatClearUnreadNumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatClearUnreadNumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatClearUnreadNumResponseValidationError) ErrorName() string {
	return "ChatClearUnreadNumResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ChatClearUnreadNumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatClearUnreadNumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatClearUnreadNumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatClearUnreadNumResponseValidationError{}

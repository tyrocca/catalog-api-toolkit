// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: catalog/v1/company.proto

package v1

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

// Validate checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Company) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Company with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CompanyMultiError, or nil if none found.
func (m *Company) ValidateAll() error {
	return m.validate(true)
}

func (m *Company) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Uid

	if l := len(m.GetDisplayName()); l < 1 || l > 256 {
		err := CompanyValidationError{
			field:  "DisplayName",
			reason: "value length must be between 1 and 256 bytes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompanyValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompanyValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompanyValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompanyValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompanyValidationError{
					field:  "UpdateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompanyValidationError{
				field:  "UpdateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CompanyMultiError(errors)
	}

	return nil
}

// CompanyMultiError is an error wrapping multiple validation errors returned
// by Company.ValidateAll() if the designated constraints aren't met.
type CompanyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompanyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompanyMultiError) AllErrors() []error { return m }

// CompanyValidationError is the validation error returned by Company.Validate
// if the designated constraints aren't met.
type CompanyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompanyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompanyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompanyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompanyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompanyValidationError) ErrorName() string { return "CompanyValidationError" }

// Error satisfies the builtin error interface
func (e CompanyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCompany.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompanyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompanyValidationError{}

// Validate checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCompanyRequestMultiError, or nil if none found.
func (m *CreateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCompany()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompany()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateCompanyRequestValidationError{
				field:  "Company",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateCompanyRequestMultiError(errors)
	}

	return nil
}

// CreateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCompanyRequestMultiError) AllErrors() []error { return m }

// CreateCompanyRequestValidationError is the validation error returned by
// CreateCompanyRequest.Validate if the designated constraints aren't met.
type CreateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCompanyRequestValidationError) ErrorName() string {
	return "CreateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCompanyRequestValidationError{}

// Validate checks the field values on GetCompanyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompanyRequestMultiError, or nil if none found.
func (m *GetCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetCompanyRequestMultiError(errors)
	}

	return nil
}

// GetCompanyRequestMultiError is an error wrapping multiple validation errors
// returned by GetCompanyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompanyRequestMultiError) AllErrors() []error { return m }

// GetCompanyRequestValidationError is the validation error returned by
// GetCompanyRequest.Validate if the designated constraints aren't met.
type GetCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompanyRequestValidationError) ErrorName() string {
	return "GetCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompanyRequestValidationError{}

// Validate checks the field values on ListCompaniesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCompaniesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCompaniesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCompaniesRequestMultiError, or nil if none found.
func (m *ListCompaniesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCompaniesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageToken

	// no validation rules for Filter

	// no validation rules for OrderBy

	if len(errors) > 0 {
		return ListCompaniesRequestMultiError(errors)
	}

	return nil
}

// ListCompaniesRequestMultiError is an error wrapping multiple validation
// errors returned by ListCompaniesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListCompaniesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCompaniesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCompaniesRequestMultiError) AllErrors() []error { return m }

// ListCompaniesRequestValidationError is the validation error returned by
// ListCompaniesRequest.Validate if the designated constraints aren't met.
type ListCompaniesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCompaniesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCompaniesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCompaniesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCompaniesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCompaniesRequestValidationError) ErrorName() string {
	return "ListCompaniesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCompaniesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCompaniesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCompaniesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCompaniesRequestValidationError{}

// Validate checks the field values on ListCompaniesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCompaniesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCompaniesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCompaniesResponseMultiError, or nil if none found.
func (m *ListCompaniesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCompaniesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCompanies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCompaniesResponseValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCompaniesResponseValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCompaniesResponseValidationError{
					field:  fmt.Sprintf("Companies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListCompaniesResponseMultiError(errors)
	}

	return nil
}

// ListCompaniesResponseMultiError is an error wrapping multiple validation
// errors returned by ListCompaniesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListCompaniesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCompaniesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCompaniesResponseMultiError) AllErrors() []error { return m }

// ListCompaniesResponseValidationError is the validation error returned by
// ListCompaniesResponse.Validate if the designated constraints aren't met.
type ListCompaniesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCompaniesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCompaniesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCompaniesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCompaniesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCompaniesResponseValidationError) ErrorName() string {
	return "ListCompaniesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListCompaniesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCompaniesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCompaniesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCompaniesResponseValidationError{}

// Validate checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCompanyRequestMultiError, or nil if none found.
func (m *UpdateCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCompany()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCompanyRequestValidationError{
					field:  "Company",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompany()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCompanyRequestValidationError{
				field:  "Company",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCompanyRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCompanyRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCompanyRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateCompanyRequestMultiError(errors)
	}

	return nil
}

// UpdateCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCompanyRequestMultiError) AllErrors() []error { return m }

// UpdateCompanyRequestValidationError is the validation error returned by
// UpdateCompanyRequest.Validate if the designated constraints aren't met.
type UpdateCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCompanyRequestValidationError) ErrorName() string {
	return "UpdateCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCompanyRequestValidationError{}

// Validate checks the field values on DeleteCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCompanyRequestMultiError, or nil if none found.
func (m *DeleteCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Force

	if len(errors) > 0 {
		return DeleteCompanyRequestMultiError(errors)
	}

	return nil
}

// DeleteCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCompanyRequestMultiError) AllErrors() []error { return m }

// DeleteCompanyRequestValidationError is the validation error returned by
// DeleteCompanyRequest.Validate if the designated constraints aren't met.
type DeleteCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCompanyRequestValidationError) ErrorName() string {
	return "DeleteCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCompanyRequestValidationError{}
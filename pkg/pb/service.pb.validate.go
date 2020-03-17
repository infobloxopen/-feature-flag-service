// Code generated by protoc-gen-validate
// source: github.com/Infoblox-CTO/atlas.feature.flag/pkg/pb/service.proto
// DO NOT EDIT!!!

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on VersionResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *VersionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Version

	return nil
}

// VersionResponseValidationError is the validation error returned by
// VersionResponse.Validate if the designated constraints aren't met.
type VersionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VersionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VersionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VersionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VersionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VersionResponseValidationError) ErrorName() string { return "VersionResponseValidationError" }

// Error satisfies the builtin error interface
func (e VersionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVersionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VersionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VersionResponseValidationError{}

// Validate checks the field values on FeatureFlag with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FeatureFlag) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FeatureName

	// no validation rules for Value

	// no validation rules for Origin

	return nil
}

// FeatureFlagValidationError is the validation error returned by
// FeatureFlag.Validate if the designated constraints aren't met.
type FeatureFlagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeatureFlagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeatureFlagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeatureFlagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeatureFlagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeatureFlagValidationError) ErrorName() string { return "FeatureFlagValidationError" }

// Error satisfies the builtin error interface
func (e FeatureFlagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatureFlag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeatureFlagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeatureFlagValidationError{}

// Validate checks the field values on ListFeatureFlagsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeatureFlagsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeatureFlagsRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Labels

	return nil
}

// ListFeatureFlagsRequestValidationError is the validation error returned by
// ListFeatureFlagsRequest.Validate if the designated constraints aren't met.
type ListFeatureFlagsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeatureFlagsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeatureFlagsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeatureFlagsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeatureFlagsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeatureFlagsRequestValidationError) ErrorName() string {
	return "ListFeatureFlagsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeatureFlagsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeatureFlagsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeatureFlagsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeatureFlagsRequestValidationError{}

// Validate checks the field values on ListFeatureFlagsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeatureFlagsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFeatureFlagsResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListFeatureFlagsResponseValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListFeatureFlagsResponseValidationError is the validation error returned by
// ListFeatureFlagsResponse.Validate if the designated constraints aren't met.
type ListFeatureFlagsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeatureFlagsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeatureFlagsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeatureFlagsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeatureFlagsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeatureFlagsResponseValidationError) ErrorName() string {
	return "ListFeatureFlagsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeatureFlagsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeatureFlagsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeatureFlagsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeatureFlagsResponseValidationError{}

// Validate checks the field values on ReadFeatureFlagRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadFeatureFlagRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FeatureName

	if v, ok := interface{}(m.GetFields()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadFeatureFlagRequestValidationError{
				field:  "Fields",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Labels

	return nil
}

// ReadFeatureFlagRequestValidationError is the validation error returned by
// ReadFeatureFlagRequest.Validate if the designated constraints aren't met.
type ReadFeatureFlagRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadFeatureFlagRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadFeatureFlagRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadFeatureFlagRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadFeatureFlagRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadFeatureFlagRequestValidationError) ErrorName() string {
	return "ReadFeatureFlagRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadFeatureFlagRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadFeatureFlagRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadFeatureFlagRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadFeatureFlagRequestValidationError{}

// Validate checks the field values on ReadFeatureFlagResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReadFeatureFlagResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadFeatureFlagResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ReadFeatureFlagResponseValidationError is the validation error returned by
// ReadFeatureFlagResponse.Validate if the designated constraints aren't met.
type ReadFeatureFlagResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadFeatureFlagResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadFeatureFlagResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadFeatureFlagResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadFeatureFlagResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadFeatureFlagResponseValidationError) ErrorName() string {
	return "ReadFeatureFlagResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadFeatureFlagResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadFeatureFlagResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadFeatureFlagResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadFeatureFlagResponseValidationError{}

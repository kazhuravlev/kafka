// Code generated by options-gen. DO NOT EDIT.
package producer

import (
	fmt461e464ebed9 "fmt"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter func(o *Options)

func NewOptions(
	addrs []string,
	options ...OptOptionsSetter,
) Options {
	o := Options{}

	// Setting defaults from variable
	o.addrs = defaultOptions.addrs

	o.topic = defaultOptions.topic

	o.tls = defaultOptions.tls

	o.clientID = defaultOptions.clientID

	o.addrs = addrs

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithTopic(opt string) OptOptionsSetter {
	return func(o *Options) {
		o.topic = opt

	}
}

func WithTls(opt bool) OptOptionsSetter {
	return func(o *Options) {
		o.tls = opt

	}
}

func WithClientID(opt string) OptOptionsSetter {
	return func(o *Options) {
		o.clientID = opt

	}
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("addrs", _validate_Options_addrs(o)))
	errs.Add(errors461e464ebed9.NewValidationError("clientID", _validate_Options_clientID(o)))
	return errs.AsError()
}

func _validate_Options_addrs(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.addrs, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `addrs` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_clientID(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.clientID, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `clientID` did not pass the test: %w", err)
	}
	return nil
}
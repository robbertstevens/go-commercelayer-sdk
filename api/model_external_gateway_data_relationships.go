/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalGatewayDataRelationships struct for ExternalGatewayDataRelationships
type ExternalGatewayDataRelationships struct {
	PaymentMethods   *AdyenGatewayDataRelationshipsPaymentMethods      `json:"payment_methods,omitempty"`
	ExternalPayments *ExternalGatewayDataRelationshipsExternalPayments `json:"external_payments,omitempty"`
}

// NewExternalGatewayDataRelationships instantiates a new ExternalGatewayDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGatewayDataRelationships() *ExternalGatewayDataRelationships {
	this := ExternalGatewayDataRelationships{}
	return &this
}

// NewExternalGatewayDataRelationshipsWithDefaults instantiates a new ExternalGatewayDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGatewayDataRelationshipsWithDefaults() *ExternalGatewayDataRelationships {
	this := ExternalGatewayDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *ExternalGatewayDataRelationships) GetPaymentMethods() AdyenGatewayDataRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret AdyenGatewayDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataRelationships) GetPaymentMethodsOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *ExternalGatewayDataRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given AdyenGatewayDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *ExternalGatewayDataRelationships) SetPaymentMethods(v AdyenGatewayDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetExternalPayments returns the ExternalPayments field value if set, zero value otherwise.
func (o *ExternalGatewayDataRelationships) GetExternalPayments() ExternalGatewayDataRelationshipsExternalPayments {
	if o == nil || o.ExternalPayments == nil {
		var ret ExternalGatewayDataRelationshipsExternalPayments
		return ret
	}
	return *o.ExternalPayments
}

// GetExternalPaymentsOk returns a tuple with the ExternalPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGatewayDataRelationships) GetExternalPaymentsOk() (*ExternalGatewayDataRelationshipsExternalPayments, bool) {
	if o == nil || o.ExternalPayments == nil {
		return nil, false
	}
	return o.ExternalPayments, true
}

// HasExternalPayments returns a boolean if a field has been set.
func (o *ExternalGatewayDataRelationships) HasExternalPayments() bool {
	if o != nil && o.ExternalPayments != nil {
		return true
	}

	return false
}

// SetExternalPayments gets a reference to the given ExternalGatewayDataRelationshipsExternalPayments and assigns it to the ExternalPayments field.
func (o *ExternalGatewayDataRelationships) SetExternalPayments(v ExternalGatewayDataRelationshipsExternalPayments) {
	o.ExternalPayments = &v
}

func (o ExternalGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.ExternalPayments != nil {
		toSerialize["external_payments"] = o.ExternalPayments
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGatewayDataRelationships struct {
	value *ExternalGatewayDataRelationships
	isSet bool
}

func (v NullableExternalGatewayDataRelationships) Get() *ExternalGatewayDataRelationships {
	return v.value
}

func (v *NullableExternalGatewayDataRelationships) Set(val *ExternalGatewayDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGatewayDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGatewayDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGatewayDataRelationships(val *ExternalGatewayDataRelationships) *NullableExternalGatewayDataRelationships {
	return &NullableExternalGatewayDataRelationships{value: val, isSet: true}
}

func (v NullableExternalGatewayDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGatewayDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

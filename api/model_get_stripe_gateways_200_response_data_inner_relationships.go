/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETStripeGateways200ResponseDataInnerRelationships struct for GETStripeGateways200ResponseDataInnerRelationships
type GETStripeGateways200ResponseDataInnerRelationships struct {
	PaymentMethods *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods  `json:"payment_methods,omitempty"`
	StripePayments *GETStripeGateways200ResponseDataInnerRelationshipsStripePayments `json:"stripe_payments,omitempty"`
}

// NewGETStripeGateways200ResponseDataInnerRelationships instantiates a new GETStripeGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStripeGateways200ResponseDataInnerRelationships() *GETStripeGateways200ResponseDataInnerRelationships {
	this := GETStripeGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETStripeGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStripeGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStripeGateways200ResponseDataInnerRelationshipsWithDefaults() *GETStripeGateways200ResponseDataInnerRelationships {
	this := GETStripeGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETStripeGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetStripePayments returns the StripePayments field value if set, zero value otherwise.
func (o *GETStripeGateways200ResponseDataInnerRelationships) GetStripePayments() GETStripeGateways200ResponseDataInnerRelationshipsStripePayments {
	if o == nil || o.StripePayments == nil {
		var ret GETStripeGateways200ResponseDataInnerRelationshipsStripePayments
		return ret
	}
	return *o.StripePayments
}

// GetStripePaymentsOk returns a tuple with the StripePayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStripeGateways200ResponseDataInnerRelationships) GetStripePaymentsOk() (*GETStripeGateways200ResponseDataInnerRelationshipsStripePayments, bool) {
	if o == nil || o.StripePayments == nil {
		return nil, false
	}
	return o.StripePayments, true
}

// HasStripePayments returns a boolean if a field has been set.
func (o *GETStripeGateways200ResponseDataInnerRelationships) HasStripePayments() bool {
	if o != nil && o.StripePayments != nil {
		return true
	}

	return false
}

// SetStripePayments gets a reference to the given GETStripeGateways200ResponseDataInnerRelationshipsStripePayments and assigns it to the StripePayments field.
func (o *GETStripeGateways200ResponseDataInnerRelationships) SetStripePayments(v GETStripeGateways200ResponseDataInnerRelationshipsStripePayments) {
	o.StripePayments = &v
}

func (o GETStripeGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.StripePayments != nil {
		toSerialize["stripe_payments"] = o.StripePayments
	}
	return json.Marshal(toSerialize)
}

type NullableGETStripeGateways200ResponseDataInnerRelationships struct {
	value *GETStripeGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETStripeGateways200ResponseDataInnerRelationships) Get() *GETStripeGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETStripeGateways200ResponseDataInnerRelationships) Set(val *GETStripeGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStripeGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStripeGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStripeGateways200ResponseDataInnerRelationships(val *GETStripeGateways200ResponseDataInnerRelationships) *NullableGETStripeGateways200ResponseDataInnerRelationships {
	return &NullableGETStripeGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETStripeGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStripeGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

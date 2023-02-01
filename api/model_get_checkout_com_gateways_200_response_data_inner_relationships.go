/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCheckoutComGateways200ResponseDataInnerRelationships struct for GETCheckoutComGateways200ResponseDataInnerRelationships
type GETCheckoutComGateways200ResponseDataInnerRelationships struct {
	PaymentMethods      *GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods            `json:"payment_methods,omitempty"`
	CheckoutComPayments *GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments `json:"checkout_com_payments,omitempty"`
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationships instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCheckoutComGateways200ResponseDataInnerRelationships() *GETCheckoutComGateways200ResponseDataInnerRelationships {
	this := GETCheckoutComGateways200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCheckoutComGateways200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCheckoutComGateways200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCheckoutComGateways200ResponseDataInnerRelationshipsWithDefaults() *GETCheckoutComGateways200ResponseDataInnerRelationships {
	this := GETCheckoutComGateways200ResponseDataInnerRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) GetPaymentMethods() GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods {
	if o == nil || o.PaymentMethods == nil {
		var ret GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) GetPaymentMethodsOk() (*GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods, bool) {
	if o == nil || o.PaymentMethods == nil {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) HasPaymentMethods() bool {
	if o != nil && o.PaymentMethods != nil {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) SetPaymentMethods(v GETAdyenGateways200ResponseDataInnerRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetCheckoutComPayments returns the CheckoutComPayments field value if set, zero value otherwise.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) GetCheckoutComPayments() GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments {
	if o == nil || o.CheckoutComPayments == nil {
		var ret GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments
		return ret
	}
	return *o.CheckoutComPayments
}

// GetCheckoutComPaymentsOk returns a tuple with the CheckoutComPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) GetCheckoutComPaymentsOk() (*GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments, bool) {
	if o == nil || o.CheckoutComPayments == nil {
		return nil, false
	}
	return o.CheckoutComPayments, true
}

// HasCheckoutComPayments returns a boolean if a field has been set.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) HasCheckoutComPayments() bool {
	if o != nil && o.CheckoutComPayments != nil {
		return true
	}

	return false
}

// SetCheckoutComPayments gets a reference to the given GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments and assigns it to the CheckoutComPayments field.
func (o *GETCheckoutComGateways200ResponseDataInnerRelationships) SetCheckoutComPayments(v GETCheckoutComGateways200ResponseDataInnerRelationshipsCheckoutComPayments) {
	o.CheckoutComPayments = &v
}

func (o GETCheckoutComGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentMethods != nil {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if o.CheckoutComPayments != nil {
		toSerialize["checkout_com_payments"] = o.CheckoutComPayments
	}
	return json.Marshal(toSerialize)
}

type NullableGETCheckoutComGateways200ResponseDataInnerRelationships struct {
	value *GETCheckoutComGateways200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationships) Get() *GETCheckoutComGateways200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationships) Set(val *GETCheckoutComGateways200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCheckoutComGateways200ResponseDataInnerRelationships(val *GETCheckoutComGateways200ResponseDataInnerRelationships) *NullableGETCheckoutComGateways200ResponseDataInnerRelationships {
	return &NullableGETCheckoutComGateways200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCheckoutComGateways200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCheckoutComGateways200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTSatispayGateways201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSatispayGateways201ResponseDataRelationships{}

// POSTSatispayGateways201ResponseDataRelationships struct for POSTSatispayGateways201ResponseDataRelationships
type POSTSatispayGateways201ResponseDataRelationships struct {
	PaymentMethods   *POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods      `json:"payment_methods,omitempty"`
	SatispayPayments *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments `json:"satispay_payments,omitempty"`
}

// NewPOSTSatispayGateways201ResponseDataRelationships instantiates a new POSTSatispayGateways201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSatispayGateways201ResponseDataRelationships() *POSTSatispayGateways201ResponseDataRelationships {
	this := POSTSatispayGateways201ResponseDataRelationships{}
	return &this
}

// NewPOSTSatispayGateways201ResponseDataRelationshipsWithDefaults instantiates a new POSTSatispayGateways201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSatispayGateways201ResponseDataRelationshipsWithDefaults() *POSTSatispayGateways201ResponseDataRelationships {
	this := POSTSatispayGateways201ResponseDataRelationships{}
	return &this
}

// GetPaymentMethods returns the PaymentMethods field value if set, zero value otherwise.
func (o *POSTSatispayGateways201ResponseDataRelationships) GetPaymentMethods() POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods {
	if o == nil || IsNil(o.PaymentMethods) {
		var ret POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods
		return ret
	}
	return *o.PaymentMethods
}

// GetPaymentMethodsOk returns a tuple with the PaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayGateways201ResponseDataRelationships) GetPaymentMethodsOk() (*POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods, bool) {
	if o == nil || IsNil(o.PaymentMethods) {
		return nil, false
	}
	return o.PaymentMethods, true
}

// HasPaymentMethods returns a boolean if a field has been set.
func (o *POSTSatispayGateways201ResponseDataRelationships) HasPaymentMethods() bool {
	if o != nil && !IsNil(o.PaymentMethods) {
		return true
	}

	return false
}

// SetPaymentMethods gets a reference to the given POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods and assigns it to the PaymentMethods field.
func (o *POSTSatispayGateways201ResponseDataRelationships) SetPaymentMethods(v POSTAdyenGateways201ResponseDataRelationshipsPaymentMethods) {
	o.PaymentMethods = &v
}

// GetSatispayPayments returns the SatispayPayments field value if set, zero value otherwise.
func (o *POSTSatispayGateways201ResponseDataRelationships) GetSatispayPayments() POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments {
	if o == nil || IsNil(o.SatispayPayments) {
		var ret POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments
		return ret
	}
	return *o.SatispayPayments
}

// GetSatispayPaymentsOk returns a tuple with the SatispayPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayGateways201ResponseDataRelationships) GetSatispayPaymentsOk() (*POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments, bool) {
	if o == nil || IsNil(o.SatispayPayments) {
		return nil, false
	}
	return o.SatispayPayments, true
}

// HasSatispayPayments returns a boolean if a field has been set.
func (o *POSTSatispayGateways201ResponseDataRelationships) HasSatispayPayments() bool {
	if o != nil && !IsNil(o.SatispayPayments) {
		return true
	}

	return false
}

// SetSatispayPayments gets a reference to the given POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments and assigns it to the SatispayPayments field.
func (o *POSTSatispayGateways201ResponseDataRelationships) SetSatispayPayments(v POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) {
	o.SatispayPayments = &v
}

func (o POSTSatispayGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSatispayGateways201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethods) {
		toSerialize["payment_methods"] = o.PaymentMethods
	}
	if !IsNil(o.SatispayPayments) {
		toSerialize["satispay_payments"] = o.SatispayPayments
	}
	return toSerialize, nil
}

type NullablePOSTSatispayGateways201ResponseDataRelationships struct {
	value *POSTSatispayGateways201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationships) Get() *POSTSatispayGateways201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationships) Set(val *POSTSatispayGateways201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSatispayGateways201ResponseDataRelationships(val *POSTSatispayGateways201ResponseDataRelationships) *NullablePOSTSatispayGateways201ResponseDataRelationships {
	return &NullablePOSTSatispayGateways201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

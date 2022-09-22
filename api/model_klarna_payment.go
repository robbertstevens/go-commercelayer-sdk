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

// KlarnaPayment struct for KlarnaPayment
type KlarnaPayment struct {
	Data KlarnaPaymentData `json:"data"`
}

// NewKlarnaPayment instantiates a new KlarnaPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaPayment(data KlarnaPaymentData) *KlarnaPayment {
	this := KlarnaPayment{}
	this.Data = data
	return &this
}

// NewKlarnaPaymentWithDefaults instantiates a new KlarnaPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaPaymentWithDefaults() *KlarnaPayment {
	this := KlarnaPayment{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaPayment) GetData() KlarnaPaymentData {
	if o == nil {
		var ret KlarnaPaymentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaPayment) GetDataOk() (*KlarnaPaymentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaPayment) SetData(v KlarnaPaymentData) {
	o.Data = v
}

func (o KlarnaPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaPayment struct {
	value *KlarnaPayment
	isSet bool
}

func (v NullableKlarnaPayment) Get() *KlarnaPayment {
	return v.value
}

func (v *NullableKlarnaPayment) Set(val *KlarnaPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaPayment(val *KlarnaPayment) *NullableKlarnaPayment {
	return &NullableKlarnaPayment{value: val, isSet: true}
}

func (v NullableKlarnaPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

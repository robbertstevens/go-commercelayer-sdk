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

// checks if the KlarnaPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KlarnaPayment{}

// KlarnaPayment struct for KlarnaPayment
type KlarnaPayment struct {
	Data *KlarnaPaymentData `json:"data,omitempty"`
}

// NewKlarnaPayment instantiates a new KlarnaPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaPayment() *KlarnaPayment {
	this := KlarnaPayment{}
	return &this
}

// NewKlarnaPaymentWithDefaults instantiates a new KlarnaPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaPaymentWithDefaults() *KlarnaPayment {
	this := KlarnaPayment{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KlarnaPayment) GetData() KlarnaPaymentData {
	if o == nil || IsNil(o.Data) {
		var ret KlarnaPaymentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlarnaPayment) GetDataOk() (*KlarnaPaymentData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KlarnaPayment) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given KlarnaPaymentData and assigns it to the Data field.
func (o *KlarnaPayment) SetData(v KlarnaPaymentData) {
	o.Data = &v
}

func (o KlarnaPayment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlarnaPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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

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

// StripePaymentUpdate struct for StripePaymentUpdate
type StripePaymentUpdate struct {
	Data StripePaymentUpdateData `json:"data"`
}

// NewStripePaymentUpdate instantiates a new StripePaymentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripePaymentUpdate(data StripePaymentUpdateData) *StripePaymentUpdate {
	this := StripePaymentUpdate{}
	this.Data = data
	return &this
}

// NewStripePaymentUpdateWithDefaults instantiates a new StripePaymentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripePaymentUpdateWithDefaults() *StripePaymentUpdate {
	this := StripePaymentUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StripePaymentUpdate) GetData() StripePaymentUpdateData {
	if o == nil {
		var ret StripePaymentUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StripePaymentUpdate) GetDataOk() (*StripePaymentUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StripePaymentUpdate) SetData(v StripePaymentUpdateData) {
	o.Data = v
}

func (o StripePaymentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripePaymentUpdate struct {
	value *StripePaymentUpdate
	isSet bool
}

func (v NullableStripePaymentUpdate) Get() *StripePaymentUpdate {
	return v.value
}

func (v *NullableStripePaymentUpdate) Set(val *StripePaymentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStripePaymentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStripePaymentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripePaymentUpdate(val *StripePaymentUpdate) *NullableStripePaymentUpdate {
	return &NullableStripePaymentUpdate{value: val, isSet: true}
}

func (v NullableStripePaymentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripePaymentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

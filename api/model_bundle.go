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

// Bundle struct for Bundle
type Bundle struct {
	Data BundleData `json:"data"`
}

// NewBundle instantiates a new Bundle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundle(data BundleData) *Bundle {
	this := Bundle{}
	this.Data = data
	return &this
}

// NewBundleWithDefaults instantiates a new Bundle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleWithDefaults() *Bundle {
	this := Bundle{}
	return &this
}

// GetData returns the Data field value
func (o *Bundle) GetData() BundleData {
	if o == nil {
		var ret BundleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Bundle) GetDataOk() (*BundleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Bundle) SetData(v BundleData) {
	o.Data = v
}

func (o Bundle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundle struct {
	value *Bundle
	isSet bool
}

func (v NullableBundle) Get() *Bundle {
	return v.value
}

func (v *NullableBundle) Set(val *Bundle) {
	v.value = val
	v.isSet = true
}

func (v NullableBundle) IsSet() bool {
	return v.isSet
}

func (v *NullableBundle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundle(val *Bundle) *NullableBundle {
	return &NullableBundle{value: val, isSet: true}
}

func (v NullableBundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

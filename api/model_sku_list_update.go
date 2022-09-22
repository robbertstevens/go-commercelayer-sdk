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

// SkuListUpdate struct for SkuListUpdate
type SkuListUpdate struct {
	Data SkuListUpdateData `json:"data"`
}

// NewSkuListUpdate instantiates a new SkuListUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListUpdate(data SkuListUpdateData) *SkuListUpdate {
	this := SkuListUpdate{}
	this.Data = data
	return &this
}

// NewSkuListUpdateWithDefaults instantiates a new SkuListUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListUpdateWithDefaults() *SkuListUpdate {
	this := SkuListUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuListUpdate) GetData() SkuListUpdateData {
	if o == nil {
		var ret SkuListUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuListUpdate) GetDataOk() (*SkuListUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuListUpdate) SetData(v SkuListUpdateData) {
	o.Data = v
}

func (o SkuListUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListUpdate struct {
	value *SkuListUpdate
	isSet bool
}

func (v NullableSkuListUpdate) Get() *SkuListUpdate {
	return v.value
}

func (v *NullableSkuListUpdate) Set(val *SkuListUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListUpdate(val *SkuListUpdate) *NullableSkuListUpdate {
	return &NullableSkuListUpdate{value: val, isSet: true}
}

func (v NullableSkuListUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

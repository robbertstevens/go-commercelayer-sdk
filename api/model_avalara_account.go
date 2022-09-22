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

// AvalaraAccount struct for AvalaraAccount
type AvalaraAccount struct {
	Data AvalaraAccountData `json:"data"`
}

// NewAvalaraAccount instantiates a new AvalaraAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccount(data AvalaraAccountData) *AvalaraAccount {
	this := AvalaraAccount{}
	this.Data = data
	return &this
}

// NewAvalaraAccountWithDefaults instantiates a new AvalaraAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountWithDefaults() *AvalaraAccount {
	this := AvalaraAccount{}
	return &this
}

// GetData returns the Data field value
func (o *AvalaraAccount) GetData() AvalaraAccountData {
	if o == nil {
		var ret AvalaraAccountData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccount) GetDataOk() (*AvalaraAccountData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvalaraAccount) SetData(v AvalaraAccountData) {
	o.Data = v
}

func (o AvalaraAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccount struct {
	value *AvalaraAccount
	isSet bool
}

func (v NullableAvalaraAccount) Get() *AvalaraAccount {
	return v.value
}

func (v *NullableAvalaraAccount) Set(val *AvalaraAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccount(val *AvalaraAccount) *NullableAvalaraAccount {
	return &NullableAvalaraAccount{value: val, isSet: true}
}

func (v NullableAvalaraAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

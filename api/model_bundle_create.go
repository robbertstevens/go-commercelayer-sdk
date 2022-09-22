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

// BundleCreate struct for BundleCreate
type BundleCreate struct {
	Data BundleCreateData `json:"data"`
}

// NewBundleCreate instantiates a new BundleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreate(data BundleCreateData) *BundleCreate {
	this := BundleCreate{}
	this.Data = data
	return &this
}

// NewBundleCreateWithDefaults instantiates a new BundleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateWithDefaults() *BundleCreate {
	this := BundleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *BundleCreate) GetData() BundleCreateData {
	if o == nil {
		var ret BundleCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleCreate) GetDataOk() (*BundleCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleCreate) SetData(v BundleCreateData) {
	o.Data = v
}

func (o BundleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundleCreate struct {
	value *BundleCreate
	isSet bool
}

func (v NullableBundleCreate) Get() *BundleCreate {
	return v.value
}

func (v *NullableBundleCreate) Set(val *BundleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreate(val *BundleCreate) *NullableBundleCreate {
	return &NullableBundleCreate{value: val, isSet: true}
}

func (v NullableBundleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

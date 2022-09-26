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

// BundleDataRelationshipsSkus struct for BundleDataRelationshipsSkus
type BundleDataRelationshipsSkus struct {
	Data BundleDataRelationshipsSkusData `json:"data"`
}

// NewBundleDataRelationshipsSkus instantiates a new BundleDataRelationshipsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleDataRelationshipsSkus(data BundleDataRelationshipsSkusData) *BundleDataRelationshipsSkus {
	this := BundleDataRelationshipsSkus{}
	this.Data = data
	return &this
}

// NewBundleDataRelationshipsSkusWithDefaults instantiates a new BundleDataRelationshipsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleDataRelationshipsSkusWithDefaults() *BundleDataRelationshipsSkus {
	this := BundleDataRelationshipsSkus{}
	return &this
}

// GetData returns the Data field value
func (o *BundleDataRelationshipsSkus) GetData() BundleDataRelationshipsSkusData {
	if o == nil {
		var ret BundleDataRelationshipsSkusData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleDataRelationshipsSkus) GetDataOk() (*BundleDataRelationshipsSkusData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleDataRelationshipsSkus) SetData(v BundleDataRelationshipsSkusData) {
	o.Data = v
}

func (o BundleDataRelationshipsSkus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBundleDataRelationshipsSkus struct {
	value *BundleDataRelationshipsSkus
	isSet bool
}

func (v NullableBundleDataRelationshipsSkus) Get() *BundleDataRelationshipsSkus {
	return v.value
}

func (v *NullableBundleDataRelationshipsSkus) Set(val *BundleDataRelationshipsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleDataRelationshipsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleDataRelationshipsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleDataRelationshipsSkus(val *BundleDataRelationshipsSkus) *NullableBundleDataRelationshipsSkus {
	return &NullableBundleDataRelationshipsSkus{value: val, isSet: true}
}

func (v NullableBundleDataRelationshipsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleDataRelationshipsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

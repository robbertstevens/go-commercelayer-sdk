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

// InventoryModelUpdate struct for InventoryModelUpdate
type InventoryModelUpdate struct {
	Data InventoryModelUpdateData `json:"data"`
}

// NewInventoryModelUpdate instantiates a new InventoryModelUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelUpdate(data InventoryModelUpdateData) *InventoryModelUpdate {
	this := InventoryModelUpdate{}
	this.Data = data
	return &this
}

// NewInventoryModelUpdateWithDefaults instantiates a new InventoryModelUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelUpdateWithDefaults() *InventoryModelUpdate {
	this := InventoryModelUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryModelUpdate) GetData() InventoryModelUpdateData {
	if o == nil {
		var ret InventoryModelUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryModelUpdate) GetDataOk() (*InventoryModelUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryModelUpdate) SetData(v InventoryModelUpdateData) {
	o.Data = v
}

func (o InventoryModelUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelUpdate struct {
	value *InventoryModelUpdate
	isSet bool
}

func (v NullableInventoryModelUpdate) Get() *InventoryModelUpdate {
	return v.value
}

func (v *NullableInventoryModelUpdate) Set(val *InventoryModelUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelUpdate(val *InventoryModelUpdate) *NullableInventoryModelUpdate {
	return &NullableInventoryModelUpdate{value: val, isSet: true}
}

func (v NullableInventoryModelUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

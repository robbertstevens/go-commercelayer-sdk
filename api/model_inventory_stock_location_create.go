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

// InventoryStockLocationCreate struct for InventoryStockLocationCreate
type InventoryStockLocationCreate struct {
	Data InventoryStockLocationCreateData `json:"data"`
}

// NewInventoryStockLocationCreate instantiates a new InventoryStockLocationCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocationCreate(data InventoryStockLocationCreateData) *InventoryStockLocationCreate {
	this := InventoryStockLocationCreate{}
	this.Data = data
	return &this
}

// NewInventoryStockLocationCreateWithDefaults instantiates a new InventoryStockLocationCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationCreateWithDefaults() *InventoryStockLocationCreate {
	this := InventoryStockLocationCreate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryStockLocationCreate) GetData() InventoryStockLocationCreateData {
	if o == nil {
		var ret InventoryStockLocationCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationCreate) GetDataOk() (*InventoryStockLocationCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryStockLocationCreate) SetData(v InventoryStockLocationCreateData) {
	o.Data = v
}

func (o InventoryStockLocationCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryStockLocationCreate struct {
	value *InventoryStockLocationCreate
	isSet bool
}

func (v NullableInventoryStockLocationCreate) Get() *InventoryStockLocationCreate {
	return v.value
}

func (v *NullableInventoryStockLocationCreate) Set(val *InventoryStockLocationCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocationCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocationCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocationCreate(val *InventoryStockLocationCreate) *NullableInventoryStockLocationCreate {
	return &NullableInventoryStockLocationCreate{value: val, isSet: true}
}

func (v NullableInventoryStockLocationCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocationCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

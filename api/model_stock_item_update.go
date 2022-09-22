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

// StockItemUpdate struct for StockItemUpdate
type StockItemUpdate struct {
	Data StockItemUpdateData `json:"data"`
}

// NewStockItemUpdate instantiates a new StockItemUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemUpdate(data StockItemUpdateData) *StockItemUpdate {
	this := StockItemUpdate{}
	this.Data = data
	return &this
}

// NewStockItemUpdateWithDefaults instantiates a new StockItemUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemUpdateWithDefaults() *StockItemUpdate {
	this := StockItemUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StockItemUpdate) GetData() StockItemUpdateData {
	if o == nil {
		var ret StockItemUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockItemUpdate) GetDataOk() (*StockItemUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockItemUpdate) SetData(v StockItemUpdateData) {
	o.Data = v
}

func (o StockItemUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockItemUpdate struct {
	value *StockItemUpdate
	isSet bool
}

func (v NullableStockItemUpdate) Get() *StockItemUpdate {
	return v.value
}

func (v *NullableStockItemUpdate) Set(val *StockItemUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemUpdate(val *StockItemUpdate) *NullableStockItemUpdate {
	return &NullableStockItemUpdate{value: val, isSet: true}
}

func (v NullableStockItemUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

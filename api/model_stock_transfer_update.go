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

// StockTransferUpdate struct for StockTransferUpdate
type StockTransferUpdate struct {
	Data StockTransferUpdateData `json:"data"`
}

// NewStockTransferUpdate instantiates a new StockTransferUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferUpdate(data StockTransferUpdateData) *StockTransferUpdate {
	this := StockTransferUpdate{}
	this.Data = data
	return &this
}

// NewStockTransferUpdateWithDefaults instantiates a new StockTransferUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferUpdateWithDefaults() *StockTransferUpdate {
	this := StockTransferUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StockTransferUpdate) GetData() StockTransferUpdateData {
	if o == nil {
		var ret StockTransferUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockTransferUpdate) GetDataOk() (*StockTransferUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockTransferUpdate) SetData(v StockTransferUpdateData) {
	o.Data = v
}

func (o StockTransferUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferUpdate struct {
	value *StockTransferUpdate
	isSet bool
}

func (v NullableStockTransferUpdate) Get() *StockTransferUpdate {
	return v.value
}

func (v *NullableStockTransferUpdate) Set(val *StockTransferUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferUpdate(val *StockTransferUpdate) *NullableStockTransferUpdate {
	return &NullableStockTransferUpdate{value: val, isSet: true}
}

func (v NullableStockTransferUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the StockLineItemUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLineItemUpdate{}

// StockLineItemUpdate struct for StockLineItemUpdate
type StockLineItemUpdate struct {
	Data StockLineItemUpdateData `json:"data"`
}

// NewStockLineItemUpdate instantiates a new StockLineItemUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLineItemUpdate(data StockLineItemUpdateData) *StockLineItemUpdate {
	this := StockLineItemUpdate{}
	this.Data = data
	return &this
}

// NewStockLineItemUpdateWithDefaults instantiates a new StockLineItemUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLineItemUpdateWithDefaults() *StockLineItemUpdate {
	this := StockLineItemUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StockLineItemUpdate) GetData() StockLineItemUpdateData {
	if o == nil {
		var ret StockLineItemUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockLineItemUpdate) GetDataOk() (*StockLineItemUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockLineItemUpdate) SetData(v StockLineItemUpdateData) {
	o.Data = v
}

func (o StockLineItemUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLineItemUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStockLineItemUpdate struct {
	value *StockLineItemUpdate
	isSet bool
}

func (v NullableStockLineItemUpdate) Get() *StockLineItemUpdate {
	return v.value
}

func (v *NullableStockLineItemUpdate) Set(val *StockLineItemUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLineItemUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLineItemUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLineItemUpdate(val *StockLineItemUpdate) *NullableStockLineItemUpdate {
	return &NullableStockLineItemUpdate{value: val, isSet: true}
}

func (v NullableStockLineItemUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLineItemUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

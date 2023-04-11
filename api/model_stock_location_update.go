/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the StockLocationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StockLocationUpdate{}

// StockLocationUpdate struct for StockLocationUpdate
type StockLocationUpdate struct {
	Data PATCHStockLocationsStockLocationIdRequestData `json:"data"`
}

// NewStockLocationUpdate instantiates a new StockLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockLocationUpdate(data PATCHStockLocationsStockLocationIdRequestData) *StockLocationUpdate {
	this := StockLocationUpdate{}
	this.Data = data
	return &this
}

// NewStockLocationUpdateWithDefaults instantiates a new StockLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockLocationUpdateWithDefaults() *StockLocationUpdate {
	this := StockLocationUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *StockLocationUpdate) GetData() PATCHStockLocationsStockLocationIdRequestData {
	if o == nil {
		var ret PATCHStockLocationsStockLocationIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StockLocationUpdate) GetDataOk() (*PATCHStockLocationsStockLocationIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StockLocationUpdate) SetData(v PATCHStockLocationsStockLocationIdRequestData) {
	o.Data = v
}

func (o StockLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StockLocationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableStockLocationUpdate struct {
	value *StockLocationUpdate
	isSet bool
}

func (v NullableStockLocationUpdate) Get() *StockLocationUpdate {
	return v.value
}

func (v *NullableStockLocationUpdate) Set(val *StockLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStockLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStockLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockLocationUpdate(val *StockLocationUpdate) *NullableStockLocationUpdate {
	return &NullableStockLocationUpdate{value: val, isSet: true}
}

func (v NullableStockLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

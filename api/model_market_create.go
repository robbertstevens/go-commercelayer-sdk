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

// MarketCreate struct for MarketCreate
type MarketCreate struct {
	Data MarketCreateData `json:"data"`
}

// NewMarketCreate instantiates a new MarketCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCreate(data MarketCreateData) *MarketCreate {
	this := MarketCreate{}
	this.Data = data
	return &this
}

// NewMarketCreateWithDefaults instantiates a new MarketCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCreateWithDefaults() *MarketCreate {
	this := MarketCreate{}
	return &this
}

// GetData returns the Data field value
func (o *MarketCreate) GetData() MarketCreateData {
	if o == nil {
		var ret MarketCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MarketCreate) GetDataOk() (*MarketCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MarketCreate) SetData(v MarketCreateData) {
	o.Data = v
}

func (o MarketCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarketCreate struct {
	value *MarketCreate
	isSet bool
}

func (v NullableMarketCreate) Get() *MarketCreate {
	return v.value
}

func (v *NullableMarketCreate) Set(val *MarketCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCreate(val *MarketCreate) *NullableMarketCreate {
	return &NullableMarketCreate{value: val, isSet: true}
}

func (v NullableMarketCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

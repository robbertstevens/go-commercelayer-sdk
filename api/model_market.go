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

// Market struct for Market
type Market struct {
	Data MarketData `json:"data"`
}

// NewMarket instantiates a new Market object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarket(data MarketData) *Market {
	this := Market{}
	this.Data = data
	return &this
}

// NewMarketWithDefaults instantiates a new Market object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketWithDefaults() *Market {
	this := Market{}
	return &this
}

// GetData returns the Data field value
func (o *Market) GetData() MarketData {
	if o == nil {
		var ret MarketData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Market) GetDataOk() (*MarketData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Market) SetData(v MarketData) {
	o.Data = v
}

func (o Market) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMarket struct {
	value *Market
	isSet bool
}

func (v NullableMarket) Get() *Market {
	return v.value
}

func (v *NullableMarket) Set(val *Market) {
	v.value = val
	v.isSet = true
}

func (v NullableMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarket(val *Market) *NullableMarket {
	return &NullableMarket{value: val, isSet: true}
}

func (v NullableMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

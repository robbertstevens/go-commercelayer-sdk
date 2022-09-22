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

// PriceTier struct for PriceTier
type PriceTier struct {
	Data PriceTierData `json:"data"`
}

// NewPriceTier instantiates a new PriceTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTier(data PriceTierData) *PriceTier {
	this := PriceTier{}
	this.Data = data
	return &this
}

// NewPriceTierWithDefaults instantiates a new PriceTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTierWithDefaults() *PriceTier {
	this := PriceTier{}
	return &this
}

// GetData returns the Data field value
func (o *PriceTier) GetData() PriceTierData {
	if o == nil {
		var ret PriceTierData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceTier) GetDataOk() (*PriceTierData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceTier) SetData(v PriceTierData) {
	o.Data = v
}

func (o PriceTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceTier struct {
	value *PriceTier
	isSet bool
}

func (v NullablePriceTier) Get() *PriceTier {
	return v.value
}

func (v *NullablePriceTier) Set(val *PriceTier) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTier) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTier(val *PriceTier) *NullablePriceTier {
	return &NullablePriceTier{value: val, isSet: true}
}

func (v NullablePriceTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

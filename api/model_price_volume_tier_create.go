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

// PriceVolumeTierCreate struct for PriceVolumeTierCreate
type PriceVolumeTierCreate struct {
	Data PriceVolumeTierCreateData `json:"data"`
}

// NewPriceVolumeTierCreate instantiates a new PriceVolumeTierCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTierCreate(data PriceVolumeTierCreateData) *PriceVolumeTierCreate {
	this := PriceVolumeTierCreate{}
	this.Data = data
	return &this
}

// NewPriceVolumeTierCreateWithDefaults instantiates a new PriceVolumeTierCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierCreateWithDefaults() *PriceVolumeTierCreate {
	this := PriceVolumeTierCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceVolumeTierCreate) GetData() PriceVolumeTierCreateData {
	if o == nil {
		var ret PriceVolumeTierCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTierCreate) GetDataOk() (*PriceVolumeTierCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceVolumeTierCreate) SetData(v PriceVolumeTierCreateData) {
	o.Data = v
}

func (o PriceVolumeTierCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTierCreate struct {
	value *PriceVolumeTierCreate
	isSet bool
}

func (v NullablePriceVolumeTierCreate) Get() *PriceVolumeTierCreate {
	return v.value
}

func (v *NullablePriceVolumeTierCreate) Set(val *PriceVolumeTierCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTierCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTierCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTierCreate(val *PriceVolumeTierCreate) *NullablePriceVolumeTierCreate {
	return &NullablePriceVolumeTierCreate{value: val, isSet: true}
}

func (v NullablePriceVolumeTierCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTierCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

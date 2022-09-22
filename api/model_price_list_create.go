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

// PriceListCreate struct for PriceListCreate
type PriceListCreate struct {
	Data PriceListCreateData `json:"data"`
}

// NewPriceListCreate instantiates a new PriceListCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListCreate(data PriceListCreateData) *PriceListCreate {
	this := PriceListCreate{}
	this.Data = data
	return &this
}

// NewPriceListCreateWithDefaults instantiates a new PriceListCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListCreateWithDefaults() *PriceListCreate {
	this := PriceListCreate{}
	return &this
}

// GetData returns the Data field value
func (o *PriceListCreate) GetData() PriceListCreateData {
	if o == nil {
		var ret PriceListCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceListCreate) GetDataOk() (*PriceListCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceListCreate) SetData(v PriceListCreateData) {
	o.Data = v
}

func (o PriceListCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListCreate struct {
	value *PriceListCreate
	isSet bool
}

func (v NullablePriceListCreate) Get() *PriceListCreate {
	return v.value
}

func (v *NullablePriceListCreate) Set(val *PriceListCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListCreate(val *PriceListCreate) *NullablePriceListCreate {
	return &NullablePriceListCreate{value: val, isSet: true}
}

func (v NullablePriceListCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

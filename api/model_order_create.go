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

// OrderCreate struct for OrderCreate
type OrderCreate struct {
	Data OrderCreateData `json:"data"`
}

// NewOrderCreate instantiates a new OrderCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreate(data OrderCreateData) *OrderCreate {
	this := OrderCreate{}
	this.Data = data
	return &this
}

// NewOrderCreateWithDefaults instantiates a new OrderCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateWithDefaults() *OrderCreate {
	this := OrderCreate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderCreate) GetData() OrderCreateData {
	if o == nil {
		var ret OrderCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderCreate) GetDataOk() (*OrderCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderCreate) SetData(v OrderCreateData) {
	o.Data = v
}

func (o OrderCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderCreate struct {
	value *OrderCreate
	isSet bool
}

func (v NullableOrderCreate) Get() *OrderCreate {
	return v.value
}

func (v *NullableOrderCreate) Set(val *OrderCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreate(val *OrderCreate) *NullableOrderCreate {
	return &NullableOrderCreate{value: val, isSet: true}
}

func (v NullableOrderCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

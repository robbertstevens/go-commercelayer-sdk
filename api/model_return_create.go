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

// ReturnCreate struct for ReturnCreate
type ReturnCreate struct {
	Data ReturnCreateData `json:"data"`
}

// NewReturnCreate instantiates a new ReturnCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnCreate(data ReturnCreateData) *ReturnCreate {
	this := ReturnCreate{}
	this.Data = data
	return &this
}

// NewReturnCreateWithDefaults instantiates a new ReturnCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnCreateWithDefaults() *ReturnCreate {
	this := ReturnCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnCreate) GetData() ReturnCreateData {
	if o == nil {
		var ret ReturnCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnCreate) GetDataOk() (*ReturnCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnCreate) SetData(v ReturnCreateData) {
	o.Data = v
}

func (o ReturnCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnCreate struct {
	value *ReturnCreate
	isSet bool
}

func (v NullableReturnCreate) Get() *ReturnCreate {
	return v.value
}

func (v *NullableReturnCreate) Set(val *ReturnCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnCreate(val *ReturnCreate) *NullableReturnCreate {
	return &NullableReturnCreate{value: val, isSet: true}
}

func (v NullableReturnCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// checks if the SkuCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuCreate{}

// SkuCreate struct for SkuCreate
type SkuCreate struct {
	Data SkuCreateData `json:"data"`
}

// NewSkuCreate instantiates a new SkuCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreate(data SkuCreateData) *SkuCreate {
	this := SkuCreate{}
	this.Data = data
	return &this
}

// NewSkuCreateWithDefaults instantiates a new SkuCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateWithDefaults() *SkuCreate {
	this := SkuCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuCreate) GetData() SkuCreateData {
	if o == nil {
		var ret SkuCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuCreate) GetDataOk() (*SkuCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuCreate) SetData(v SkuCreateData) {
	o.Data = v
}

func (o SkuCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSkuCreate struct {
	value *SkuCreate
	isSet bool
}

func (v NullableSkuCreate) Get() *SkuCreate {
	return v.value
}

func (v *NullableSkuCreate) Set(val *SkuCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreate(val *SkuCreate) *NullableSkuCreate {
	return &NullableSkuCreate{value: val, isSet: true}
}

func (v NullableSkuCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

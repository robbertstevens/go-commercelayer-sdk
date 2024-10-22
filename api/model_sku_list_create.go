/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SkuListCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuListCreate{}

// SkuListCreate struct for SkuListCreate
type SkuListCreate struct {
	Data SkuListCreateData `json:"data"`
}

// NewSkuListCreate instantiates a new SkuListCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListCreate(data SkuListCreateData) *SkuListCreate {
	this := SkuListCreate{}
	this.Data = data
	return &this
}

// NewSkuListCreateWithDefaults instantiates a new SkuListCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListCreateWithDefaults() *SkuListCreate {
	this := SkuListCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuListCreate) GetData() SkuListCreateData {
	if o == nil {
		var ret SkuListCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuListCreate) GetDataOk() (*SkuListCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuListCreate) SetData(v SkuListCreateData) {
	o.Data = v
}

func (o SkuListCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuListCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSkuListCreate struct {
	value *SkuListCreate
	isSet bool
}

func (v NullableSkuListCreate) Get() *SkuListCreate {
	return v.value
}

func (v *NullableSkuListCreate) Set(val *SkuListCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListCreate(val *SkuListCreate) *NullableSkuListCreate {
	return &NullableSkuListCreate{value: val, isSet: true}
}

func (v NullableSkuListCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

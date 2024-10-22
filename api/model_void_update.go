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

// checks if the VoidUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoidUpdate{}

// VoidUpdate struct for VoidUpdate
type VoidUpdate struct {
	Data VoidUpdateData `json:"data"`
}

// NewVoidUpdate instantiates a new VoidUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidUpdate(data VoidUpdateData) *VoidUpdate {
	this := VoidUpdate{}
	this.Data = data
	return &this
}

// NewVoidUpdateWithDefaults instantiates a new VoidUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidUpdateWithDefaults() *VoidUpdate {
	this := VoidUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *VoidUpdate) GetData() VoidUpdateData {
	if o == nil {
		var ret VoidUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VoidUpdate) GetDataOk() (*VoidUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VoidUpdate) SetData(v VoidUpdateData) {
	o.Data = v
}

func (o VoidUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoidUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableVoidUpdate struct {
	value *VoidUpdate
	isSet bool
}

func (v NullableVoidUpdate) Get() *VoidUpdate {
	return v.value
}

func (v *NullableVoidUpdate) Set(val *VoidUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidUpdate(val *VoidUpdate) *NullableVoidUpdate {
	return &NullableVoidUpdate{value: val, isSet: true}
}

func (v NullableVoidUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

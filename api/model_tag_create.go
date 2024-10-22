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

// checks if the TagCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagCreate{}

// TagCreate struct for TagCreate
type TagCreate struct {
	Data TagCreateData `json:"data"`
}

// NewTagCreate instantiates a new TagCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCreate(data TagCreateData) *TagCreate {
	this := TagCreate{}
	this.Data = data
	return &this
}

// NewTagCreateWithDefaults instantiates a new TagCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCreateWithDefaults() *TagCreate {
	this := TagCreate{}
	return &this
}

// GetData returns the Data field value
func (o *TagCreate) GetData() TagCreateData {
	if o == nil {
		var ret TagCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TagCreate) GetDataOk() (*TagCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TagCreate) SetData(v TagCreateData) {
	o.Data = v
}

func (o TagCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableTagCreate struct {
	value *TagCreate
	isSet bool
}

func (v NullableTagCreate) Get() *TagCreate {
	return v.value
}

func (v *NullableTagCreate) Set(val *TagCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCreate(val *TagCreate) *NullableTagCreate {
	return &NullableTagCreate{value: val, isSet: true}
}

func (v NullableTagCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

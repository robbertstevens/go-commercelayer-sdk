/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LineItemOptionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemOptionUpdate{}

// LineItemOptionUpdate struct for LineItemOptionUpdate
type LineItemOptionUpdate struct {
	Data PATCHLineItemOptionsLineItemOptionIdRequestData `json:"data"`
}

// NewLineItemOptionUpdate instantiates a new LineItemOptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionUpdate(data PATCHLineItemOptionsLineItemOptionIdRequestData) *LineItemOptionUpdate {
	this := LineItemOptionUpdate{}
	this.Data = data
	return &this
}

// NewLineItemOptionUpdateWithDefaults instantiates a new LineItemOptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionUpdateWithDefaults() *LineItemOptionUpdate {
	this := LineItemOptionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemOptionUpdate) GetData() PATCHLineItemOptionsLineItemOptionIdRequestData {
	if o == nil {
		var ret PATCHLineItemOptionsLineItemOptionIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdate) GetDataOk() (*PATCHLineItemOptionsLineItemOptionIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemOptionUpdate) SetData(v PATCHLineItemOptionsLineItemOptionIdRequestData) {
	o.Data = v
}

func (o LineItemOptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemOptionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableLineItemOptionUpdate struct {
	value *LineItemOptionUpdate
	isSet bool
}

func (v NullableLineItemOptionUpdate) Get() *LineItemOptionUpdate {
	return v.value
}

func (v *NullableLineItemOptionUpdate) Set(val *LineItemOptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionUpdate(val *LineItemOptionUpdate) *NullableLineItemOptionUpdate {
	return &NullableLineItemOptionUpdate{value: val, isSet: true}
}

func (v NullableLineItemOptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

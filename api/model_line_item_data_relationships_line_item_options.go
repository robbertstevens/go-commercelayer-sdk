/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemDataRelationshipsLineItemOptions struct for LineItemDataRelationshipsLineItemOptions
type LineItemDataRelationshipsLineItemOptions struct {
	Data LineItemDataRelationshipsLineItemOptionsData `json:"data"`
}

// NewLineItemDataRelationshipsLineItemOptions instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsLineItemOptions(data LineItemDataRelationshipsLineItemOptionsData) *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	this.Data = data
	return &this
}

// NewLineItemDataRelationshipsLineItemOptionsWithDefaults instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsLineItemOptionsWithDefaults() *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemDataRelationshipsLineItemOptions) GetData() LineItemDataRelationshipsLineItemOptionsData {
	if o == nil {
		var ret LineItemDataRelationshipsLineItemOptionsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsLineItemOptions) GetDataOk() (*LineItemDataRelationshipsLineItemOptionsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemDataRelationshipsLineItemOptions) SetData(v LineItemDataRelationshipsLineItemOptionsData) {
	o.Data = v
}

func (o LineItemDataRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemDataRelationshipsLineItemOptions struct {
	value *LineItemDataRelationshipsLineItemOptions
	isSet bool
}

func (v NullableLineItemDataRelationshipsLineItemOptions) Get() *LineItemDataRelationshipsLineItemOptions {
	return v.value
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) Set(val *LineItemDataRelationshipsLineItemOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemDataRelationshipsLineItemOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemDataRelationshipsLineItemOptions(val *LineItemDataRelationshipsLineItemOptions) *NullableLineItemDataRelationshipsLineItemOptions {
	return &NullableLineItemDataRelationshipsLineItemOptions{value: val, isSet: true}
}

func (v NullableLineItemDataRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemDataRelationshipsLineItemOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

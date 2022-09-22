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

// LineItemOptionDataRelationshipsLineItem struct for LineItemOptionDataRelationshipsLineItem
type LineItemOptionDataRelationshipsLineItem struct {
	Data LineItemOptionDataRelationshipsLineItemData `json:"data"`
}

// NewLineItemOptionDataRelationshipsLineItem instantiates a new LineItemOptionDataRelationshipsLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataRelationshipsLineItem(data LineItemOptionDataRelationshipsLineItemData) *LineItemOptionDataRelationshipsLineItem {
	this := LineItemOptionDataRelationshipsLineItem{}
	this.Data = data
	return &this
}

// NewLineItemOptionDataRelationshipsLineItemWithDefaults instantiates a new LineItemOptionDataRelationshipsLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataRelationshipsLineItemWithDefaults() *LineItemOptionDataRelationshipsLineItem {
	this := LineItemOptionDataRelationshipsLineItem{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemOptionDataRelationshipsLineItem) GetData() LineItemOptionDataRelationshipsLineItemData {
	if o == nil {
		var ret LineItemOptionDataRelationshipsLineItemData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationshipsLineItem) GetDataOk() (*LineItemOptionDataRelationshipsLineItemData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemOptionDataRelationshipsLineItem) SetData(v LineItemOptionDataRelationshipsLineItemData) {
	o.Data = v
}

func (o LineItemOptionDataRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionDataRelationshipsLineItem struct {
	value *LineItemOptionDataRelationshipsLineItem
	isSet bool
}

func (v NullableLineItemOptionDataRelationshipsLineItem) Get() *LineItemOptionDataRelationshipsLineItem {
	return v.value
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) Set(val *LineItemOptionDataRelationshipsLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataRelationshipsLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataRelationshipsLineItem(val *LineItemOptionDataRelationshipsLineItem) *NullableLineItemOptionDataRelationshipsLineItem {
	return &NullableLineItemOptionDataRelationshipsLineItem{value: val, isSet: true}
}

func (v NullableLineItemOptionDataRelationshipsLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataRelationshipsLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

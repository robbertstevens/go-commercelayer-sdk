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

// ParcelLineItem struct for ParcelLineItem
type ParcelLineItem struct {
	Data ParcelLineItemData `json:"data"`
}

// NewParcelLineItem instantiates a new ParcelLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItem(data ParcelLineItemData) *ParcelLineItem {
	this := ParcelLineItem{}
	this.Data = data
	return &this
}

// NewParcelLineItemWithDefaults instantiates a new ParcelLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemWithDefaults() *ParcelLineItem {
	this := ParcelLineItem{}
	return &this
}

// GetData returns the Data field value
func (o *ParcelLineItem) GetData() ParcelLineItemData {
	if o == nil {
		var ret ParcelLineItemData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItem) GetDataOk() (*ParcelLineItemData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ParcelLineItem) SetData(v ParcelLineItemData) {
	o.Data = v
}

func (o ParcelLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableParcelLineItem struct {
	value *ParcelLineItem
	isSet bool
}

func (v NullableParcelLineItem) Get() *ParcelLineItem {
	return v.value
}

func (v *NullableParcelLineItem) Set(val *ParcelLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItem(val *ParcelLineItem) *NullableParcelLineItem {
	return &NullableParcelLineItem{value: val, isSet: true}
}

func (v NullableParcelLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



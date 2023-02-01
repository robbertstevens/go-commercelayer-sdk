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

// LineItemOptionCreateDataRelationshipsSkuOption struct for LineItemOptionCreateDataRelationshipsSkuOption
type LineItemOptionCreateDataRelationshipsSkuOption struct {
	Data LineItemOptionDataRelationshipsSkuOptionData `json:"data"`
}

// NewLineItemOptionCreateDataRelationshipsSkuOption instantiates a new LineItemOptionCreateDataRelationshipsSkuOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionCreateDataRelationshipsSkuOption(data LineItemOptionDataRelationshipsSkuOptionData) *LineItemOptionCreateDataRelationshipsSkuOption {
	this := LineItemOptionCreateDataRelationshipsSkuOption{}
	this.Data = data
	return &this
}

// NewLineItemOptionCreateDataRelationshipsSkuOptionWithDefaults instantiates a new LineItemOptionCreateDataRelationshipsSkuOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionCreateDataRelationshipsSkuOptionWithDefaults() *LineItemOptionCreateDataRelationshipsSkuOption {
	this := LineItemOptionCreateDataRelationshipsSkuOption{}
	return &this
}

// GetData returns the Data field value
func (o *LineItemOptionCreateDataRelationshipsSkuOption) GetData() LineItemOptionDataRelationshipsSkuOptionData {
	if o == nil {
		var ret LineItemOptionDataRelationshipsSkuOptionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataRelationshipsSkuOption) GetDataOk() (*LineItemOptionDataRelationshipsSkuOptionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LineItemOptionCreateDataRelationshipsSkuOption) SetData(v LineItemOptionDataRelationshipsSkuOptionData) {
	o.Data = v
}

func (o LineItemOptionCreateDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionCreateDataRelationshipsSkuOption struct {
	value *LineItemOptionCreateDataRelationshipsSkuOption
	isSet bool
}

func (v NullableLineItemOptionCreateDataRelationshipsSkuOption) Get() *LineItemOptionCreateDataRelationshipsSkuOption {
	return v.value
}

func (v *NullableLineItemOptionCreateDataRelationshipsSkuOption) Set(val *LineItemOptionCreateDataRelationshipsSkuOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionCreateDataRelationshipsSkuOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionCreateDataRelationshipsSkuOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionCreateDataRelationshipsSkuOption(val *LineItemOptionCreateDataRelationshipsSkuOption) *NullableLineItemOptionCreateDataRelationshipsSkuOption {
	return &NullableLineItemOptionCreateDataRelationshipsSkuOption{value: val, isSet: true}
}

func (v NullableLineItemOptionCreateDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionCreateDataRelationshipsSkuOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

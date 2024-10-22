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

// checks if the LineItemOptionDataRelationshipsSkuOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemOptionDataRelationshipsSkuOption{}

// LineItemOptionDataRelationshipsSkuOption struct for LineItemOptionDataRelationshipsSkuOption
type LineItemOptionDataRelationshipsSkuOption struct {
	Data *LineItemOptionDataRelationshipsSkuOptionData `json:"data,omitempty"`
}

// NewLineItemOptionDataRelationshipsSkuOption instantiates a new LineItemOptionDataRelationshipsSkuOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataRelationshipsSkuOption() *LineItemOptionDataRelationshipsSkuOption {
	this := LineItemOptionDataRelationshipsSkuOption{}
	return &this
}

// NewLineItemOptionDataRelationshipsSkuOptionWithDefaults instantiates a new LineItemOptionDataRelationshipsSkuOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataRelationshipsSkuOptionWithDefaults() *LineItemOptionDataRelationshipsSkuOption {
	this := LineItemOptionDataRelationshipsSkuOption{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemOptionDataRelationshipsSkuOption) GetData() LineItemOptionDataRelationshipsSkuOptionData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemOptionDataRelationshipsSkuOptionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationshipsSkuOption) GetDataOk() (*LineItemOptionDataRelationshipsSkuOptionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemOptionDataRelationshipsSkuOption) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemOptionDataRelationshipsSkuOptionData and assigns it to the Data field.
func (o *LineItemOptionDataRelationshipsSkuOption) SetData(v LineItemOptionDataRelationshipsSkuOptionData) {
	o.Data = &v
}

func (o LineItemOptionDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemOptionDataRelationshipsSkuOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLineItemOptionDataRelationshipsSkuOption struct {
	value *LineItemOptionDataRelationshipsSkuOption
	isSet bool
}

func (v NullableLineItemOptionDataRelationshipsSkuOption) Get() *LineItemOptionDataRelationshipsSkuOption {
	return v.value
}

func (v *NullableLineItemOptionDataRelationshipsSkuOption) Set(val *LineItemOptionDataRelationshipsSkuOption) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataRelationshipsSkuOption) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataRelationshipsSkuOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataRelationshipsSkuOption(val *LineItemOptionDataRelationshipsSkuOption) *NullableLineItemOptionDataRelationshipsSkuOption {
	return &NullableLineItemOptionDataRelationshipsSkuOption{value: val, isSet: true}
}

func (v NullableLineItemOptionDataRelationshipsSkuOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataRelationshipsSkuOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

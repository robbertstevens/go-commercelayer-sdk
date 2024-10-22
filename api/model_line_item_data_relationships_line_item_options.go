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

// checks if the LineItemDataRelationshipsLineItemOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemDataRelationshipsLineItemOptions{}

// LineItemDataRelationshipsLineItemOptions struct for LineItemDataRelationshipsLineItemOptions
type LineItemDataRelationshipsLineItemOptions struct {
	Data *LineItemDataRelationshipsLineItemOptionsData `json:"data,omitempty"`
}

// NewLineItemDataRelationshipsLineItemOptions instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemDataRelationshipsLineItemOptions() *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	return &this
}

// NewLineItemDataRelationshipsLineItemOptionsWithDefaults instantiates a new LineItemDataRelationshipsLineItemOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemDataRelationshipsLineItemOptionsWithDefaults() *LineItemDataRelationshipsLineItemOptions {
	this := LineItemDataRelationshipsLineItemOptions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LineItemDataRelationshipsLineItemOptions) GetData() LineItemDataRelationshipsLineItemOptionsData {
	if o == nil || IsNil(o.Data) {
		var ret LineItemDataRelationshipsLineItemOptionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemDataRelationshipsLineItemOptions) GetDataOk() (*LineItemDataRelationshipsLineItemOptionsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LineItemDataRelationshipsLineItemOptions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LineItemDataRelationshipsLineItemOptionsData and assigns it to the Data field.
func (o *LineItemDataRelationshipsLineItemOptions) SetData(v LineItemDataRelationshipsLineItemOptionsData) {
	o.Data = &v
}

func (o LineItemDataRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemDataRelationshipsLineItemOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
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

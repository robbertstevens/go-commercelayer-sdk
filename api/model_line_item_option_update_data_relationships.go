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

// LineItemOptionUpdateDataRelationships struct for LineItemOptionUpdateDataRelationships
type LineItemOptionUpdateDataRelationships struct {
	SkuOption *LineItemOptionDataRelationshipsSkuOption `json:"sku_option,omitempty"`
}

// NewLineItemOptionUpdateDataRelationships instantiates a new LineItemOptionUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionUpdateDataRelationships() *LineItemOptionUpdateDataRelationships {
	this := LineItemOptionUpdateDataRelationships{}
	return &this
}

// NewLineItemOptionUpdateDataRelationshipsWithDefaults instantiates a new LineItemOptionUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionUpdateDataRelationshipsWithDefaults() *LineItemOptionUpdateDataRelationships {
	this := LineItemOptionUpdateDataRelationships{}
	return &this
}

// GetSkuOption returns the SkuOption field value if set, zero value otherwise.
func (o *LineItemOptionUpdateDataRelationships) GetSkuOption() LineItemOptionDataRelationshipsSkuOption {
	if o == nil || o.SkuOption == nil {
		var ret LineItemOptionDataRelationshipsSkuOption
		return ret
	}
	return *o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionUpdateDataRelationships) GetSkuOptionOk() (*LineItemOptionDataRelationshipsSkuOption, bool) {
	if o == nil || o.SkuOption == nil {
		return nil, false
	}
	return o.SkuOption, true
}

// HasSkuOption returns a boolean if a field has been set.
func (o *LineItemOptionUpdateDataRelationships) HasSkuOption() bool {
	if o != nil && o.SkuOption != nil {
		return true
	}

	return false
}

// SetSkuOption gets a reference to the given LineItemOptionDataRelationshipsSkuOption and assigns it to the SkuOption field.
func (o *LineItemOptionUpdateDataRelationships) SetSkuOption(v LineItemOptionDataRelationshipsSkuOption) {
	o.SkuOption = &v
}

func (o LineItemOptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuOption != nil {
		toSerialize["sku_option"] = o.SkuOption
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionUpdateDataRelationships struct {
	value *LineItemOptionUpdateDataRelationships
	isSet bool
}

func (v NullableLineItemOptionUpdateDataRelationships) Get() *LineItemOptionUpdateDataRelationships {
	return v.value
}

func (v *NullableLineItemOptionUpdateDataRelationships) Set(val *LineItemOptionUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionUpdateDataRelationships(val *LineItemOptionUpdateDataRelationships) *NullableLineItemOptionUpdateDataRelationships {
	return &NullableLineItemOptionUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableLineItemOptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

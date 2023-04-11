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

// checks if the PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships{}

// PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships struct for PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships
type PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships struct {
	SkuOption *POSTLineItemOptionsRequestDataRelationshipsSkuOption `json:"sku_option,omitempty"`
}

// NewPATCHLineItemOptionsLineItemOptionIdRequestDataRelationships instantiates a new PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemOptionsLineItemOptionIdRequestDataRelationships() *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships {
	this := PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships{}
	return &this
}

// NewPATCHLineItemOptionsLineItemOptionIdRequestDataRelationshipsWithDefaults instantiates a new PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemOptionsLineItemOptionIdRequestDataRelationshipsWithDefaults() *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships {
	this := PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships{}
	return &this
}

// GetSkuOption returns the SkuOption field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) GetSkuOption() POSTLineItemOptionsRequestDataRelationshipsSkuOption {
	if o == nil || IsNil(o.SkuOption) {
		var ret POSTLineItemOptionsRequestDataRelationshipsSkuOption
		return ret
	}
	return *o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) GetSkuOptionOk() (*POSTLineItemOptionsRequestDataRelationshipsSkuOption, bool) {
	if o == nil || IsNil(o.SkuOption) {
		return nil, false
	}
	return o.SkuOption, true
}

// HasSkuOption returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) HasSkuOption() bool {
	if o != nil && !IsNil(o.SkuOption) {
		return true
	}

	return false
}

// SetSkuOption gets a reference to the given POSTLineItemOptionsRequestDataRelationshipsSkuOption and assigns it to the SkuOption field.
func (o *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) SetSkuOption(v POSTLineItemOptionsRequestDataRelationshipsSkuOption) {
	o.SkuOption = &v
}

func (o PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkuOption) {
		toSerialize["sku_option"] = o.SkuOption
	}
	return toSerialize, nil
}

type NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships struct {
	value *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) Get() *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) Set(val *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships(val *PATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) *NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships {
	return &NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemOptionsLineItemOptionIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

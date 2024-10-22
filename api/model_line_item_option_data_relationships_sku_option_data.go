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

// checks if the LineItemOptionDataRelationshipsSkuOptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemOptionDataRelationshipsSkuOptionData{}

// LineItemOptionDataRelationshipsSkuOptionData struct for LineItemOptionDataRelationshipsSkuOptionData
type LineItemOptionDataRelationshipsSkuOptionData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewLineItemOptionDataRelationshipsSkuOptionData instantiates a new LineItemOptionDataRelationshipsSkuOptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataRelationshipsSkuOptionData() *LineItemOptionDataRelationshipsSkuOptionData {
	this := LineItemOptionDataRelationshipsSkuOptionData{}
	return &this
}

// NewLineItemOptionDataRelationshipsSkuOptionDataWithDefaults instantiates a new LineItemOptionDataRelationshipsSkuOptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataRelationshipsSkuOptionDataWithDefaults() *LineItemOptionDataRelationshipsSkuOptionData {
	this := LineItemOptionDataRelationshipsSkuOptionData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemOptionDataRelationshipsSkuOptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemOptionDataRelationshipsSkuOptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LineItemOptionDataRelationshipsSkuOptionData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *LineItemOptionDataRelationshipsSkuOptionData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LineItemOptionDataRelationshipsSkuOptionData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LineItemOptionDataRelationshipsSkuOptionData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LineItemOptionDataRelationshipsSkuOptionData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *LineItemOptionDataRelationshipsSkuOptionData) SetId(v interface{}) {
	o.Id = v
}

func (o LineItemOptionDataRelationshipsSkuOptionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemOptionDataRelationshipsSkuOptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableLineItemOptionDataRelationshipsSkuOptionData struct {
	value *LineItemOptionDataRelationshipsSkuOptionData
	isSet bool
}

func (v NullableLineItemOptionDataRelationshipsSkuOptionData) Get() *LineItemOptionDataRelationshipsSkuOptionData {
	return v.value
}

func (v *NullableLineItemOptionDataRelationshipsSkuOptionData) Set(val *LineItemOptionDataRelationshipsSkuOptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataRelationshipsSkuOptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataRelationshipsSkuOptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataRelationshipsSkuOptionData(val *LineItemOptionDataRelationshipsSkuOptionData) *NullableLineItemOptionDataRelationshipsSkuOptionData {
	return &NullableLineItemOptionDataRelationshipsSkuOptionData{value: val, isSet: true}
}

func (v NullableLineItemOptionDataRelationshipsSkuOptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataRelationshipsSkuOptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

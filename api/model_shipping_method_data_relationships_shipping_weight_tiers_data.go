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

// checks if the ShippingMethodDataRelationshipsShippingWeightTiersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingMethodDataRelationshipsShippingWeightTiersData{}

// ShippingMethodDataRelationshipsShippingWeightTiersData struct for ShippingMethodDataRelationshipsShippingWeightTiersData
type ShippingMethodDataRelationshipsShippingWeightTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewShippingMethodDataRelationshipsShippingWeightTiersData instantiates a new ShippingMethodDataRelationshipsShippingWeightTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodDataRelationshipsShippingWeightTiersData() *ShippingMethodDataRelationshipsShippingWeightTiersData {
	this := ShippingMethodDataRelationshipsShippingWeightTiersData{}
	return &this
}

// NewShippingMethodDataRelationshipsShippingWeightTiersDataWithDefaults instantiates a new ShippingMethodDataRelationshipsShippingWeightTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodDataRelationshipsShippingWeightTiersDataWithDefaults() *ShippingMethodDataRelationshipsShippingWeightTiersData {
	this := ShippingMethodDataRelationshipsShippingWeightTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *ShippingMethodDataRelationshipsShippingWeightTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o ShippingMethodDataRelationshipsShippingWeightTiersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingMethodDataRelationshipsShippingWeightTiersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableShippingMethodDataRelationshipsShippingWeightTiersData struct {
	value *ShippingMethodDataRelationshipsShippingWeightTiersData
	isSet bool
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiersData) Get() *ShippingMethodDataRelationshipsShippingWeightTiersData {
	return v.value
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiersData) Set(val *ShippingMethodDataRelationshipsShippingWeightTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodDataRelationshipsShippingWeightTiersData(val *ShippingMethodDataRelationshipsShippingWeightTiersData) *NullableShippingMethodDataRelationshipsShippingWeightTiersData {
	return &NullableShippingMethodDataRelationshipsShippingWeightTiersData{value: val, isSet: true}
}

func (v NullableShippingMethodDataRelationshipsShippingWeightTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodDataRelationshipsShippingWeightTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

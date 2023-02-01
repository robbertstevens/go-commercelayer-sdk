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

// ShippingMethodCreateDataRelationshipsShippingZone struct for ShippingMethodCreateDataRelationshipsShippingZone
type ShippingMethodCreateDataRelationshipsShippingZone struct {
	Data ShippingMethodDataRelationshipsShippingZoneData `json:"data"`
}

// NewShippingMethodCreateDataRelationshipsShippingZone instantiates a new ShippingMethodCreateDataRelationshipsShippingZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreateDataRelationshipsShippingZone(data ShippingMethodDataRelationshipsShippingZoneData) *ShippingMethodCreateDataRelationshipsShippingZone {
	this := ShippingMethodCreateDataRelationshipsShippingZone{}
	this.Data = data
	return &this
}

// NewShippingMethodCreateDataRelationshipsShippingZoneWithDefaults instantiates a new ShippingMethodCreateDataRelationshipsShippingZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateDataRelationshipsShippingZoneWithDefaults() *ShippingMethodCreateDataRelationshipsShippingZone {
	this := ShippingMethodCreateDataRelationshipsShippingZone{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingMethodCreateDataRelationshipsShippingZone) GetData() ShippingMethodDataRelationshipsShippingZoneData {
	if o == nil {
		var ret ShippingMethodDataRelationshipsShippingZoneData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationshipsShippingZone) GetDataOk() (*ShippingMethodDataRelationshipsShippingZoneData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingMethodCreateDataRelationshipsShippingZone) SetData(v ShippingMethodDataRelationshipsShippingZoneData) {
	o.Data = v
}

func (o ShippingMethodCreateDataRelationshipsShippingZone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodCreateDataRelationshipsShippingZone struct {
	value *ShippingMethodCreateDataRelationshipsShippingZone
	isSet bool
}

func (v NullableShippingMethodCreateDataRelationshipsShippingZone) Get() *ShippingMethodCreateDataRelationshipsShippingZone {
	return v.value
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingZone) Set(val *ShippingMethodCreateDataRelationshipsShippingZone) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreateDataRelationshipsShippingZone) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreateDataRelationshipsShippingZone(val *ShippingMethodCreateDataRelationshipsShippingZone) *NullableShippingMethodCreateDataRelationshipsShippingZone {
	return &NullableShippingMethodCreateDataRelationshipsShippingZone{value: val, isSet: true}
}

func (v NullableShippingMethodCreateDataRelationshipsShippingZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreateDataRelationshipsShippingZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

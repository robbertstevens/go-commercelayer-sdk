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

// InventoryModelDataRelationshipsInventoryReturnLocations struct for InventoryModelDataRelationshipsInventoryReturnLocations
type InventoryModelDataRelationshipsInventoryReturnLocations struct {
	Data InventoryModelDataRelationshipsInventoryReturnLocationsData `json:"data"`
}

// NewInventoryModelDataRelationshipsInventoryReturnLocations instantiates a new InventoryModelDataRelationshipsInventoryReturnLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelDataRelationshipsInventoryReturnLocations(data InventoryModelDataRelationshipsInventoryReturnLocationsData) *InventoryModelDataRelationshipsInventoryReturnLocations {
	this := InventoryModelDataRelationshipsInventoryReturnLocations{}
	this.Data = data
	return &this
}

// NewInventoryModelDataRelationshipsInventoryReturnLocationsWithDefaults instantiates a new InventoryModelDataRelationshipsInventoryReturnLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelDataRelationshipsInventoryReturnLocationsWithDefaults() *InventoryModelDataRelationshipsInventoryReturnLocations {
	this := InventoryModelDataRelationshipsInventoryReturnLocations{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryModelDataRelationshipsInventoryReturnLocations) GetData() InventoryModelDataRelationshipsInventoryReturnLocationsData {
	if o == nil {
		var ret InventoryModelDataRelationshipsInventoryReturnLocationsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryModelDataRelationshipsInventoryReturnLocations) GetDataOk() (*InventoryModelDataRelationshipsInventoryReturnLocationsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryModelDataRelationshipsInventoryReturnLocations) SetData(v InventoryModelDataRelationshipsInventoryReturnLocationsData) {
	o.Data = v
}

func (o InventoryModelDataRelationshipsInventoryReturnLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelDataRelationshipsInventoryReturnLocations struct {
	value *InventoryModelDataRelationshipsInventoryReturnLocations
	isSet bool
}

func (v NullableInventoryModelDataRelationshipsInventoryReturnLocations) Get() *InventoryModelDataRelationshipsInventoryReturnLocations {
	return v.value
}

func (v *NullableInventoryModelDataRelationshipsInventoryReturnLocations) Set(val *InventoryModelDataRelationshipsInventoryReturnLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelDataRelationshipsInventoryReturnLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelDataRelationshipsInventoryReturnLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelDataRelationshipsInventoryReturnLocations(val *InventoryModelDataRelationshipsInventoryReturnLocations) *NullableInventoryModelDataRelationshipsInventoryReturnLocations {
	return &NullableInventoryModelDataRelationshipsInventoryReturnLocations{value: val, isSet: true}
}

func (v NullableInventoryModelDataRelationshipsInventoryReturnLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelDataRelationshipsInventoryReturnLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


